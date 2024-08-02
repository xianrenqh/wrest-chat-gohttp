package deliver

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wrest-chat/wclient"
	"strings"
)

// 将执行结果投递到微信

func wechatMessage(args []string, message string) int32 {

	// 增加多个接收者参数，分割符”|“
	argsStr := strings.Split(strings.TrimSpace(args[0]), "|")

	wxid := ""
	if len(args) > 1 {
		wxid = args[1]
	}

	wc := wclient.Register()
	if wc == nil {
		logman.Error("deliver", "error", "wclient is nil")
		return -1
	}

	for _, value := range argsStr {
		wclient.SendFlexMsg(message, wxid, value)
	}
	return 1
}
