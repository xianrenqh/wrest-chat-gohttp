package robot

import (
	"github.com/opentdp/wrest-chat/dbase/forward"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func forwardMsgData(msg *wcferry.WxMsg) {
	time.Sleep(1 * time.Second) // 等待一秒

	// 判断是否转发群
	rq := &forward.FetchParam{
		Roomid: msg.Roomid,
		Status: 1,
	}
	res, err := forward.Fetch(rq)
	if err != nil {
		// 记录或处理错误
		log.Error().Err(err).Msg("获取转发信息失败")
		return
	}

	// 检查是否需要转发此消息
	if res.Rd == 0 || (res.Wxid != "" && res.Wxid != msg.Sender) || res.SendRoomids == "" {
		return
	}

	// 检查接收群是否一致并获取接收群列表
	if msg.Roomid != res.Roomid {
		return // 如果源群和目标群不一致，则不执行转发
	}
	sendRoomids := strings.Split(res.SendRoomids, ",")

	time.Sleep(2 * time.Second) // 再次延迟一秒
	for _, receive := range sendRoomids {
		if wc.CmdClient.ForwardMsg(msg.Id, receive) == 1 {
			time.Sleep(time.Duration(0.5 * float64(time.Second))) // 成功转发后延迟半秒
			log.Info().Msg("消息转发成功")
		} else {
			// 处理转发失败的情况，例如记录日志等。
			log.Warn().Msg("消息转发失败")
		}
	}
}
