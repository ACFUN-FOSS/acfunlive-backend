package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/leemcloughlin/logfile"
	"github.com/orzogc/acfundanmu"
	"github.com/orzogc/fastws"
	"github.com/segmentio/encoding/json"
	"github.com/ugjka/messenger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

//go:generate go run github.com/ACFUN-FOSS/acfunlive-backend/cmd -o cmd_gen.go

func main() {
	port := flag.Uint("port", 0, "WebSocket server port, default 15368")
	isDebug = flag.Bool("debug", false, "debug")
	isTCP = flag.Bool("tcp", false, "Danmu client connects server with TCP")
	isLogAll = flag.Bool("logall", false, "log all debug message")
	flag.Parse()

	if !(*port > 1023 && *port < 65536) {
		// 默认端口为 15368
		*port = 15368
	}

	if *isLogAll {
		*isDebug = true
	}

	if logfile.Defaults.FileName != "" {
		var maxSize int64 = 20 * 1024 * 1024
		if logfile.Defaults.MaxSize > 0 {
			maxSize = logfile.Defaults.MaxSize
		}
		oldVersions := 2
		if logfile.Defaults.OldVersions > 0 {
			oldVersions = logfile.Defaults.OldVersions
		}
		logFile, err := logfile.New(
			&logfile.LogFile{
				Flags:       logfile.FileOnly,
				FileName:    logfile.Defaults.FileName,
				MaxSize:     maxSize,
				OldVersions: oldVersions,
			})
		if err != nil {
			log.Panicf("Failed to create logFile %s: %s\n", logfile.Defaults.FileName, err)
		}
		defer logFile.Close()
		*isDebug = true
		log.SetOutput(logFile)
		panicFile, err := os.OpenFile(logfile.Defaults.FileName+".panic", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Panicf("Failed to open logFile %s: %s\n", logfile.Defaults.FileName, err)
		}
		redirectStderr(panicFile)
	}

	server_ch = messenger.New(1024, false)

	server := &fasthttp.Server{
		Handler:      fastws.Upgrade(wsHandler),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  idleTimeout,
		TCPKeepalive: true,
	}

	go func() {
		if err := server.ListenAndServe(fmt.Sprintf(":%d", *port)); err != nil {
			log.Printf("WebSocket server error: %v", err)
			os.Exit(1)
		}
	}()
	debug("WebSocket server is running, the port is %d", *port)
	defer debug("WebSocket server is stopping")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	signal.Stop(ch)
	signal.Reset(os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	server.Shutdown()
}

func danmuClient() acfundanmu.DanmuClient {
	if *isTCP {
		return &acfundanmu.TCPDanmuClient{}
	} else {
		return &acfundanmu.WebSocketDanmuClient{}
	}
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

// 打印调试信息，isLogAll 为 true 才会打印
func (conn *wsConn) debugAll(format string, v ...interface{}) {
	if *isDebug && *isLogAll {
		addr := fmt.Sprintf("[%s] ", conn.remoteAddr)
		log.Printf(addr+format, v...)
	}
}

// 发送 WebSocket 消息
func (conn *wsConn) send(msg string) error {
	conn.debugAll("Send message: %s", msg)
	_, err := conn.c.WriteString(msg)
	if err != nil {
		conn.debug("Failed to send message: %s, error: %v", msg, err)
	}
	return err
}

// 处理 WebSocket 连接
func wsHandler(c *fastws.Conn) {
	c.SetReadTimeout(timeout)
	c.SetWriteTimeout(timeout)
	conn := &wsConn{
		c:          c,
		remoteAddr: c.RemoteAddr().String(),
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer conn.c.Close()
	defer conn.debug("WebSocket connection close")
	conn.debug("WebSocket connection open")

	forward_ch, err := server_ch.Sub()
	if err != nil {
		conn.debug("Server's main channel has been killed")
		return
	}
	defer server_ch.Unsub(forward_ch)

	var pool fastjson.ParserPool
	var msg []byte
	var clientID string
	// map(int64, *acLive)
	acMap := new(sync.Map)
	var mu sync.RWMutex
	var ac *acLive
	var data []byte

	go func() {
		for {
			msg, ok := <-forward_ch
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
		if reqType != loginType && reqType != setTokenType {
			conn.debugAll("Recieve message: %s", string(msg))
		}
		mu.RLock()
		if ac == nil && reqType != heartbeatType && reqType != loginType && reqType != setClientIDType && reqType != requestForwardDataType && reqType != setTokenType && reqType != QRCodeLoginType {
			go conn.send(fmt.Sprintf(respErrJSON, reqType, quote(reqID), needLogin, quote("Need login or token")))
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
				_, _ = c.WriteString(resp)
			}()
			pool.Put(p)
		case QRCodeLoginType:
			go func() {
				resp := conn.loginWithQRCode(acMap, reqID)
				if aci, ok := acMap.Load(0); ok {
					mu.Lock()
					ac = aci.(*acLive)
					mu.Unlock()
				}
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
		case setTokenType:
			conn.debug("Client sets token")
			data = v.Get("data").MarshalTo(data[:0])
			token := new(acfundanmu.TokenInfo)
			if err := json.Unmarshal(data, token); err != nil {
				go conn.send(fmt.Sprintf(respErrJSON, reqType, quote(reqID), invalidReqData, quote(fmt.Sprintf("Failed to unmarshal data to TokenInfo: %v", err))))
				pool.Put(p)
				continue
			}
			newAC, err := acfundanmu.NewAcFunLive(acfundanmu.SetTokenInfo(token), acfundanmu.SetDanmuClient(danmuClient()))
			if err != nil {
				go conn.send(fmt.Sprintf(respErrJSON, reqType, quote(reqID), reqHandleErr, quote(fmt.Sprintf("Failed to set TokenInfo: %v", err))))
				pool.Put(p)
				continue
			}
			mu.Lock()
			ac = &acLive{
				conn: conn,
				ac:   newAC,
			}
			acMap.Store(0, ac)
			mu.Unlock()
			go conn.send(fmt.Sprintf(respNoDataJSON, reqType, quote(reqID)))
			pool.Put(p)
		case getDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				conn.debug("getDanmu: liverUID not exist or less than 1")
				go conn.send(fmt.Sprintf(respErrJSON, getDanmuType, quote(reqID), invalidReqData, quote("liverUID not exist or less than 1")))
			} else {
				go conn.getDanmu(ctx, cancel, acMap, uid, reqID)
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
		case getLiveCutInfoType:
			go func() {
				mu.RLock()
				resp := ac.getLiveCutInfo(v, reqID)
				mu.RUnlock()
				_, _ = c.WriteString(resp)
				conn.debug("Return the live cut info's response to the client")
				pool.Put(p)
			}()
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
