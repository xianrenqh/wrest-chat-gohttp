package robot

import (
	"fmt"
	"github.com/opentdp/wrest-chat/dbase/point"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/wcferry"
)

func pointHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    500,
		Roomid:   "*",
		Command:  "积分查询",
		Emoij:    "💯",
		Describe: "查询当前群聊积分",
		Callback: pointCallback,
	})
	return cmds
}
func pointCallback(msg *wcferry.WxMsg) string {

	myPoint := 0
	upPoint, _ := point.Fetch(&point.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
	if upPoint.Rd > 0 {
		myPoint = upPoint.Point
	}

	//判断权限级别
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
	if up.Level < 7 {
		return fmt.Sprintf("您当前剩余积分: %d \n还望好好努力，平时舔舔管理员，让管理给你施舍点！", myPoint)
	} else {
		return fmt.Sprintf("您当前剩余积分: %d \n尊敬的管理员，您可以赠送积分呦~", myPoint)
	}

}
