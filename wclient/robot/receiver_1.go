package robot

import (
	"encoding/xml"
	"fmt"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"strings"

	"github.com/opentdp/wrest-chat/wcferry"
)

// 处理新消息
func receiver1(msg *wcferry.WxMsg) {
	selfWxid := wc.CmdClient.GetSelfWxid()
	// 判断是否是群聊
	if msg.IsGroup {
		ret := &types.MsgXmlAtUser{}
		err := xml.Unmarshal([]byte(msg.Xml), ret)
		if err == nil && ret.AtUserList != "" {
			users := strings.Split(ret.AtUserList, ",")
			fmt.Println("获取到的@用户列表：", users)
			for _, v := range users {
				if v == "" {
					continue
				}
				if v == selfWxid {
					//wc.CmdClient.SendTxt("你@我干嘛？", msg.Roomid, msg.Sender)
					continue
				}
			}
		}
	}

	// 处理聊天指令
	if msg.IsGroup || wcferry.ContactType(msg.Sender) == "好友" {
		output := ApplyHandlers(msg)
		if strings.Trim(output, "-") != "" {
			reply(msg, output)
		}
		return
	}

}
