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
	isDebug = flag.Bool("debug", false, "Debug")
	flag.Parse()
	if !(*port > 1023 && *port < 65536) {
		// 默认端口为15368
		*port = 15368
	}
	debug("WebSocket server port is %d", *port)

	server_ch = messenger.New(100, true)

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

// 发送WebSocket消息
func send(conn *fastws.Conn, msg string) error {
	debug("Send message: %s", msg)
	_, err := conn.WriteString(msg)
	if err != nil {
		debug("Failed to send message: %s, error: %v", msg, err)
	}
	return err
}

// 处理WebSocket连接
func wsHandler(conn *fastws.Conn) {
	defer debug("WebSocket connection from %s close", conn.RemoteAddr().String())

	debug("WebSocket connection open, local address: %s, remote address: %s", conn.LocalAddr().String(), conn.RemoteAddr().String())

	ch, err := server_ch.Sub()
	if err != nil {
		debug("Server's main channel has been killed")
		return
	}
	defer server_ch.Unsub(ch)

	var pool fastjson.ParserPool
	var msg []byte
	var clientID string
	// map(int64, *acLive)
	acMap := new(sync.Map)
	var ac *acLive

	for {
		select {
		case msg, ok := <-ch:
			if ok {
				msg := msg.(*forwardMsg)
				if msg.clientID == "" || msg.clientID == clientID {
					data, err := json.Marshal(msg)
					if err != nil {
						debug("Forward message error: cannot marshal to json: %+v", msg)
						_ = send(conn, fmt.Sprintf(respErrJSON, forwardDataType, quote(msg.requestID), reqHandleErr, quote(err.Error())))
					}
					_ = send(conn, fmt.Sprintf(respJSON, forwardDataType, quote(msg.requestID), string(data)))
				}
			}
		default:
		}

		_, msg, err = conn.ReadMessage(msg[:0])
		if err != nil {
			if !errors.Is(err, fastws.EOF) {
				debug("WebSocket error: %v", err)
			}
			break
		}
		debug("Recieve message: %s", string(msg))

		p := pool.Get()
		v, err := p.ParseBytes(msg)
		if err != nil {
			debug("Parsing json error: %v, request data: %s", err, string(msg))
			err := send(conn, fmt.Sprintf(respErrJSON, 0, "", jsonParseErr, quote(fmt.Sprintf("error: %v, request data: %s", err, string(msg)))))
			if err != nil {
				pool.Put(p)
				break
			}
			pool.Put(p)
			continue
		}

		reqType := v.GetInt("type")
		reqID := string(v.GetStringBytes("requestID"))
		if reqType != heartbeatType && reqType != loginType && reqType != setClientIDType && reqType != requestForwardDataType && ac == nil {
			err := send(conn, fmt.Sprintf(respErrJSON, reqType, quote(reqID), needLogin, quote("Need login")))
			if err != nil {
				pool.Put(p)
				break
			}
			pool.Put(p)
			continue
		}

		switch reqType {
		case heartbeatType:
			go func() {
				_ = send(conn, heartbeatJSON)
			}()
			pool.Put(p)
		case loginType:
			account := string(v.GetStringBytes("data", "account"))
			password := string(v.GetStringBytes("data", "password"))
			go func() {
				resp := login(acMap, account, password, reqID)
				if aci, ok := acMap.Load(0); ok {
					ac = aci.(*acLive)
				}
				_ = send(conn, resp)
			}()
			pool.Put(p)
		case setClientIDType:
			clientID = string(v.GetStringBytes("data", "clientID"))
			_ = send(conn, fmt.Sprintf(respNoDataJSON, setClientIDType, quote(reqID)))
			pool.Put(p)
		case requestForwardDataType:
			msg := new(forwardMsg)
			msg.requestID = reqID
			msg.SourceID = clientID
			msg.clientID = string(v.GetStringBytes("data", "clientID"))
			msg.Message = string(v.GetStringBytes("data", "message"))
			server_ch.Broadcast(msg)
			_ = send(conn, fmt.Sprintf(respNoDataJSON, requestForwardDataType, quote(reqID)))
			pool.Put(p)
		case getDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				debug("getDanmu: liverUID not exist or less than 1")
				_ = send(conn, fmt.Sprintf(respErrJSON, getDanmuType, quote(reqID), invalidReqData, quote("liverUID not exist or less than 1")))
				pool.Put(p)
				break
			}
			go getDanmu(acMap, conn, uid, reqID)
			pool.Put(p)
		case stopDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				debug("stopDanmu: liverUID not exist or less than 1")
				_ = send(conn, fmt.Sprintf(respErrJSON, stopDanmuType, quote(reqID), invalidReqData, quote("liverUID not exist or less than 1")))
				pool.Put(p)
				break
			}
			go stopDanmu(acMap, conn, uid, reqID)
			pool.Put(p)
		default:
			if f, ok := cmdDispatch[reqType]; ok {
				go func() {
					resp := f(ac, v, reqID)
					_ = send(conn, resp)
					pool.Put(p)
				}()
			} else {
				debug("Error: unknown request type: %d", reqType)
				_ = send(conn, fmt.Sprintf(respErrJSON, reqType, quote(reqID), invalidReqType, quote("Unknown request type")))
				pool.Put(p)
			}
		}
	}
}
