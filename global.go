package main

import (
	"context"
	"strconv"

	"github.com/orzogc/acfundanmu"
)

const (
	heartbeatJSON  = `{"type":1}`                                        // 心跳包
	respJSON       = `{"type":%d,"requestID":%s,"result":1,"data":%s}`   // 响应
	respNoDataJSON = `{"type":%d,"requestID":%s,"result":1}`             // 没有data的响应
	respErrJSON    = `{"type":%d,"requestID":%s,"result":%d,"error":%s}` // 错误响应
	danmuJSON      = `{"liverUID":%d,"type":%d,"data":%s}`               // 弹幕和信号数据
)

// 基础类型
const (
	heartbeatType = iota + 1
	loginType
)

// 命令类型
const (
	getDanmuType = iota + 100
	stopDanmuType
	getWatchlingListType
	getBillboardType
	getSummaryType
	getLuckListType
	getPlaybackType
	getAllGiftListType
	getWalletBalanceType
	getUserLiveInfoType
	getAllLiveListType
)

const (
	getManagerListType = iota + 200
	addManagerType
	deleteManagerType
	getKickHistoryType
	managerKickType
	authorKickType
)

const (
	getMedalDetailType = iota + 300
	getMedalListType
	getMedalRankListType
	getUserMedalType
)

const (
	checkLiveAuthType = iota + 900
	getLiveTypeListType
	getOBSConfigType
)

// 弹幕类型
const (
	commentType = iota + 1000
	likeType
	enterRoomType
	followAuthorType
	throwBananaType
	giftType
	richTextType
	joinClubType
)

// 状态信号类型
const (
	danmuStopType = iota + 2000
	bananaCountType
	displayInfoType
	topUsersType
	recentCommentType
	redpackListType
)

// 连麦类型
const (
	chatCallType = iota + 2100
	chatAcceptType
	chatReadyType
	chatEndType
)

// 通知信号类型
const (
	kickedOutType = iota + 3000
	violationAlertType
	managerStateType
)

// 富文本类型
const (
	richTextUserInfoType = iota + 1900
	richTextPlainType
	richTextImageType
)

// 错误码
const (
	jsonParseErr   = iota + 10 // 请求的json解析错误
	invalidReqType             // 无效的请求type
	invalidReqData             // 无效的请求data
	reqHandleErr               // 处理请求时出现错误
	needLogin                  // 需要登陆
)

var (
	isDebug *bool           // 是否调试
	quote   = strconv.Quote // 给字符串加上双引号
)

// 直播相关信息
type acLive struct {
	ac     *acfundanmu.AcFunLive
	cancel context.CancelFunc
}
