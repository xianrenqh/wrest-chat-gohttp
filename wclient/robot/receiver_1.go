package robot

import (
	"encoding/xml"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"github.com/opentdp/wrest-chat/wclient/aichat"
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
	videoKeyWords := args.FunctionKeyWord.VideoWord
	FishKeyWords := args.FunctionKeyWord.FishWord
	KfcKeyWords := args.FunctionKeyWord.KfcWord
	DogKeyWords := args.FunctionKeyWord.DogWord

	if judgeEqualListWord(msg.Content, videoKeyWords) {
		// 美女视频

	} else if judgeEqualListWord(msg.Content, FishKeyWords) {
		// 摸鱼日记

	} else if judgeEqualListWord(msg.Content, KfcKeyWords) {
		// kfc疯狂星期四

	} else if judgeEqualListWord(msg.Content, DogKeyWords) {
		// 舔狗日记

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

func judgeEqualListWord(content string, picKeyWords []string) bool {
	words := strings.Split(content, " ") // 假设 content 中的单词由空格分隔
	for _, word := range words {
		found := false
		for _, keyWord := range picKeyWords {
			if word == keyWord {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
