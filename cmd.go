package main

import (
	"fmt"
	"sort"
	"sync"

	"github.com/orzogc/acfundanmu"
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fastjson"
)

var cmdDispatch = map[int]func(*acLive, *fastjson.Value, string) string{
	getWatchingListType:     (*acLive).getWatchingList,
	getBillboardType:        (*acLive).getBillboard,
	getSummaryType:          (*acLive).getSummary,
	getLuckListType:         (*acLive).getLuckList,
	getPlaybackType:         (*acLive).getPlayback,
	getAllGiftListType:      (*acLive).getAllGiftList,
	getWalletBalanceType:    (*acLive).getWalletBalance,
	getUserLiveInfoType:     (*acLive).getUserLiveInfo,
	getAllLiveListType:      (*acLive).getAllLiveList,
	getManagerListType:      (*acLive).getManagerList,
	addManagerType:          (*acLive).addManager,
	deleteManagerType:       (*acLive).deleteManager,
	managerKickType:         (*acLive).managerKick,
	authorKickType:          (*acLive).authorKick,
	getMedalDetailType:      (*acLive).getMedalDetail,
	getMedalListType:        (*acLive).getMedalList,
	getMedalRankListType:    (*acLive).getMedalRankList,
	getUserMedalType:        (*acLive).getUserMedal,
	checkLiveAuthType:       (*acLive).checkLiveAuth,
	getLiveTypeListType:     (*acLive).getLiveTypeList,
	getPushConfigType:       (*acLive).getPushConfig,
	getLiveStatusType:       (*acLive).getLiveStatus,
	getTranscodeInfoType:    (*acLive).getTranscodeInfo,
	startLiveType:           (*acLive).startLive,
	stopLiveType:            (*acLive).stopLive,
	changeTitleAndCoverType: (*acLive).changeTitleAndCover,
}

// 处理登陆命令
func login(acMap *sync.Map, account, password, reqID string) string {
	var newAC *acfundanmu.AcFunLive
	var err error
	if account == "" || password == "" {
		newAC, err = acfundanmu.NewAcFunLive()
		if err != nil {
			debug("login() error: cannot login as anonymous: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
	} else {
		cookies, err := acfundanmu.Login(account, password)
		if err != nil {
			debug("login() error: cannot login as AcFun user: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
		newAC, err = acfundanmu.NewAcFunLive(acfundanmu.SetCookies(cookies))
		if err != nil {
			debug("login() error: %+v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
	}

	ac := new(acLive)
	ac.ac = newAC
	acMap.Store(0, ac)

	info := ac.ac.GetTokenInfo()
	data, err := json.Marshal(info)
	if err != nil {
		debug("login() error: cannot marshal to json: %+v", info)
		return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, loginType, quote(reqID), fmt.Sprintf(`{"tokenInfo":%s}`, string(data)))
}

// 获取全部礼物的列表
func (ac *acLive) getAllGiftList(v *fastjson.Value, reqID string) string {
	gift, err := ac.ac.GetAllGiftList()
	if err != nil {
		debug("getAllGiftList() error: %v", err)
		return fmt.Sprintf(respErrJSON, getAllGiftListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}
	list := make([]acfundanmu.GiftDetail, 0, len(gift))
	for _, g := range gift {
		list = append(list, g)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].GiftID < list[j].GiftID
	})
	data, err := json.Marshal(list)
	if err != nil {
		debug("getAllGiftList() error: cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getAllGiftListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getAllGiftListType, quote(reqID), string(data))
}

// 获取账户钱包数据
func (ac *acLive) getWalletBalance(v *fastjson.Value, reqID string) string {
	acCoin, banana, err := ac.ac.GetWalletBalance()
	if err != nil {
		debug("getWalletBalance() error: %v", err)
		return fmt.Sprintf(respErrJSON, getWalletBalanceType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getWalletBalanceType, quote(reqID), fmt.Sprintf(`{"acCoin":%d,"banana":%d}`, acCoin, banana))
}

// 检测是否有直播权限
func (ac *acLive) checkLiveAuth(v *fastjson.Value, reqID string) string {
	auth, err := ac.ac.CheckLiveAuth()
	if err != nil {
		debug("checkLiveAuth() error: %v", err)
		return fmt.Sprintf(respErrJSON, checkLiveAuthType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, checkLiveAuthType, quote(reqID), fmt.Sprintf(`{"liveAuth":%v}`, auth))
}

// 启动直播
func (ac *acLive) startLive(v *fastjson.Value, reqID string) string {
	title := string(v.GetStringBytes("data", "title"))
	if title == "" {
		debug("startLive() error: No title")
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), invalidReqData, quote("Need title"))
	}
	coverFile := string(v.GetStringBytes("data", "coverFile"))
	if coverFile == "" {
		debug("startLive() error: No coverFile")
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), invalidReqData, quote("Need coverFile"))
	}
	streamName := string(v.GetStringBytes("data", "streamName"))
	if streamName == "" {
		debug("startLive() error: No streamName")
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), invalidReqData, quote("Need streamName"))
	}
	portrait := v.GetBool("data", "portrait")
	panoramic := v.GetBool("data", "panoramic")
	categoryID := v.GetInt("data", "categoryID")
	subCategoryID := v.GetInt("data", "subCategoryID")

	liveID, err := ac.ac.StartLive(title, coverFile, streamName, portrait, panoramic,
		&acfundanmu.LiveType{
			CategoryID:    categoryID,
			SubCategoryID: subCategoryID,
		})
	if err != nil {
		debug("startLive() error: %v", err)
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, startLiveType, quote(reqID), fmt.Sprintf(`{"liveID":%s}`, quote(liveID)))
}

// 更改直播间标题和封面
func (ac *acLive) changeTitleAndCover(v *fastjson.Value, reqID string) string {
	title := string(v.GetStringBytes("data", "title"))
	if title == "" {
		debug("changeTitleAndCover() error: No title")
		return fmt.Sprintf(respErrJSON, changeTitleAndCoverType, quote(reqID), invalidReqData, quote("Need title"))
	}
	coverFile := string(v.GetStringBytes("data", "coverFile"))
	if coverFile == "" {
		debug("changeTitleAndCover() error: No coverFile")
		return fmt.Sprintf(respErrJSON, changeTitleAndCoverType, quote(reqID), invalidReqData, quote("Need coverFile"))
	}
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		debug("changeTitleAndCover() error: No liveID")
		return fmt.Sprintf(respErrJSON, changeTitleAndCoverType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	err := ac.ac.ChangeTitleAndCover(title, coverFile, liveID)
	if err != nil {
		debug("changeTitleAndCover() error: %v", err)
		return fmt.Sprintf(respErrJSON, changeTitleAndCoverType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, changeTitleAndCoverType, quote(reqID))
}
