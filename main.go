package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/dgrr/fastws"
	"github.com/segmentio/encoding/json"
	"github.com/ugjka/messenger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

//go:generate go run github.com/ACFUN-FOSS/acfunlive-backend/cmd -o cmd_gen.go

func main() {
	port := flag.Uint("port", 0, "WebSocket server port, default 15368")
	isDebug = flag.Bool("debug", false, "debug")
	logFileName := flag.String("logfile", "", "log file location")
	flag.Parse()
	if !(*port > 1023 && *port < 65536) {
		// 默认端口为15368
		*port = 15368
	}
	if *logFileName != "" {
		if _, err := os.Stat(*logFileName); err == nil {
			if err := os.Rename(*logFileName, *logFileName+".bak"); err != nil {
				panic(err)
			}
		} else if !os.IsNotExist(err) {
			panic(err)
		}
		logFile, err := os.OpenFile(*logFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		defer logFile.Close()
		*isDebug = true
		log.SetOutput(logFile)
		syscall.Dup2(int(logFile.Fd()), 2)
	}
	debug("WebSocket server port is %d", *port)

	server_ch = messenger.New(1024, false)

	server := &fasthttp.Server{
		Handler: fastws.Upgrade(wsHandler),
	}

	go func() {
		if err := server.ListenAndServe(fmt.Sprintf(":%d", *port)); err != nil {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	signal.Stop(ch)
	signal.Reset(os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	debug("Server shutdown")
	server.Shutdown()
}

// 打印调试信息
func debug(format string, v ...interface{}) {
	if *isDebug {
		log.Printf(format, v...)
	}
}

// 打印调试信息
func (conn *wsConn) debug(format string, v ...interface{}) {
	if *isDebug {
		addr := fmt.Sprintf("[%s] ", conn.remoteAddr)
		log.Printf(addr+format, v...)
	}
}

// 发送WebSocket消息
func (conn *wsConn) send(msg string) error {
	conn.debug("Send message: %s", msg)
	_, err := conn.c.WriteString(msg)
	if err != nil {
		conn.debug("Failed to send message: %s, error: %v", msg, err)
	}
	return err
}

// 处理WebSocket连接
func wsHandler(c *fastws.Conn) {
	conn := &wsConn{
		c:          c,
		remoteAddr: c.RemoteAddr().String(),
	}
	defer conn.debug("WebSocket connection close")
	conn.debug("WebSocket connection open")

	ch, err := server_ch.Sub()
	if err != nil {
		conn.debug("Server's main channel has been killed")
		return
	}
	defer server_ch.Unsub(ch)

	var pool fastjson.ParserPool
	var msg []byte
	var clientID string
	// map(int64, *acLive)
	acMap := new(sync.Map)
	var mu sync.RWMutex
	var ac *acLive

	go func() {
		for {
			msg, ok := <-ch
			if ok {
				msg := msg.(*forwardMsg)
				if msg.clientID == "" || msg.clientID == clientID {
					data, err := json.Marshal(msg)
					if err != nil {
						conn.debug("Forward message error: cannot marshal to json: %+v", msg)
						go conn.send(fmt.Sprintf(respErrJSON, forwardDataType, quote(msg.requestID), reqHandleErr, quote(err.Error())))
					} else {
						go conn.send(fmt.Sprintf(respJSON, forwardDataType, quote(msg.requestID), string(data)))
					}
				}
			} else {
				break
			}
		}
	}()

	for {
		_, msg, err = c.ReadMessage(msg[:0])
		if err != nil {
			if !errors.Is(err, fastws.EOF) {
				conn.debug("WebSocket error: %v", err)
			}
			break
		}

		p := pool.Get()
		v, err := p.ParseBytes(msg)
		if err != nil {
			conn.debug("Parsing json error: %v, request data: %s", err, string(msg))
			go conn.send(fmt.Sprintf(respErrJSON, 0, "", jsonParseErr, quote(fmt.Sprintf("error: %v, request data: %s", err, string(msg)))))
			pool.Put(p)
			continue
		}

		reqType := v.GetInt("type")
		reqID := string(v.GetStringBytes("requestID"))
		if reqType != loginType {
			conn.debug("Recieve message: %s", string(msg))
		}
		mu.RLock()
		if ac == nil && reqType != heartbeatType && reqType != loginType && reqType != setClientIDType && reqType != requestForwardDataType {
			go conn.send(fmt.Sprintf(respErrJSON, reqType, quote(reqID), needLogin, quote("Need login")))
			pool.Put(p)
			mu.RUnlock()
			continue
		}
		mu.RUnlock()

		switch reqType {
		case heartbeatType:
			go conn.send(heartbeatJSON)
			pool.Put(p)
		case loginType:
			account := string(v.GetStringBytes("data", "account"))
			password := string(v.GetStringBytes("data", "password"))
			go func() {
				resp := conn.login(acMap, account, password, reqID)
				if aci, ok := acMap.Load(0); ok {
					mu.Lock()
					ac = aci.(*acLive)
					mu.Unlock()
				}
				//_ = conn.send(resp)
				_, _ = c.WriteString(resp)
			}()
			pool.Put(p)
		case setClientIDType:
			clientID = string(v.GetStringBytes("data", "clientID"))
			go conn.send(fmt.Sprintf(respNoDataJSON, setClientIDType, quote(reqID)))
			pool.Put(p)
		case requestForwardDataType:
			msg := new(forwardMsg)
			msg.requestID = reqID
			msg.SourceID = clientID
			msg.clientID = string(v.GetStringBytes("data", "clientID"))
			msg.Message = string(v.GetStringBytes("data", "message"))
			go func() {
				server_ch.Broadcast(msg)
				_ = conn.send(fmt.Sprintf(respNoDataJSON, requestForwardDataType, quote(reqID)))
			}()
			pool.Put(p)
		case getDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				conn.debug("getDanmu: liverUID not exist or less than 1")
				go conn.send(fmt.Sprintf(respErrJSON, getDanmuType, quote(reqID), invalidReqData, quote("liverUID not exist or less than 1")))
			} else {
				go conn.getDanmu(acMap, uid, reqID)
			}
			pool.Put(p)
		case stopDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				conn.debug("stopDanmu: liverUID not exist or less than 1")
				go conn.send(fmt.Sprintf(respErrJSON, stopDanmuType, quote(reqID), invalidReqData, quote("liverUID not exist or less than 1")))
			} else {
				go conn.stopDanmu(acMap, uid, reqID)
			}
			pool.Put(p)
		default:
			if f, ok := cmdDispatch[reqType]; ok {
				go func() {
					mu.RLock()
					resp := f(ac, v, reqID)
					mu.RUnlock()
					_ = conn.send(resp)
					pool.Put(p)
				}()
			} else {
				conn.debug("Error: unknown request type: %d", reqType)
				go conn.send(fmt.Sprintf(respErrJSON, reqType, quote(reqID), invalidReqType, quote("Unknown request type")))
				pool.Put(p)
			}
		}
	}
}
