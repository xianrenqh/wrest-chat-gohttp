package robot

import (
	"encoding/xml"
	"fmt"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"github.com/opentdp/wrest-chat/wclient/deliver"
	"strconv"
	"strings"
	"time"
)

func pointAddHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    500,
		Roomid:   "*",
		Command:  "加",
		Emoij:    "💯",
		Describe: "添加积分给某人，例如：加 100 @机器人",
		Callback: pointAddCallback,
	})
	return cmds
}

func pointAddCallback(msg *wcferry.WxMsg) string {
	wxId := msg.Sender
	roomId := msg.Roomid
	ret := &types.MsgXmlAtUser{}
	err := xml.Unmarshal([]byte(msg.Xml), ret)
	if err == nil && ret.AtUserList != "" {

		users := strings.Split(ret.AtUserList, ",")
		params := strings.SplitN(msg.Content, " ", 2)
		getPoint, _ := strconv.Atoi(params[0])
		if getPoint < 1 {
			return "积分不能小于1"
		}
		// 判断权限级别
		up, _ := profile.Fetch(&profile.FetchParam{Wxid: wxId, Roomid: roomId})
		if up.Level < 9 {
			// 非超级管理员最多增加的积分值
			AddPointManager := args.Others.AddPointManager //管理员最多增加的积分值
			if getPoint > AddPointManager {
				return fmt.Sprintf("管理员最多增加的积分值是： %d", AddPointManager)
			}
		}
		for _, v := range users {
			v := strings.TrimSpace(v)
			if v != "" {
				pointNew := deliver.UpdateOrCreatePoints(v, roomId, 2, getPoint, 1, "管理赠送")
				resMsg := fmt.Sprintf("基于您的表现，尊贵的管理员给您施舍了 %d 分，请您好好珍惜。\n当前可用积分: %d", getPoint, pointNew)

				if u := wc.CmdClient.GetInfoByWxid(v); u != nil && u.Name != "" {
					resMsg = "@" + u.Name + "\n" + resMsg
				}

				time.Sleep(1200 * time.Millisecond)
				wc.CmdClient.SendTxt(resMsg, roomId, v)
			}
		}
		time.Sleep(2500 * time.Millisecond)
		wc.CmdClient.SendTxt("其他人还望好好努力，平时舔舔管理员 让管理给你施舍点", roomId, "")

	} else {
		return "请指定要添加积分的用户"
	}

	return ""
}
