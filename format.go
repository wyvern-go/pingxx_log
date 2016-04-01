package pingxx_log

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

type LogInfo struct {
	LogId       string `json:"log_id"`
	LogLevel    string    `json:"log_level"`
	Module      string  `json:"module"`
	LogType     string  `json:"type,omitempty"`
	LogTime     string `json:"time"`
	Filename    string  `json:"filename"`
	Line        int   `json:"line"`
	Remark      string   `json:"remark,omitempty"`

	AcctId      string  `json:"acct_id,omitempty"`
	AppId       string  `json:"app_id,omitempty"`
	Channel     string `json:"channel,omitempty"`
	ObjectId    string `json:"object_id,omitempty"`
	Mode        int `json:"mode,omitempty"`
	Agent       string `json:"agent,omitempty"`
	Refer       string  `json:"refer,omitempty"`
	Url         string `json:"url,omitempty"`
	ReqMethod   string  `json:"req_method,omitempty"`
	ReqHeader   string `json:"req_header,omitempty"`
	ReqParam    string `json:"req_param,omitempty"`
	RepHeader   string `json:"rep_header,omitempty"`
	RepHttpcode int `json:"rep_httpcode,omitempty"`
	RepResult   string `json:"rep_result,omitempty"`
	Runtime     int `json:"runtime,omitempty"`
	Ip          string `json:"ip,omitempty"`
}

func NewLogInfo() *LogInfo {
	return new(LogInfo)
}

func (info LogInfo) ToJson() ([]byte, error) {
	return json.Marshal(&info)
}

func (info LogInfo) ToStd() string {
	return fmt.Sprintf("%s [%s:%d] <%s> %s: %s", info.LogTime, info.Filename, info.Line, info.Module, info.LogLevel, info.Remark)
}

func (info *LogInfo) SetAcctId(acctid string) *LogInfo {
	info.AcctId = acctid
	return info
}

func (info *LogInfo) SetAppId(appid string) *LogInfo {
	info.AppId = appid
	return info
}

func (info *LogInfo) SetChannel(channel string) *LogInfo {
	info.Channel = channel
	return info
}

func (info *LogInfo) SetMode(mode bool) *LogInfo {
	if mode {
		info.Mode = 1
	}else {
		info.Mode = 0
	}
	return info
}

func (info *LogInfo) SetAgent(agent string) *LogInfo {
	info.Agent = agent
	return info
}

func (info *LogInfo) SetRefer(refer string) *LogInfo {
	info.Refer = refer
	return info
}

func (info *LogInfo) SetUrl(requrl string) *LogInfo {
	info.Url = requrl
	return info
}

func (info *LogInfo) SetRequestInfo(req *http.Request) *LogInfo {
	info.ReqMethod = req.Method
	info.ReqHeader = ""
	for k, v := range req.Header {
		info.ReqHeader = info.ReqHeader + fmt.Sprintf("%s:%v;", k, v[0])
	}
	info.Url = strings.Split(info.Url, "?")[0]
	info.ReqParam = strings.Split(info.Url, "?")[1]
	return info
}

func (info *LogInfo)SetResponseInfo(response http.Response) *LogInfo {
	var by []byte
	by, _ = ioutil.ReadAll(response.Body)
	info.RepResult = string(by)
	info.RepHttpcode = response.StatusCode
	for k, v := range response.Header {
		info.RepHeader = info.RepHeader + fmt.Sprintf("%s:%v;", k, v[0])
	}
	return info
}

func (info *LogInfo) SetReqMethod(mothod string) *LogInfo {
	info.ReqMethod = mothod
	return info
}

func (info *LogInfo) SetReqHeader(reqheader string) *LogInfo {
	info.ReqHeader = reqheader
	return info
}

func (info *LogInfo) SetReqParam(reqparam string) *LogInfo {
	info.ReqParam = reqparam
	return info
}

func (info *LogInfo) SetRepHeader(repheader string) *LogInfo {
	info.RepHeader = repheader
	return info
}

func (info *LogInfo) SetRepHttpcode(httpcode int) *LogInfo {
	info.RepHttpcode = httpcode
	return info
}

func (info *LogInfo) SetRuntime(run_time int) *LogInfo {
	info.Runtime = run_time
	return info
}

func (info *LogInfo) SetRepResult(represult string) *LogInfo {
	info.RepResult = represult
	return info
}

func (info *LogInfo) SetIp(ip string) *LogInfo {
	info.Ip = ip
	return info
}