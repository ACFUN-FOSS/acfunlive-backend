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
	getWatchlingListType: (*acLive).getWatchingList,
	getBillboardType:     (*acLive).getBillboard,
	getSummaryType:       (*acLive).getSummary,
	getLuckListType:      (*acLive).getLuckList,
	getPlaybackType:      (*acLive).getPlayback,
	getAllGiftListType:   (*acLive).getAllGiftList,
	getWalletBalanceType: (*acLive).getWalletBalance,
	getManagerListType:   (*acLive).getManagerList,
	addManagerType:       (*acLive).addManager,
	deleteManagerType:    (*acLive).deleteManager,
	managerKickType:      (*acLive).managerKick,
	authorKickType:       (*acLive).authorKick,
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

// 获取直播间观众列表
func (ac *acLive) getWatchingList(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		debug("getWatchingList(): No liveID")
		return fmt.Sprintf(respErrJSON, getWatchlingListType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	list, err := ac.ac.GetWatchingListWithLiveID(liveID)
	if err != nil {
		debug("getWatchingList(): call GetWatchingListWithLiveID() error: %v", err)
		return fmt.Sprintf(respErrJSON, getWatchlingListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}
	data, err := json.Marshal(list)
	if err != nil {
		debug("getWatchingList(): cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getWatchlingListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getWatchlingListType, quote(reqID), string(data))
}

// 获取礼物贡献榜
func (ac *acLive) getBillboard(v *fastjson.Value, reqID string) string {
	uid := v.GetInt64("data", "liverUID")
	if uid <= 0 {
		debug("getBillboard(): liverUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, getBillboardType, quote(reqID), invalidReqData, quote("liverUID not exist or less than 1"))
	}

	list, err := ac.ac.GetBillboard(uid)
	if err != nil {
		debug("getBillboard(): call GetBillboard() error: %v", err)
		return fmt.Sprintf(respErrJSON, getBillboardType, quote(reqID), reqHandleErr, quote(err.Error()))
	}
	data, err := json.Marshal(list)
	if err != nil {
		debug("getBillboard(): cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getBillboardType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getBillboardType, quote(reqID), string(data))
}

// 获取直播总结信息
func (ac *acLive) getSummary(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		debug("getSummary(): No liveID")
		return fmt.Sprintf(respErrJSON, getSummaryType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	info, err := ac.ac.GetSummaryWithLiveID(liveID)
	if err != nil {
		debug("getSummary(): call GetSummaryWithLiveID() error: %v", err)
		return fmt.Sprintf(respErrJSON, getSummaryType, quote(reqID), reqHandleErr, quote(err.Error()))
	}
	data, err := json.Marshal(info)
	if err != nil {
		debug("getSummary(): cannot marshal to json: %+v", info)
		return fmt.Sprintf(respErrJSON, getSummaryType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getSummaryType, quote(reqID), string(data))
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

// 获取直播回放
func (ac *acLive) getPlayback(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		debug("getPlayback(): No liveID")
		return fmt.Sprintf(respErrJSON, getPlaybackType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	info, err := ac.ac.GetPlayback(liveID)
	if err != nil {
		debug("getPlayback(): call GetPlayback() error: %v", err)
		return fmt.Sprintf(respErrJSON, getPlaybackType, quote(reqID), reqHandleErr, quote(err.Error()))
	}
	data, err := json.Marshal(info)
	if err != nil {
		debug("getPlayback(): cannot marshal to json: %+v", info)
		return fmt.Sprintf(respErrJSON, getPlaybackType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getPlaybackType, quote(reqID), string(data))
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

func (ac *acLive) getMedalInfo(v *fastjson.Value) {

}

// 获取主播的房管列表
func (ac *acLive) getManagerList(v *fastjson.Value, reqID string) string {
	list, err := ac.ac.GetManagerList()
	if err != nil {
		debug("getManagerList(): call GetManagerList() error: %v", err)
		return fmt.Sprintf(respErrJSON, getManagerListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	data, err := json.Marshal(list)
	if err != nil {
		debug("getManagerList(): cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getManagerListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getManagerListType, quote(reqID), string(data))
}

// 添加房管
func (ac *acLive) addManager(v *fastjson.Value, reqID string) string {
	uid := v.GetInt64("data", "managerUID")
	if uid <= 0 {
		debug("addManager(): managerUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, addManagerType, quote(reqID), invalidReqData, quote("managerUID not exist or less than 1"))
	}

	err := ac.ac.AddManager(uid)
	if err != nil {
		debug("addManager(): call AddManager() error: %v", err)
		return fmt.Sprintf(respErrJSON, addManagerType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, addManagerType, quote(reqID))
}

// 删除房管
func (ac *acLive) deleteManager(v *fastjson.Value, reqID string) string {
	uid := v.GetInt64("data", "managerUID")
	if uid <= 0 {
		debug("deleteManager(): managerUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, deleteManagerType, quote(reqID), invalidReqData, quote("managerUID not exist or less than 1"))
	}

	err := ac.ac.DeleteManager(uid)
	if err != nil {
		debug("deleteManager(): call DeleteManager() error: %v", err)
		return fmt.Sprintf(respErrJSON, deleteManagerType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, deleteManagerType, quote(reqID))
}

// 房管踢人
func (ac *acLive) managerKick(v *fastjson.Value, reqID string) string {
	uid := v.GetInt64("data", "kickedUID")
	if uid <= 0 {
		debug("managerKick(): kickedUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, managerKickType, quote(reqID), invalidReqData, quote("kickedUID not exist or less than 1"))
	}

	err := ac.ac.ManagerKick(uid)
	if err != nil {
		debug("managerKick(): call ManagerKick() error: %v", err)
		return fmt.Sprintf(respErrJSON, managerKickType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, managerKickType, quote(reqID))
}

// 主播踢人
func (ac *acLive) authorKick(v *fastjson.Value, reqID string) string {
	uid := v.GetInt64("data", "kickedUID")
	if uid <= 0 {
		debug("authorKick(): kickedUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, authorKickType, quote(reqID), invalidReqData, quote("kickedUID not exist or less than 1"))
	}

	err := ac.ac.AuthorKick(uid)
	if err != nil {
		debug("authorKick(): call AuthorKick() error: %v", err)
		return fmt.Sprintf(respErrJSON, authorKickType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, authorKickType, quote(reqID))
}
