package robot

import (
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/wcferry"
)

func jokeHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    500,
		Roomid:   "*",
		Command:  "内涵段子",
		Emoij:    "🤣",
		Describe: "内涵的不只是段子",
		Callback: jokeCallback,
	})
	return cmds
}
func jokeCallback(msg *wcferry.WxMsg) string {
	jokeApi := args.ApiServer.JokeApi
	wc.CmdClient.SendImg(jokeApi, msg.Roomid)
	return ""
}
