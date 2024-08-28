package robot

import (
	"encoding/json"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/wcferry"
	"io/ioutil"
	"net/http"
)

func fishHandler() []*Handler {
	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    500,
		Roomid:   "*",
		Command:  "摸鱼日记",
		Emoij:    "📖",
		Describe: "发送摸鱼图片",
		Callback: fishCallback,
	})

	return cmds
}

type ApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func fishCallback(msg *wcferry.WxMsg) string {
	picApi := args.ApiServer.FishApi
	resp, err := http.Get(picApi)
	if err != nil {
		return "请求失败"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "读取响应数据失败"
	}
	var apiResp ApiResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return "解析JSON失败"
	}

	if apiResp.Code == 200 && apiResp.Msg == "success" {
		wc.CmdClient.SendImg(apiResp.Data, msg.Roomid)
		return ""
	} else {
		return "API返回错误"
	}
}
