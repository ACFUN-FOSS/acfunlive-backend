package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

const header = `// Code generated by github.com/ACFUN-FOSS/acfunlive-backend/cmd . DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/orzogc/acfundanmu"
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fastjson"
)
`

const funcHeader = `
func (ac *acLive) __FUNC__(v *fastjson.Value, reqID string) string {`

const singleString = `
	__PARAM__ := string(v.GetStringBytes("data", "__PARAM__"))
	if __PARAM__ == "" {
		ac.conn.debug("__FUNC__() error: No __PARAM__")
		return fmt.Sprintf(respErrJSON, __FUNCTYPE__, quote(reqID), invalidReqData, quote("Need __PARAM__"))
	}
`

const singleInt64 = `
	__PARAM__ := v.GetInt64("data", "__PARAM__")
	if __PARAM__ <= 0 {
		ac.conn.debug("__FUNC__() error: __PARAM__ not exist or less than 1")
		return fmt.Sprintf(respErrJSON, __FUNCTYPE__, quote(reqID), invalidReqData, quote("__PARAM__ not exist or less than 1"))
	}
`

const singleInt = `
	__PARAM__ := v.GetInt("data", "__PARAM__")
	if __PARAM__ <= 0 {
		ac.conn.debug("__FUNC__() error: __PARAM__ not exist or less than 1")
		return fmt.Sprintf(respErrJSON, __FUNCTYPE__, quote(reqID), invalidReqData, quote("__PARAM__ not exist or less than 1"))
	}
`

const callNoParamFunc = `
	ret, err := __CALLFUNC__()` + callFuncErrHandle

const callNoParamNoDataFunc = `
	err := __CALLFUNC__()` + callFuncErrHandle

const callFunc = `
	ret, err := __CALLFUNC__(__ALLPARAM__)` + callFuncErrHandle

const callNoDataFunc = `
	err := __CALLFUNC__(__ALLPARAM__)` + callFuncErrHandle

const callFuncErrHandle = `
	if err != nil {
		ac.conn.debug("__FUNC__() error: %v", err)
		return fmt.Sprintf(respErrJSON, __FUNCTYPE__, quote(reqID), reqHandleErr, quote(err.Error()))
	}
`

const jsonMarshal = `
	data, err := json.Marshal(ret)
	if err != nil {
		ac.conn.debug("__FUNC__() error: cannot marshal to json: %+v", ret)
		return fmt.Sprintf(respErrJSON, __FUNCTYPE__, quote(reqID), reqHandleErr, quote(err.Error()))
	}
`

const dataReturn = `
	return fmt.Sprintf(respJSON, __FUNCTYPE__, quote(reqID), string(data))
}
`

const noDataReturn = `
	return fmt.Sprintf(respNoDataJSON, __FUNCTYPE__, quote(reqID))
}
`

const noParamFunc = funcHeader + callNoParamFunc + jsonMarshal + dataReturn

const noParamNodataFunc = funcHeader + callNoParamNoDataFunc + noDataReturn

// 函数内容
type funcContent struct {
	funcName string   // 函数名字
	funcType string   // 函数类型
	callFunc string   // 调用的函数
	params   []string // 调用函数的参数
}

// 调用的函数没有参数，有返回
var noParamFuncContent = []funcContent{
	{"getAllLiveList", "getAllLiveListType", "ac.ac.GetAllLiveList", []string{}},
	{"getScheduleList", "getScheduleListType", "ac.ac.GetScheduleList", []string{}},
	{"getManagerList", "getManagerListType", "ac.ac.GetManagerList", []string{}},
	{"getMedalList", "getMedalListType", "ac.ac.GetMedalList", []string{}},
	{"getLiveTypeList", "getLiveTypeListType", "ac.ac.GetLiveTypeList", []string{}},
	{"getPushConfig", "getPushConfigType", "ac.ac.GetPushConfig", []string{}},
	{"getLiveStatus", "getLiveStatusType", "ac.ac.GetLiveStatus", []string{}},
}

// 调用的函数没有参数，没有返回
var noParamNoDataFuncContent = []funcContent{
	{"cancelWearMedal", "cancelWearMedalType", "ac.ac.CancelWearMedal", []string{}},
}

// 调用的函数只有string参数，有返回
var stringFuncContent = []funcContent{
	{"getWatchingList", "getWatchingListType", "ac.ac.GetWatchingList", []string{"liveID"}},
	{"getSummary", "getSummaryType", "ac.ac.GetSummary", []string{"liveID"}},
	{"getLuckList", "getLuckListType", "ac.ac.GetLuckList", []string{"liveID", "redpackID", "redpackBizUnit"}},
	{"getPlayback", "getPlaybackType", "ac.ac.GetPlayback", []string{"liveID"}},
	{"getAllKickHistory", "getAllKickHistoryType", "ac.ac.GetAllKickHistory", []string{"liveID"}},
	{"getTranscodeInfo", "getTranscodeInfoType", "ac.ac.GetTranscodeInfo", []string{"streamName"}},
	{"stopLive", "stopLiveType", "ac.ac.StopLive", []string{"liveID"}},
}

// 调用的函数只有int64参数，有返回
var int64FuncContent = []funcContent{
	{"getBillboard", "getBillboardType", "ac.ac.GetBillboard", []string{"liverUID"}},
	{"getUserLiveInfo", "getUserLiveInfoType", "ac.ac.GetUserLiveInfo", []string{"userID"}},
	{"getMedalDetail", "getMedalDetailType", "ac.ac.GetMedalDetail", []string{"liverUID"}},
	{"getMedalRankList", "getMedalRankListType", "ac.ac.GetMedalRankList", []string{"liverUID"}},
	{"getUserMedal", "getUserMedalType", "acfundanmu.GetUserMedal", []string{"userID"}},
}

// 调用的函数只有int64参数，没有返回
var int64NoDataFuncContent = []funcContent{
	{"addManager", "addManagerType", "ac.ac.AddManager", []string{"managerUID"}},
	{"deleteManager", "deleteManagerType", "ac.ac.DeleteManager", []string{"managerUID"}},
	{"wearMedal", "wearMedalType", "ac.ac.WearMedal", []string{"liverUID"}},
}

// 调用的函数只有int参数，有返回
var intFuncContent = []funcContent{
	{"getLiveData", "getLiveDataType", "ac.ac.GetLiveData", []string{"days"}},
}

func main() {
	output := flag.String("o", "", "output Go file")
	flag.Parse()
	if *output == "" {
		log.Panicln("Need -o to specify output Go file")
	}

	file, err := os.Create(*output)
	if err != nil {
		log.Panicf("Cannot create %s", *output)
	}
	defer file.Close()

	_, err = file.WriteString(header)
	if err != nil {
		log.Panicf("Cannot write content to %s", *output)
	}
	for _, c := range noParamFuncContent {
		f := strings.ReplaceAll(noParamFunc, "__FUNC__", c.funcName)
		f = strings.ReplaceAll(f, "__FUNCTYPE__", c.funcType)
		f = strings.ReplaceAll(f, "__CALLFUNC__", c.callFunc)
		_, err = file.WriteString(f)
		if err != nil {
			log.Panicf("Cannot write content to %s", *output)
		}
	}
	for _, c := range noParamNoDataFuncContent {
		f := strings.ReplaceAll(noParamNodataFunc, "__FUNC__", c.funcName)
		f = strings.ReplaceAll(f, "__FUNCTYPE__", c.funcType)
		f = strings.ReplaceAll(f, "__CALLFUNC__", c.callFunc)
		_, err = file.WriteString(f)
		if err != nil {
			log.Panicf("Cannot write content to %s", *output)
		}
	}
	for _, c := range stringFuncContent {
		stringFunc := funcHeader
		for _, s := range c.params {
			stringFunc += singleString
			stringFunc = strings.ReplaceAll(stringFunc, "__PARAM__", s)
		}
		stringFunc += callFunc + jsonMarshal + dataReturn
		stringFunc = strings.ReplaceAll(stringFunc, "__FUNC__", c.funcName)
		stringFunc = strings.ReplaceAll(stringFunc, "__FUNCTYPE__", c.funcType)
		stringFunc = strings.ReplaceAll(stringFunc, "__CALLFUNC__", c.callFunc)
		stringFunc = strings.ReplaceAll(stringFunc, "__ALLPARAM__", strings.Join(c.params, ", "))
		_, err = file.WriteString(stringFunc)
		if err != nil {
			log.Panicf("Cannot write content to %s", *output)
		}
	}
	for _, c := range int64FuncContent {
		int64Func := funcHeader
		for _, s := range c.params {
			int64Func += singleInt64
			int64Func = strings.ReplaceAll(int64Func, "__PARAM__", s)
		}
		int64Func += callFunc + jsonMarshal + dataReturn
		int64Func = strings.ReplaceAll(int64Func, "__FUNC__", c.funcName)
		int64Func = strings.ReplaceAll(int64Func, "__FUNCTYPE__", c.funcType)
		int64Func = strings.ReplaceAll(int64Func, "__CALLFUNC__", c.callFunc)
		int64Func = strings.ReplaceAll(int64Func, "__ALLPARAM__", strings.Join(c.params, ", "))
		_, err = file.WriteString(int64Func)
		if err != nil {
			log.Panicf("Cannot write content to %s", *output)
		}
	}
	for _, c := range int64NoDataFuncContent {
		int64NoDataFunc := funcHeader
		for _, s := range c.params {
			int64NoDataFunc += singleInt64
			int64NoDataFunc = strings.ReplaceAll(int64NoDataFunc, "__PARAM__", s)
		}
		int64NoDataFunc += callNoDataFunc + noDataReturn
		int64NoDataFunc = strings.ReplaceAll(int64NoDataFunc, "__FUNC__", c.funcName)
		int64NoDataFunc = strings.ReplaceAll(int64NoDataFunc, "__FUNCTYPE__", c.funcType)
		int64NoDataFunc = strings.ReplaceAll(int64NoDataFunc, "__CALLFUNC__", c.callFunc)
		int64NoDataFunc = strings.ReplaceAll(int64NoDataFunc, "__ALLPARAM__", strings.Join(c.params, ", "))
		_, err = file.WriteString(int64NoDataFunc)
		if err != nil {
			log.Panicf("Cannot write content to %s", *output)
		}
	}
	for _, c := range intFuncContent {
		intFunc := funcHeader
		for _, s := range c.params {
			intFunc += singleInt
			intFunc = strings.ReplaceAll(intFunc, "__PARAM__", s)
		}
		intFunc += callFunc + jsonMarshal + dataReturn
		intFunc = strings.ReplaceAll(intFunc, "__FUNC__", c.funcName)
		intFunc = strings.ReplaceAll(intFunc, "__FUNCTYPE__", c.funcType)
		intFunc = strings.ReplaceAll(intFunc, "__CALLFUNC__", c.callFunc)
		intFunc = strings.ReplaceAll(intFunc, "__ALLPARAM__", strings.Join(c.params, ", "))
		_, err = file.WriteString(intFunc)
		if err != nil {
			log.Panicf("Cannot write content to %s", *output)
		}
	}
}
