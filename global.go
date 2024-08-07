package main

import (
	"context"
	"strconv"
	"time"

	"github.com/orzogc/acfundanmu"
	"github.com/orzogc/fastws"
	"github.com/ugjka/messenger"
)

const (
	heartbeatJSON   = `{"type":1}`                                        // 心跳包
	respJSON        = `{"type":%d,"requestID":%s,"result":1,"data":%s}`   // 响应
	respNoDataJSON  = `{"type":%d,"requestID":%s,"result":1}`             // 没有 data 的响应
	respErrJSON     = `{"type":%d,"requestID":%s,"result":%d,"error":%s}` // 错误响应
	danmuJSON       = `{"liverUID":%d,"type":%d,"data":%s}`               // 弹幕和信号数据
	danmuNoDataJSON = `{"liverUID":%d,"type":%d}`                         // 没有数据的弹幕和信号数据
)

// 基础类型
const (
	heartbeatType = iota + 1
	loginType
	setClientIDType
	requestForwardDataType
	forwardDataType
	setTokenType
	QRCodeLoginType
	QRCodeScannedType
	QRCodeLoginCancelType
	QRCodeLoginSuccessType
)

// 命令类型
const (
	getDanmuType = iota + 100
	stopDanmuType
	getWatchingListType
	getBillboardType
	getSummaryType
	getLuckListType
	getPlaybackType
	getAllGiftListType
	getWalletBalanceType
	getUserLiveInfoType
	getAllLiveListType
	uploadImageType
	getLiveDataType
	getScheduleListType
	getGiftListType
	getUserInfoType
	getLiveCutInfoType
)

const (
	getManagerListType = iota + 200
	addManagerType
	deleteManagerType
	getAllKickHistoryType
	managerKickType
	authorKickType
)

const (
	getMedalDetailType = iota + 300
	getMedalListType
	getMedalRankListType
	getUserMedalType
	wearMedalType
	cancelWearMedalType
)

const (
	checkLiveAuthType = iota + 900
	getLiveTypeListType
	getPushConfigType
	getLiveStatusType
	getTranscodeInfoType
	startLiveType
	stopLiveType
	changeTitleAndCoverType
	getLiveCutStatusType
	setLiveCutStatusType
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
	shareLiveType
)

// 状态信号类型
const (
	danmuStopType = iota + 2000
	bananaCountType
	displayInfoType
	topUsersType
	recentCommentType
	redpackListType
	danmuStopErrType = 2999
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
	jsonParseErr   = iota + 10 // 请求的 json 解析错误
	invalidReqType             // 无效的请求 type
	invalidReqData             // 无效的请求 data
	reqHandleErr               // 处理请求时出现错误
	needLogin                  // 需要登陆
)

const timeout = 10 * time.Second
const idleTimeout = 60 * time.Second

var (
	isDebug   *bool                // 是否调试
	isTCP     *bool                // 弹幕客户端是否使用 TCP 连接
	isLogAll  *bool                // 是否记录所有调试信息
	quote     = strconv.Quote      // 给字符串加上双引号
	server_ch *messenger.Messenger // server 间通讯的 channel
)

// WebSocket 连接
type wsConn struct {
	c          *fastws.Conn
	remoteAddr string
}

// 直播相关信息
type acLive struct {
	conn   *wsConn
	ac     *acfundanmu.AcFunLive
	cancel context.CancelFunc
}

// 转发数据
type forwardMsg struct {
	requestID string
	SourceID  string `json:"clientID"`
	clientID  string
	Message   string `json:"message"`
}
