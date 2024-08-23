package robot

import (
	"encoding/xml"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"strings"
)

func addManagerHandler() []*Handler {
	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    9,
		Order:    400,
		Roomid:   "*",
		Command:  "添加管理",
		Describe: "把**添加成管理员",
		Callback: func(msg *wcferry.WxMsg) string {
			if !msg.IsGroup {
				return ""
			}
			roomId := msg.Roomid
			selfWxid := wc.CmdClient.GetSelfWxid()
			ret := &types.MsgXmlAtUser{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				users := strings.Split(ret.AtUserList, ",")
				for _, v := range users {
					if v == "" {
						continue
					}
					if v == selfWxid {
						return "添加管理失败，不能添加自己"
					}
					//先查询是否已经添加，如果存在数据，且级别小于7，则更改级别
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: roomId})
					if up.Rd > 0 {
						if up.Level < 7 {
							//修改级别
							profile.Update(&profile.UpdateParam{Rd: up.Rd, Level: 7})
						} else {
							return "已经是管理员啦"
						}
					} else {
						// 添加数据
						profile.Create(&profile.CreateParam{Wxid: v, Roomid: roomId, Level: 7})
					}
				}
			}
			return "添加管理 成功"
		},
	})

	return cmds
}
