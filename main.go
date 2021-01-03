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
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

func main() {
	port := flag.Uint("port", 0, "WebSocket server port, default 15368")
	isDebug = flag.Bool("debug", false, "Debug")
	flag.Parse()
	if !(*port > 1023 && *port < 65536) {
		// 默认端口为15368
		*port = 15368
	}
	debug("WebSocket server port is %d", *port)

	server := &fasthttp.Server{
		Handler: fastws.Upgrade(wsHandler),
	}

	go func() {
		if err := server.ListenAndServe(fmt.Sprintf(":%d", *port)); err != nil {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-ch
	signal.Stop(ch)
	signal.Reset(os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
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
	defer debug("WebSocket connection close")

	debug("WebSocket connection open, local address: %s, remote address: %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
	var pool fastjson.ParserPool
	var msg []byte
	var err error
	// map(int64, *acLive)
	acMap := new(sync.Map)
	var ac *acLive

	for {
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
			err := send(conn, fmt.Sprintf(respErrJSON, 0, jsonParseErr, quote(fmt.Sprintf("%v, request data: %s", err, string(msg)))))
			if err != nil {
				pool.Put(p)
				break
			}
			pool.Put(p)
			continue
		}

		reqType := v.GetInt("type")
		if reqType != heartbeatType && reqType != loginType && ac == nil {
			err := send(conn, fmt.Sprintf(respErrJSON, reqType, needLogin, quote("Need login")))
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
				resp := login(acMap, account, password)
				if aci, ok := acMap.Load(0); ok {
					ac = aci.(*acLive)
				}
				_ = send(conn, resp)
			}()
			pool.Put(p)
		case getDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				debug("getDanmu: liverUID not exist or less than 1")
				_ = send(conn, fmt.Sprintf(respErrJSON, getDanmuType, invalidReqData, quote("liverUID not exist or less than 1")))
				pool.Put(p)
				break
			}
			go getDanmu(acMap, conn, uid)
			pool.Put(p)
		case stopDanmuType:
			uid := v.GetInt64("data", "liverUID")
			if uid <= 0 {
				debug("stopDanmu: liverUID not exist or less than 1")
				_ = send(conn, fmt.Sprintf(respErrJSON, stopDanmuType, invalidReqData, quote("liverUID not exist or less than 1")))
				pool.Put(p)
				break
			}
			go stopDanmu(acMap, conn, uid)
			pool.Put(p)
		default:
			if f, ok := cmdDispatch[reqType]; ok {
				go func() {
					resp := f(ac, v)
					_ = send(conn, resp)
					pool.Put(p)
				}()
			} else {
				debug("Unknown request type:%d", reqType)
				_ = send(conn, fmt.Sprintf(respErrJSON, reqType, invalidReqType, quote("Unknown request type")))
				pool.Put(p)
			}
		}
	}
}
