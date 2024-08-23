package robot

import (
	"fmt"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/pointlist"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient/deliver"
	"time"
)

func signInHandler() []*Handler {
	cmds := []*Handler{}
	sendPoint := args.Others.SignInPoint //赠送的签到积分
	cmds = append(cmds, &Handler{
		Level:   0,
		Order:   1,
		Roomid:  "*",
		Command: "签到",
		//Describe: "签到获取积分" + sendPoint,
		Describe: fmt.Sprintf("签到获取 %d 积分", sendPoint),
		Callback: func(msg *wcferry.WxMsg) string {
			if !msg.IsGroup {
				return ""
			}

			wxid := msg.Sender
			roomId := msg.Roomid
			// 查询是否有记录，如果有记录，判断今天是否有数据，如果有数据，提示已经签到，如果没有数据，插入数据，提示签到成功
			up, _ := pointlist.Fetch(&pointlist.FetchParam{Wxid: wxid, Roomid: roomId, Type: 2})

			if up.Rd <= 0 {
				pointNew := deliver.UpdateOrCreatePoints(wxid, roomId, 2, sendPoint, 1, "签到")

				return fmt.Sprintf("恭喜你签到成功, 当前群聊可用积分: %d", pointNew)
			} else {
				now := time.Now().Unix()
				startOfDay := now - (now % 86400)                           // 获取今天开始时间的时间戳（凌晨零点）
				endOfDay := startOfDay + 86400 - 1                          // 获取今天的结束时间（明天凌晨零点之前的一秒）
				if up.CreatedAt >= startOfDay && up.CreatedAt <= endOfDay { // 注意使用 >= 和 <= 来判断时间范围
					return "你干嘛~ 你今天已经签到过了！"
				}

				pointNew := deliver.UpdateOrCreatePoints(wxid, roomId, 2, sendPoint, 1, "签到")
				return fmt.Sprintf("恭喜你签到成功, 当前群聊可用积分: %d", pointNew)
			}
		},
	})

	return cmds
}
