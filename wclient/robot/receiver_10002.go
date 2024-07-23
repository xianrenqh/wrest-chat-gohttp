package robot

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/message"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
)

// 处理撤回消息
func receiver10002(msg *wcferry.WxMsg) {

	var output string

	// 获取撤回提示
	if msg.IsGroup {

		// 自动回应群聊拍一拍（私聊不支持）
		if strings.Contains(msg.Content, "拍了拍我") {
			room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
			if room.PatReturn == "true" {
				wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
			}
			return
		}
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		output = room.RevokeMsg
	} else {
		output = setting.RevokeMsg
	}

	// 防撤回提示已关闭
	if len(output) < 2 {
		return
	}

	// 解析已撤回的消息
	revoke := types.MsgContent10002{}
	err := xml.Unmarshal([]byte(msg.Content), &revoke)

	if err != nil || revoke.RevokeMsg.NewMsgID == "" {
		return
	}

	// 获取已撤回消息的 Id
	id, err := strconv.Atoi(revoke.RevokeMsg.NewMsgID)
	if err != nil || id == 0 {
		return
	}

	// 取回已撤回的消息内容
	origin, err := message.Fetch(&message.FetchParam{Id: uint64(id)})
	if err != nil || origin.Content == "" {
		return
	}

	// 提示已撤回的消息内容
	str := strings.TrimSpace(origin.Content)
	strImg := strings.TrimSpace(origin.Remark)

	xmlPrefixes := []string{"<?xml", "<sysmsg", "<msg"}
	for _, prefix := range xmlPrefixes {
		if strings.HasPrefix(str, prefix) {
			str = ""
		}
	}

	if str != "" {
		output += "一则流芳百世的金句？\n\n"
		output += "小样，我还治不了你了：\n"
		output += str + "\n"
		reply(msg, output)
		return
	}

	if origin.Type == 3 {
		output += "一张不可描述的图片？\n\n"
		output += "小样，我还治不了你了：\n"
		reply(msg, output)
		replyImg(msg, strImg)
		return
	}

	if origin.Type == 47 {
		xmlStr := origin.Content
		msgContent, err := parseXML(xmlStr)

		if err != nil {
			return
		}
		CDNURL := msgContent.Emoji.CDNURL
		output += "一个震惊四座的表情？\n\n"
		output += "小样，我还治不了你了：\n"
		reply(msg, output)
		replyImg(msg, CDNURL)
		return
	}

	if origin.Type == 49 {
		appmsg := types.MsgContent49{}
		err := xml.Unmarshal([]byte(origin.Content), &appmsg)
		if err == nil {
			switch appmsg.AppMsg.Type {
			case 6:
				output += "一份暗藏机密的文件？"
			case 19:
				output += "多条来自异界的消息？"
			case 57:
				output += "一则流芳百世的金句？"
			default:
				output += "系统也无法揣测的消息？"
			}
			reply(msg, output)
			return
		}
	}

	output += "系统也无法揣测的消息？"
	reply(msg, output)

}

// 解析 XML 字符串并将其存储在 MsgContent47 结构体中
func parseXML(xmlStr string) (*types.MsgContent47, error) {
	var msgContent types.MsgContent47
	err := xml.Unmarshal([]byte(xmlStr), &msgContent)
	if err != nil {
		return nil, err
	}
	return &msgContent, nil
}
