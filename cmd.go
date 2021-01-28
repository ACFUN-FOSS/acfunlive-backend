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
	getWatchingListType:  (*acLive).getWatchingList,
	getBillboardType:     (*acLive).getBillboard,
	getSummaryType:       (*acLive).getSummary,
	getLuckListType:      (*acLive).getLuckList,
	getPlaybackType:      (*acLive).getPlayback,
	getAllGiftListType:   (*acLive).getAllGiftList,
	getWalletBalanceType: (*acLive).getWalletBalance,
	getUserLiveInfoType:  (*acLive).getUserLiveInfo,
	getAllLiveListType:   (*acLive).getAllLiveList,
	getManagerListType:   (*acLive).getManagerList,
	addManagerType:       (*acLive).addManager,
	deleteManagerType:    (*acLive).deleteManager,
	managerKickType:      (*acLive).managerKick,
	authorKickType:       (*acLive).authorKick,
	getMedalDetailType:   (*acLive).getMedalDetail,
	getMedalListType:     (*acLive).getMedalList,
	getMedalRankListType: (*acLive).getMedalRankList,
	getUserMedalType:     (*acLive).getUserMedal,
}

// 处理登陆命令
func login(acMap *sync.Map, account, password, reqID string) string {
	var newAC *acfundanmu.AcFunLive
	var err error
	if account == "" || password == "" {
		newAC, err = acfundanmu.NewAcFunLive()
		if err != nil {
			debug("login(): cannot login as anonymous: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
	} else {
		cookies, err := acfundanmu.Login(account, password)
		if err != nil {
			debug("login(): cannot login as AcFun user: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
		newAC, err = acfundanmu.NewAcFunLive(acfundanmu.SetCookies(cookies))
		if err != nil {
			debug("login(): call Init() error: %+v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
	}

	ac := new(acLive)
	ac.ac = newAC
	acMap.Store(0, ac)

	info := ac.ac.GetTokenInfo()
	data, err := json.Marshal(info)
	if err != nil {
		debug("login(): cannot marshal to json: %+v", info)
		return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, loginType, quote(reqID), fmt.Sprintf(`{"tokenInfo":%s}`, string(data)))
}

// 获取抢到红包的用户列表
func (ac *acLive) getLuckList(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		debug("getLuckList(): No liveID")
		return fmt.Sprintf(respErrJSON, getLuckListType, quote(reqID), invalidReqData, quote("Need liveID"))
	}
	redpackID := string(v.GetStringBytes("data", "redpackID"))
	if redpackID == "" {
		debug("getLuckList(): No redpackID")
		return fmt.Sprintf(respErrJSON, getLuckListType, quote(reqID), invalidReqData, quote("Need redpackID"))
	}

	list, err := ac.ac.GetLuckList(liveID, redpackID)
	if err != nil {
		debug("getLuckList(): call GetLuckList() error: %v", err)
		return fmt.Sprintf(respErrJSON, getLuckListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}
	data, err := json.Marshal(list)
	if err != nil {
		debug("getLuckList(): cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getLuckListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getLuckListType, quote(reqID), string(data))
}

// 获取全部礼物的列表
func (ac *acLive) getAllGiftList(v *fastjson.Value, reqID string) string {
	gift, err := ac.ac.GetAllGiftList()
	if err != nil {
		debug("getAllGiftList(): call GetAllGiftList() error: %v", err)
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
		debug("getAllGiftList(): cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getAllGiftListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getAllGiftListType, quote(reqID), string(data))
}

// 获取账户钱包数据
func (ac *acLive) getWalletBalance(v *fastjson.Value, reqID string) string {
	acCoin, banana, err := ac.ac.GetWalletBalance()
	if err != nil {
		debug("getWalletBalance(): call GetWalletBalance() error: %v", err)
		return fmt.Sprintf(respErrJSON, getWalletBalanceType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getWalletBalanceType, quote(reqID), fmt.Sprintf(`{"acCoin":%d,"banana":%d}`, acCoin, banana))
}
