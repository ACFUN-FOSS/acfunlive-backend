package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrr/fastws"
	"github.com/valyala/fastjson"
)

const (
	heartbeatJSON          = `{"type":1}`
	loginJSON              = `{"type":2,"requestID":"abc","data":{"account":%s,"password":%s}}`
	setClientIDJSON        = `{"type":3,"requestID":"abc","data":{"clientID":%s}}`
	requestForwardDataJSON = `{"type":4,"requestID":"abc","data":{"clientID":%s,"message":%s}}`
	getDanmuJSON           = `{"type":100,"requestID":"abc","data":{"liverUID":%d}}`
	stopDanmuJSON          = `{"type":101,"requestID":"abc","data":{"liverUID":%d}}`
	getWatchingListJSON    = `{"type":102,"requestID":"abc","data":{"liveID":%s}}`
	getBillboardJSON       = `{"type":103,"requestID":"abc","data":{"liverUID":%d}}`
	getSummaryJSON         = `{"type":104,"requestID":"abc","data":{"liveID":%s}}`
	getLuckListJSON        = `{"type":105,"requestID":"abc","data":{"liveID":%s,"redpackID":%s}}`
	getPlaybackJSON        = `{"type":106,"requestID":"abc","data":{"liveID":%s}}`
	getAllGiftJSON         = `{"type":107,"requestID":"abc"}`
	getWalletBalanceJSON   = `{"type":108,"requestID":"abc"}`
	getUserLiveInfoJSON    = `{"type":109,"requestID":"abc","data":{"userID":%d}}`
	getAllLiveListJSON     = `{"type":110,"requestID":"abc"}`
	uploadImageJSON        = `{"type":111,"requestID":"abc","data":{"imageFile":%s}}`
	getLiveDataJSON        = `{"type":112,"requestID":"abc","data":{"days":20}}`
	getScheduleListJSON    = `{"type":113,"requestID":"abc"}`
	getManagerListJSON     = `{"type":200,"requestID":"abc"}`
	addManagerJSON         = `{"type":201,"requestID":"abc","data":{"managerUID":%d}}`
	deleteManagerJSON      = `{"type":202,"requestID":"abc","data":{"managerUID":%d}}`
	managerKickJSON        = `{"type":204,"requestID":"abc","data":{"kickedUID":%d}}`
	getMedalDetailJSON     = `{"type":300,"requestID":"abc","data":{"liverUID":%d}}`
	getMedalListJSON       = `{"type":301,"requestID":"abc","data":{"liverUID":%d}}`
	getMedalRankListJSON   = `{"type":302,"requestID":"abc","data":{"liverUID":%d}}`
	getUserMedalJSON       = `{"type":303,"requestID":"abc","data":{"userID":%d}}`
	wearMedalJSON          = `{"type":304,"requestID":"abc","data":{"liverUID":%d}}`
	cancelWearMedalJSON    = `{"type":305,"requestID":"abc"}`
)

var quote = strconv.Quote

func main() {
	account := flag.String("account", "", "AcFun account")
	password := flag.String("password", "", "AcFun account password")
	imageFile := flag.String("image", "", "Image file")
	liverUID := flag.Int64("uid", 0, "AcFun liver uid")
	flag.Parse()

	conn, err := fastws.Dial("ws://127.0.0.1:15368")
	checkErr(err)
	defer log.Println("WebSocket client shutdown")

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
				log.Printf("Login sucess, account uid: %d", v.GetInt64("data", "tokenInfo", "userID"))
				ch <- struct{}{}
			case 3:
			case 4:
			case 5:
				log.Printf("Receive broadcast from %s : %s", string(v.GetStringBytes("data", "clientID")), string(v.GetStringBytes("data", "message")))
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
			case 903:
			case 904:
			case 905:
			case 906:
			case 907:
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

	_, err = conn.WriteString(fmt.Sprintf(setClientIDJSON, quote("client1")))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(loginJSON, quote(*account), quote(*password)))
	checkErr(err)
	<-ch
	_, err = conn.WriteString(fmt.Sprintf(getDanmuJSON, *liverUID))
	checkErr(err)
	<-ch

	_, err = conn.WriteString(fmt.Sprintf(getWatchingListJSON, quote(liveID)))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getBillboardJSON, *liverUID))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getSummaryJSON, quote(liveID)))
	checkErr(err)

	//_, err = conn.WriteString(fmt.Sprintf(getLuckListJSON, quote("7McE2WZl9Xc"), quote("b-D8XOlAlxI")))
	//checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getPlaybackJSON, quote(liveID)))
	checkErr(err)

	_, err = conn.WriteString(getAllGiftJSON)
	checkErr(err)

	_, err = conn.WriteString(getWalletBalanceJSON)
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getUserLiveInfoJSON, *liverUID))
	checkErr(err)

	//_, err = conn.WriteString(getAllLiveListJSON)
	//checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(uploadImageJSON, quote(*imageFile)))
	checkErr(err)

	_, err = conn.WriteString(getLiveDataJSON)
	checkErr(err)

	_, err = conn.WriteString(getScheduleListJSON)
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(addManagerJSON, *liverUID))
	checkErr(err)
	time.Sleep(2 * time.Second)

	_, err = conn.WriteString(getManagerListJSON)
	checkErr(err)
	time.Sleep(2 * time.Second)

	_, err = conn.WriteString(fmt.Sprintf(deleteManagerJSON, *liverUID))
	checkErr(err)

	//_, err = conn.WriteString(fmt.Sprintf(managerKickJSON, *liverUID))
	//checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getMedalDetailJSON, *liverUID))
	checkErr(err)

	//_, err = conn.WriteString(fmt.Sprintf(getMedalListJSON, *liverUID))
	//checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getMedalRankListJSON, *liverUID))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(getUserMedalJSON, *liverUID))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(wearMedalJSON, *liverUID))
	checkErr(err)

	_, err = conn.WriteString(cancelWearMedalJSON)
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(requestForwardDataJSON, quote(""), quote("hello all")))
	checkErr(err)

	_, err = conn.WriteString(fmt.Sprintf(requestForwardDataJSON, quote("client2"), quote("hello client2")))
	checkErr(err)

	time.Sleep(10 * time.Second)
	_, err = conn.WriteString(fmt.Sprintf(stopDanmuJSON, *liverUID))
	checkErr(err)
	time.Sleep(10 * time.Second)

	_ = conn.Close()
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
