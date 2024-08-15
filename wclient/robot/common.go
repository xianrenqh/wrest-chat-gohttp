package robot

import (
	"strings"

	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient"
	"github.com/rs/zerolog/log"
)

var wc *wcferry.Client

func Start() {

	log.Info().Msg("正在初始化Robot服务中。。")

	if !setting.BotEnable {
		log.Warn().Msg("Robot已关闭")
		return
	}

	if wc != nil {
		log.Info().Msg("Robot重启成功")
		return
	}

	wc = wclient.Register()
	_, err := wc.EnrollReceiver(true, receiver)
	if err != nil {
		log.Error().Err(err).Msg("Robot启动失败")
	}

	ResetHandlers()

}

func Reset() {

	setting.Laod()
	ResetHandlers()

}

///////////////////////// COMMON METHODS /////////////////////////

// 会话场景
func prid(msg *wcferry.WxMsg) string {

	if msg.IsGroup {
		return msg.Roomid
	}
	return "-"

}

// 回复消息
func reply(msg *wcferry.WxMsg, text string) int32 {

	if msg.IsSelf {
		return -2
	}

	if text = strings.TrimSpace(text); text == "" {
		return -1
	}

	if msg.IsGroup {
		if msg.Sender != "" && wcferry.ContactType(msg.Sender) == "好友" {
			user := wc.CmdClient.GetInfoByWxid(msg.Sender)
			if user != nil && user.Name != "" {
				text = "@" + user.Name + "\n" + text
			}
		}
		return wc.CmdClient.SendTxt(text, msg.Roomid, msg.Sender)
	}

	return wc.CmdClient.SendTxt(text, msg.Sender, "")

}

// 回复图片消息
func replyImg(msg *wcferry.WxMsg, path string) int32 {

	if msg.IsSelf {
		return -2
	}

	if path = strings.TrimSpace(path); path == "" {
		return -1
	}

	if msg.IsGroup {
		if msg.Sender != "" && wcferry.ContactType(msg.Sender) == "好友" {
			user := wc.CmdClient.GetInfoByWxid(msg.Sender)
			if user != nil && user.Name != "" {

			}
		}
		return wc.CmdClient.SendImg(path, msg.Roomid)
	}
	return wc.CmdClient.SendImg(path, msg.Sender)
}
