package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/orzogc/fastws"
	"github.com/valyala/fastjson"
)

const (
	heartbeatJSON           = `{"type":1}`
	loginJSON               = `{"type":2,"requestID":"abc","data":{"account":%s,"password":%s}}`
	getDanmuJSON            = `{"type":100,"requestID":"abc","data":{"liverUID":%d}}`
	stopDanmuJSON           = `{"type":101,"requestID":"abc","data":{"liverUID":%d}}`
	getAllKickHistoryJSON   = `{"type":203,"requestID":"abc","data":{"liveID":%s}}`
	authorKickJSON          = `{"type":205,"requestID":"abc","data":{"liveID":%s,"kickedUID":%d}}`
	checkLiveAuthJSON       = `{"type":900,"requestID":"abc"}`
	getLiveTypeListJSON     = `{"type":901,"requestID":"abc"}`
	getPushConfigJSON       = `{"type":902,"requestID":"abc"}`
	getLiveStatusJSON       = `{"type":903,"requestID":"abc"}`
	getTranscodeInfoJSON    = `{"type":904,"requestID":"abc","data":{"streamName":%s}}`
	startLiveJSON           = `{"type":905,"requestID":"abc","data":{"title":%s,"coverFile":%s,"streamName":%s,"portrait":false,"panoramic":false,"categoryID":3,"subCategoryID":399}}`
	stopLiveJSON            = `{"type":906,"requestID":"abc","data":{"liveID":%s}}`
	changeTitleAndCoverJSON = `{"type":907,"requestID":"abc","data":{"title":%s,"coverFile":%s,"liveID":%s}}`
)

var quote = strconv.Quote

func main() {
	account := flag.String("account", "", "AcFun account")
	password := flag.String("password", "", "AcFun account password")
	title1 := flag.String("title1", "", "Live title before changed")
	title2 := flag.String("title2", "", "Live title after changed")
	cover1 := flag.String("cover1", "", "Live cover before changed")
	cover2 := flag.String("cover2", "", "Live cover after changed")
	flag.Parse()

	conn, err := fastws.Dial("ws://127.0.0.1:15368")
	checkErr(err)
	defer log.Println("WebSocket client shutdown")
	defer conn.Close()

	go func() {
		for {
			_, err := conn.WriteString(heartbeatJSON)
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()

	var liveID string
	var streamName string
	var userID int64
	ch := make(chan struct{}, 1)
	go func() {
		var pool fastjson.ParserPool
		var msg []byte
		var err error
		for {
			_, msg, err = conn.ReadMessage(msg[:0])
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}
			log.Println(string(msg))

			p := pool.Get()
			v, err := p.ParseBytes(msg)
			checkErr(err)

			respType := v.GetInt("type")
			switch respType {
			case 1:
				log.Println("Receive heartbeat")
			case 2:
				if v.GetInt("result") != 1 {
					log.Println("Login failed")
					pool.Put(p)
					continue
				}
				userID = v.GetInt64("data", "tokenInfo", "userID")
				log.Printf("Login sucess, account uid: %d", userID)
				ch <- struct{}{}
			case 3:
			case 4:
			case 5:
				log.Printf("Receive broadcast from %s : %s", string(v.GetStringBytes("data", "clientID")), string(v.GetStringBytes("data", "message")))
			case 6:
			case 100:
				if v.GetInt("result") != 1 {
					log.Printf("Cannot get danmu, response: %s", string(msg))
					pool.Put(p)
					continue
				}
				liveID = string(v.GetStringBytes("data", "StreamInfo", "liveID"))
				log.Printf("LiveID: %s", liveID)
				ch <- struct{}{}
			case 101:
				if v.GetInt("result") != 1 {
					log.Printf("Cannot stop danmu, response: %s", string(msg))
					pool.Put(p)
					continue
				}
			case 102:
			case 103:
			case 104:
			case 105:
			case 106:
			case 107:
			case 108:
			case 109:
			case 110:
			case 111:
			case 112:
			case 113:
			case 114:
			case 115:
			case 116:
			case 200:
			case 201:
			case 202:
			case 203:
			case 204:
			case 205:
			case 300:
			case 301:
			case 302:
			case 303:
			case 304:
			case 305:
			case 900:
			case 901:
			case 902:
				streamName = string(v.GetStringBytes("data", "streamName"))
				log.Printf("Stream name: %s", streamName)
				log.Printf("RTMP server: %s", string(v.GetStringBytes("data", "rtmpServer")))
				log.Printf("Stream key: %s", string(v.GetStringBytes("data", "streamKey")))
			case 903:
			case 904:
			case 905:
				liveID = string(v.GetStringBytes("data", "liveID"))
				log.Printf("Live ID: %s", liveID)
			case 906:
			case 907:
			case 908:
			case 909:
			case 1000:
				v = v.Get("data")
				log.Printf("%s %d %s(%d): %s",
					string(v.GetStringBytes("danmuInfo", "userInfo", "medal", "clubName")),
					v.GetInt("danmuInfo", "userInfo", "medal", "level"),
					string(v.GetStringBytes("danmuInfo", "userInfo", "nickname")),
					v.GetInt64("danmuInfo", "userInfo", "userID"),
					string(v.GetStringBytes("content")),
				)
			case 1001:
				v = v.Get("data", "userInfo")
				log.Printf("%s %d %s(%d) like",
					string(v.GetStringBytes("medal", "clubName")),
					v.GetInt("medal", "level"),
					string(v.GetStringBytes("nickname")),
					v.GetInt64("userID"),
				)
			case 1002:
				v = v.Get("data", "userInfo")
				log.Printf("%s %d %s(%d) enter live room",
					string(v.GetStringBytes("medal", "clubName")),
					v.GetInt("medal", "level"),
					string(v.GetStringBytes("nickname")),
					v.GetInt64("userID"),
				)
			case 1003:
				v = v.Get("data", "userInfo")
				log.Printf("%s %d %s(%d) follow liver",
					string(v.GetStringBytes("medal", "clubName")),
					v.GetInt("medal", "level"),
					string(v.GetStringBytes("nickname")),
					v.GetInt64("userID"),
				)
			case 1004:
				v = v.Get("data")
				log.Printf("%s(%d) give %d bananas",
					string(v.GetStringBytes("danmuInfo", "userInfo", "nickname")),
					v.GetInt64("danmuInfo", "userInfo", "userID"),
					v.GetInt("bananaCount"),
				)
			case 1005:
				v = v.Get("data")
				log.Printf("%s %d %s(%d) give %d %s",
					string(v.GetStringBytes("danmuInfo", "userInfo", "medal", "clubName")),
					v.GetInt("danmuInfo", "userInfo", "medal", "level"),
					string(v.GetStringBytes("danmuInfo", "userInfo", "nickname")),
					v.GetInt64("danmuInfo", "userInfo", "userID"),
					v.GetInt("count"),
					string(v.GetStringBytes("giftDetail", "giftName")),
				)
			case 1006:
				log.Printf("%s", string(msg))
			case 1007:
				v = v.Get("data")
				log.Printf("%s(%d) join %s(%d) club",
					string(v.GetStringBytes("fansInfo", "nickname")),
					v.GetInt64("fansInfo", "userID"),
					string(v.GetStringBytes("uperInfo", "nickname")),
					v.GetInt64("uperInfo", "userID"),
				)
			case 1008:
			case 2000:
				log.Printf("%s", string(msg))
				ch <- struct{}{}
			case 2001:
				log.Printf("Banana count: %s", string(v.GetStringBytes("data", "bananaCount")))
			case 2002:
				v = v.Get("data")
				log.Printf("Watching count: %s, like count: %s, like delta: %d",
					string(v.GetStringBytes("watchingCount")),
					string(v.GetStringBytes("likeCount")),
					v.GetInt("likeDelta"),
				)
			case 2003:
				list := v.GetArray("data")
				log.Printf("Top users: %s(%d) %s  %s(%d) %s  %s(%d) %s",
					string(list[0].GetStringBytes("userInfo", "nickname")),
					list[0].GetInt64("userInfo", "userID"),
					string(list[0].GetStringBytes("displaySendAmount")),
					string(list[1].GetStringBytes("userInfo", "nickname")),
					list[1].GetInt64("userInfo", "userID"),
					string(list[1].GetStringBytes("displaySendAmount")),
					string(list[2].GetStringBytes("userInfo", "nickname")),
					list[2].GetInt64("userInfo", "userID"),
					string(list[2].GetStringBytes("displaySendAmount")),
				)
			case 2004:
				list := v.GetArray("data")
				for _, v := range list {
					log.Printf("%s %d %s(%d): %s",
						string(v.GetStringBytes("danmuInfo", "userInfo", "medal", "clubName")),
						v.GetInt("danmuInfo", "userInfo", "medal", "level"),
						string(v.GetStringBytes("danmuInfo", "userInfo", "nickname")),
						v.GetInt64("danmuInfo", "userInfo", "userID"),
						string(v.GetStringBytes("content")),
					)
				}
			case 2005:
				list := v.GetArray("data")
				if len(list) != 0 {
					for _, v := range list {
						log.Printf("%s(%d) send %d redpack",
							string(v.GetStringBytes("userInfo", "nickname")),
							v.GetInt64("userInfo", "userID"),
							v.GetInt("redpackAmount"),
						)
					}
				}
			case 2100:
				log.Printf("%s", string(msg))
			case 2101:
				log.Printf("%s", string(msg))
			case 2102:
				log.Printf("%s", string(msg))
			case 2103:
				log.Printf("%s", string(msg))
			case 2999:
				log.Printf("%s", string(msg))
				ch <- struct{}{}
			case 3000:
				log.Printf("%s", string(msg))
			case 3001:
				log.Printf("%s", string(msg))
			case 3002:
				log.Printf("%s", string(msg))
			default:
				log.Printf("Error: unknown response type: %d", respType)
			}

			pool.Put(p)
		}
	}()

	_, err = conn.WriteString(fmt.Sprintf(loginJSON, quote(*account), quote(*password)))
	checkErr(err)
	<-ch
	//_, err = conn.WriteString(fmt.Sprintf(getDanmuJSON, userID))
	//checkErr(err)
	//<-ch

	//time.Sleep(10 * time.Second)
	//_, err = conn.WriteString(fmt.Sprintf(stopDanmuJSON, *liverUID))
	//checkErr(err)
	//time.Sleep(10 * time.Second)

	_, err = conn.WriteString(checkLiveAuthJSON)
	checkErr(err)

	_, err = conn.WriteString(getLiveTypeListJSON)
	checkErr(err)

	_, err = conn.WriteString(getPushConfigJSON)
	checkErr(err)

	time.Sleep(time.Minute)
	_, err = conn.WriteString(fmt.Sprintf(getTranscodeInfoJSON, quote(streamName)))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(startLiveJSON, quote(*title1), quote(*cover1), quote(streamName)))
	checkErr(err)

	time.Sleep(time.Minute)
	_, err = conn.WriteString(getLiveStatusJSON)
	checkErr(err)

	//time.Sleep(10 * time.Second)
	//_, err = conn.WriteString(fmt.Sprintf(authorKickJSON, quote(liveID), 12345))
	//checkErr(err)

	//time.Sleep(10 * time.Second)
	//_, err = conn.WriteString(fmt.Sprintf(getAllKickHistoryJSON, quote(liveID)))
	//checkErr(err)

	time.Sleep(20 * time.Minute)
	_, err = conn.WriteString(fmt.Sprintf(changeTitleAndCoverJSON, quote(*title2), quote(*cover2), quote(liveID)))
	checkErr(err)

	time.Sleep(20 * time.Minute)
	_, err = conn.WriteString(fmt.Sprintf(stopLiveJSON, quote(liveID)))
	checkErr(err)
	time.Sleep(time.Minute)
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
