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
	getLiveDataType:         (*acLive).getLiveData,
	getGiftListType:         (*acLive).getGiftList,
	getUserInfoType:         (*acLive).getUserInfo,
	getManagerListType:      (*acLive).getManagerList,
	addManagerType:          (*acLive).addManager,
	deleteManagerType:       (*acLive).deleteManager,
	getAllKickHistoryType:   (*acLive).getAllKickHistory,
	managerKickType:         (*acLive).managerKick,
	authorKickType:          (*acLive).authorKick,
	getMedalDetailType:      (*acLive).getMedalDetail,
	getMedalListType:        (*acLive).getMedalList,
	getMedalRankListType:    (*acLive).getMedalRankList,
	getUserMedalType:        (*acLive).getUserMedal,
	wearMedalType:           (*acLive).wearMedal,
	cancelWearMedalType:     (*acLive).cancelWearMedal,
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
func (conn *wsConn) login(acMap *sync.Map, account, password, reqID string) string {
	var newAC *acfundanmu.AcFunLive
	var err error
	conn.debug("Client request login")
	if account == "" || password == "" {
		newAC, err = acfundanmu.NewAcFunLive()
		if err != nil {
			conn.debug("login() error: cannot login as anonymous: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
	} else {
		cookies, err := acfundanmu.Login(account, password)
		if err != nil {
			conn.debug("login() error: cannot login as AcFun user: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
		newAC, err = acfundanmu.NewAcFunLive(acfundanmu.SetCookies(cookies))
		if err != nil {
			conn.debug("login() error: %v", err)
			return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
		}
	}
	conn.debug("Client login is successful, uid is %d", newAC.GetUserID())

	ac := new(acLive)
	ac.conn = conn
	ac.ac = newAC
	acMap.Store(0, ac)

	info := ac.ac.GetTokenInfo()
	data, err := json.Marshal(info)
	if err != nil {
		//conn.debug("login() error: cannot marshal to json: %+v", info)
		return fmt.Sprintf(respErrJSON, loginType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, loginType, quote(reqID), fmt.Sprintf(`{"tokenInfo":%s}`, string(data)))
}

// 获取全部礼物的列表
func (ac *acLive) getAllGiftList(v *fastjson.Value, reqID string) string {
	gift, err := ac.ac.GetAllGiftList()
	if err != nil {
		ac.conn.debug("getAllGiftList() error: %v", err)
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
		ac.conn.debug("getAllGiftList() error: cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getAllGiftListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getAllGiftListType, quote(reqID), string(data))
}

// 获取账户钱包数据
func (ac *acLive) getWalletBalance(v *fastjson.Value, reqID string) string {
	acCoin, banana, err := ac.ac.GetWalletBalance()
	if err != nil {
		ac.conn.debug("getWalletBalance() error: %v", err)
		return fmt.Sprintf(respErrJSON, getWalletBalanceType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getWalletBalanceType, quote(reqID), fmt.Sprintf(`{"acCoin":%d,"banana":%d}`, acCoin, banana))
}

// 上传图片
/*
func (ac *acLive) uploadImage(v *fastjson.Value, reqID string) string {
	imageFile := string(v.GetStringBytes("data", "imageFile"))
	if imageFile == "" {
		ac.conn.debug("uploadImage() error: No imageFile")
		return fmt.Sprintf(respErrJSON, uploadImageType, quote(reqID), invalidReqData, quote("Need imageFile"))
	}

	imageURL, err := ac.ac.UploadImage(imageFile)
	if err != nil {
		ac.conn.debug("uploadImage() error: %v", err)
		return fmt.Sprintf(respErrJSON, uploadImageType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, uploadImageType, quote(reqID), fmt.Sprintf(`{"imageURL":%s}`, quote(imageURL)))
}
*/

// 获取直播间礼物列表
func (ac *acLive) getGiftList(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		ac.conn.debug("getGiftList() error: No liveID")
		return fmt.Sprintf(respErrJSON, getGiftListType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	gift, err := ac.ac.GetGiftList(liveID)
	if err != nil {
		ac.conn.debug("getGiftList() error: %v", err)
		return fmt.Sprintf(respErrJSON, getGiftListType, quote(reqID), reqHandleErr, quote(err.Error()))
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
		ac.conn.debug("getGiftList() error: cannot marshal to json: %+v", list)
		return fmt.Sprintf(respErrJSON, getGiftListType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, getGiftListType, quote(reqID), string(data))
}

// 房管踢人
func (ac *acLive) managerKick(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		ac.conn.debug("managerKick() error: No liveID")
		return fmt.Sprintf(respErrJSON, managerKickType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	kickedUID := v.GetInt64("data", "kickedUID")
	if kickedUID <= 0 {
		ac.conn.debug("managerKick() error: kickedUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, managerKickType, quote(reqID), invalidReqData, quote("kickedUID not exist or less than 1"))
	}

	err := ac.ac.ManagerKick(liveID, kickedUID)
	if err != nil {
		ac.conn.debug("managerKick() error: %v", err)
		return fmt.Sprintf(respErrJSON, managerKickType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, managerKickType, quote(reqID))
}

// 主播踢人
func (ac *acLive) authorKick(v *fastjson.Value, reqID string) string {
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		ac.conn.debug("authorKick() error: No liveID")
		return fmt.Sprintf(respErrJSON, authorKickType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	kickedUID := v.GetInt64("data", "kickedUID")
	if kickedUID <= 0 {
		ac.conn.debug("authorKick() error: kickedUID not exist or less than 1")
		return fmt.Sprintf(respErrJSON, authorKickType, quote(reqID), invalidReqData, quote("kickedUID not exist or less than 1"))
	}

	err := ac.ac.AuthorKick(liveID, kickedUID)
	if err != nil {
		ac.conn.debug("authorKick() error: %v", err)
		return fmt.Sprintf(respErrJSON, authorKickType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, authorKickType, quote(reqID))
}

// 检测是否有直播权限
func (ac *acLive) checkLiveAuth(v *fastjson.Value, reqID string) string {
	auth, err := ac.ac.CheckLiveAuth()
	if err != nil {
		ac.conn.debug("checkLiveAuth() error: %v", err)
		return fmt.Sprintf(respErrJSON, checkLiveAuthType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, checkLiveAuthType, quote(reqID), fmt.Sprintf(`{"liveAuth":%v}`, auth))
}

// 启动直播
func (ac *acLive) startLive(v *fastjson.Value, reqID string) string {
	title := string(v.GetStringBytes("data", "title"))
	if title == "" {
		ac.conn.debug("startLive() error: No title")
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), invalidReqData, quote("Need title"))
	}
	coverFile := string(v.GetStringBytes("data", "coverFile"))
	if coverFile == "" {
		ac.conn.debug("startLive() error: No coverFile")
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), invalidReqData, quote("Need coverFile"))
	}
	streamName := string(v.GetStringBytes("data", "streamName"))
	if streamName == "" {
		ac.conn.debug("startLive() error: No streamName")
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
		ac.conn.debug("startLive() error: %v", err)
		return fmt.Sprintf(respErrJSON, startLiveType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respJSON, startLiveType, quote(reqID), fmt.Sprintf(`{"liveID":%s}`, quote(liveID)))
}

// 更改直播间标题和封面
func (ac *acLive) changeTitleAndCover(v *fastjson.Value, reqID string) string {
	title := string(v.GetStringBytes("data", "title"))
	coverFile := string(v.GetStringBytes("data", "coverFile"))
	liveID := string(v.GetStringBytes("data", "liveID"))
	if liveID == "" {
		ac.conn.debug("changeTitleAndCover() error: No liveID")
		return fmt.Sprintf(respErrJSON, changeTitleAndCoverType, quote(reqID), invalidReqData, quote("Need liveID"))
	}

	err := ac.ac.ChangeTitleAndCover(title, coverFile, liveID)
	if err != nil {
		ac.conn.debug("changeTitleAndCover() error: %v", err)
		return fmt.Sprintf(respErrJSON, changeTitleAndCoverType, quote(reqID), reqHandleErr, quote(err.Error()))
	}

	return fmt.Sprintf(respNoDataJSON, changeTitleAndCoverType, quote(reqID))
}
