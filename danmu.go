package main

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/dgrr/fastws"
	"github.com/orzogc/acfundanmu"
	"github.com/segmentio/encoding/json"
)

// 获取弹幕
func getDanmu(acMap *sync.Map, conn *fastws.Conn, uid int64, reqID string) {
	if _, ok := acMap.Load(uid); ok {
		return
	}
	aci, _ := acMap.Load(0)

	newAC, err := aci.(*acLive).ac.SetLiverUID(uid)
	if err != nil {
		debug("getDanmu(): call ReInit() error: %v", err)
		_ = send(conn, fmt.Sprintf(respErrJSON, getDanmuType, quote(reqID), reqHandleErr, quote(err.Error())))
		return
	}
	ac := new(acLive)
	ac.ac = newAC
	info := ac.ac.GetStreamInfo()
	data, err := json.Marshal(info)
	if err != nil {
		debug("getDanmu(): cannot marshal to json: %v", err)
		_ = send(conn, fmt.Sprintf(respErrJSON, getDanmuType, quote(reqID), reqHandleErr, quote(err.Error())))
		return
	}
	err = send(conn, fmt.Sprintf(respJSON, getDanmuType, quote(reqID), fmt.Sprintf(`{"StreamInfo":%s}`, string(data))))
	if err != nil {
		return
	}

	errCh := make(chan error, 100)

	ac.ac.OnComment(func(ac *acfundanmu.AcFunLive, d *acfundanmu.Comment) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnComment(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, commentType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnLike(func(ac *acfundanmu.AcFunLive, d *acfundanmu.Like) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnLike(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, likeType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnEnterRoom(func(ac *acfundanmu.AcFunLive, d *acfundanmu.EnterRoom) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnEnterRoom(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, enterRoomType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnFollowAuthor(func(ac *acfundanmu.AcFunLive, d *acfundanmu.FollowAuthor) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnFollowAuthor(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, followAuthorType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnThrowBanana(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ThrowBanana) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnThrowBanana(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, throwBananaType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnGift(func(ac *acfundanmu.AcFunLive, d *acfundanmu.Gift) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnGift(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, giftType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnRichText(func(ac *acfundanmu.AcFunLive, d *acfundanmu.RichText) {
		data := `{"sendTime":%d,"segments":[%s]}`
		richText := `{"type":%d,"segment":%s}`
		s := make([]string, len(d.Segments))
		for i, r := range d.Segments {
			switch r := r.(type) {
			case *acfundanmu.RichTextUserInfo:
				t, err := json.Marshal(r)
				if err != nil {
					debug("OnRichText(): cannot marshal to json: %+v", r)
					return
				}
				s[i] = fmt.Sprintf(richText, richTextUserInfoType, string(t))
			case *acfundanmu.RichTextPlain:
				t, err := json.Marshal(r)
				if err != nil {
					debug("OnRichText(): cannot marshal to json: %+v", r)
					return
				}
				s[i] = fmt.Sprintf(richText, richTextPlainType, string(t))
			case *acfundanmu.RichTextImage:
				t, err := json.Marshal(r)
				if err != nil {
					debug("OnRichText(): cannot marshal to json: %+v", r)
					return
				}
				s[i] = fmt.Sprintf(richText, richTextImageType, string(t))
			}
		}
		data = fmt.Sprintf(data, d.SendTime, strings.Join(s, `,`))
		err := send(conn, fmt.Sprintf(danmuJSON, uid, richTextType, data))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnJoinClub(func(ac *acfundanmu.AcFunLive, d *acfundanmu.JoinClub) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnJoinClub(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, joinClubType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnDanmuStop(func(ac *acfundanmu.AcFunLive, err error) {
		var msg string
		if err == nil {
			msg = `{"liveClosed":true,"reason":""}`
		} else {
			msg = fmt.Sprintf(`{"liveClosed":false,"reason":%s}`, quote(err.Error()))
		}
		e := send(conn, fmt.Sprintf(danmuJSON, uid, danmuStopType, quote(msg)))
		if e != nil {
			errCh <- e
		}
	})

	ac.ac.OnBananaCount(func(ac *acfundanmu.AcFunLive, allBananaCount string) {
		data := fmt.Sprintf(`{"bananaCount":%s}`, quote(allBananaCount))
		err := send(conn, fmt.Sprintf(danmuJSON, uid, bananaCountType, data))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnDisplayInfo(func(ac *acfundanmu.AcFunLive, d *acfundanmu.DisplayInfo) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnDisplayInfo(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, displayInfoType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnTopUsers(func(ac *acfundanmu.AcFunLive, d []acfundanmu.TopUser) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnTopUsers(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, topUsersType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnRecentComment(func(ac *acfundanmu.AcFunLive, d []acfundanmu.Comment) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnRecentComment(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, recentCommentType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnRedpackList(func(ac *acfundanmu.AcFunLive, d []acfundanmu.Redpack) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnRedpackList(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, redpackListType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnChatCall(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ChatCall) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnChatCall(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, chatCallType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnChatAccept(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ChatAccept) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnChatAccept(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, chatAcceptType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnChatReady(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ChatReady) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnChatReady(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, chatReadyType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnChatEnd(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ChatEnd) {
		data, err := json.Marshal(d)
		if err != nil {
			debug("OnChatEnd(): cannot marshal to json: %+v", d)
			return
		}
		err = send(conn, fmt.Sprintf(danmuJSON, uid, chatEndType, string(data)))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnKickedOut(func(ac *acfundanmu.AcFunLive, kickedOutReason string) {
		data := fmt.Sprintf(`{"kickedOutReason":%s}`, quote(kickedOutReason))
		err := send(conn, fmt.Sprintf(danmuJSON, uid, kickedOutType, data))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnViolationAlert(func(ac *acfundanmu.AcFunLive, violationContent string) {
		data := fmt.Sprintf(`{"violationContent":%s}`, quote(violationContent))
		err := send(conn, fmt.Sprintf(danmuJSON, uid, violationAlertType, data))
		if err != nil {
			errCh <- err
		}
	})

	ac.ac.OnManagerState(func(ac *acfundanmu.AcFunLive, d acfundanmu.ManagerState) {
		data := fmt.Sprintf(`{"managerState":%d}`, d)
		err := send(conn, fmt.Sprintf(danmuJSON, uid, managerStateType, data))
		if err != nil {
			errCh <- err
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ac.cancel = cancel
	acMap.Store(uid, ac)
	defer acMap.Delete(uid)
	danmuCh := ac.ac.StartDanmu(ctx, true)
	debug("Start getting liver(%d) danmu", uid)
	defer debug("Stop getting liver(%d) danmu", uid)
	select {
	case <-danmuCh:
	case <-errCh:
	}
}

// 停止获取弹幕
func stopDanmu(acMap *sync.Map, conn *fastws.Conn, uid int64, reqID string) {
	aci, ok := acMap.Load(uid)
	if !ok {
		debug("Not getting liver(%d) danmu", uid)
		_ = send(conn, fmt.Sprintf(respErrJSON, stopDanmuType, quote(reqID), reqHandleErr, quote(fmt.Sprintf("Not getting liver(%d) danmu", uid))))
		return
	}
	ac := aci.(*acLive)
	ac.cancel()
	_ = send(conn, fmt.Sprintf(respNoDataJSON, stopDanmuType, quote(reqID)))
}
