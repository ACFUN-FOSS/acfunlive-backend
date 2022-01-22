package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/orzogc/fastws"
	"github.com/segmentio/encoding/json"
	"github.com/thanhpk/randstr"
	"github.com/valyala/fastjson"
)

const (
	minCharNum             = 100
	maxCharNum             = 1000
	msgNum                 = 100
	heartbeatJSON          = `{"type":1}`
	setClientIDJSON        = `{"type":3,"requestID":"abc","data":{"clientID":%s}}`
	requestForwardDataJSON = `{"type":4,"requestID":"abc","data":{"clientID":%s,"message":%s}}`
)

var (
	//char_num = [...]int{100, 1000, 10000, 100000, 1000000}
	quote = strconv.Quote
)

type forwardMsg struct {
	N int    `json:"n"`
	S string `json:"s"`
}

func main() {
	conn1, err := fastws.Dial("ws://127.0.0.1:15368")
	checkErr(err)
	defer log.Println("WebSocket client1 shutdown")
	defer conn1.Close()

	go func() {
		for {
			_, err := conn1.WriteString(heartbeatJSON)
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		var msg []byte
		var err error
		for {
			_, msg, err = conn1.ReadMessage(msg[:0])
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}
		}
	}()

	_, err = conn1.WriteString(fmt.Sprintf(setClientIDJSON, quote("client1")))
	checkErr(err)

	conn2, err := fastws.Dial("ws://127.0.0.1:15368")
	checkErr(err)
	defer log.Println("WebSocket client2 shutdown")
	defer conn2.Close()

	go func() {
		for {
			_, err := conn2.WriteString(heartbeatJSON)
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()

	_, err = conn2.WriteString(fmt.Sprintf(setClientIDJSON, quote("client2")))
	checkErr(err)

	msgs := make([]forwardMsg, msgNum)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < msgNum; i++ {
		charNum := rand.Intn(maxCharNum-minCharNum) + maxCharNum
		s := randstr.String(charNum)
		msgs[i] = forwardMsg{
			N: i,
			S: s,
		}
	}

	ch := make(chan struct{}, 2)
	var mu sync.Mutex
	var failed int

	go func() {
		for i := 0; i < msgNum; i++ {
			m, err := json.Marshal(msgs[i])
			checkErr(err)
			_, err = conn1.WriteString(fmt.Sprintf(requestForwardDataJSON, quote("client2"), quote(string(m))))
			checkErr(err)
		}

		ch <- struct{}{}
	}()

	go func() {
		var pool fastjson.ParserPool
		var msg []byte
		var err error
		var wg sync.WaitGroup
		var n int
		for {
			if n == msgNum {
				break
			}

			_, msg, err = conn2.ReadMessage(msg[:0])
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}

			p := pool.Get()
			v, err := p.ParseBytes(msg)
			checkErr(err)

			respType := v.GetInt("type")
			if respType != 5 {
				continue
			}

			clientID := string(v.GetStringBytes("data", "clientID"))
			if clientID != "client1" {
				mu.Lock()
				failed++
				mu.Unlock()
				n++
				continue
			}
			m := string(v.GetStringBytes("data", "message"))

			wg.Add(1)
			n++
			go func(m string) {
				var fm forwardMsg
				err := json.Unmarshal([]byte(m), &fm)
				checkErr(err)

				if fm.S != msgs[fm.N].S {
					mu.Lock()
					failed++
					mu.Unlock()
				}

				wg.Done()
			}(m)
		}

		wg.Wait()
		ch <- struct{}{}
	}()

	<-ch
	<-ch

	log.Printf("failed: %v", failed)
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
