package robot

import (
	"encoding/xml"
	"fmt"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"github.com/opentdp/wrest-chat/wclient/aichat"
	"github.com/opentdp/wrest-chat/wclient/deliver"
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
			//fmt.Println("获取到的@用户列表：", users)
			for _, v := range users {
				if v == "" {
					continue
				}
				if v == selfWxid {
					msgContent := msg.Content
					data := aichat.Text(msgContent, selfWxid, msg.Sender)
					if u := wc.CmdClient.GetInfoByWxid(msg.Sender); u != nil && u.Name != "" {
						data = "@" + u.Name + "\n" + data
					}
					wc.CmdClient.SendTxt(data, msg.Roomid, msg.Sender)
				}
			}
		}
	}

	// 娱乐模式
	FishKeyWords := args.FunctionKeyWord.FishWord
	DogKeyWords := args.FunctionKeyWord.DogWord

	if deliver.JudgeEqualListWord(msg.Content, FishKeyWords) {
		// 摸鱼日记

	} else if deliver.JudgeEqualListWord(msg.Content, DogKeyWords) {
		// 舔狗日记

	}

	//处理转账
	if msg.IsGroup && strings.Contains(msg.Content, "转账待你接收") {
		ret := &types.MsgXmlSilence{}
		err := xml.Unmarshal([]byte(msg.Xml), ret)
		fmt.Println("代收款")
		fmt.Println(err)
		fmt.Println(msg.Sender)
		fmt.Println(msg.Roomid)
		silence := ret.Msgsource.Silence
		fmt.Println(silence)
	}

	// 处理收款
	if msg.IsGroup && strings.Contains(msg.Content, "已收款") {
		fmt.Println("已收款")
		fmt.Println(msg.Sender)
		fmt.Println(msg.Roomid)
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
