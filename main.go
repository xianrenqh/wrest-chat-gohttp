package main

import (
	"embed"
	"github.com/opentdp/wrest-chat/dbase/forward"
	"strconv"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase"
	"github.com/opentdp/wrest-chat/httpd"
	"github.com/opentdp/wrest-chat/wclient/crond"
	"github.com/opentdp/wrest-chat/wclient/plugin"
	"github.com/opentdp/wrest-chat/wclient/robot"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

//go:embed public
var efs embed.FS

func main() {
	//日志输出
	initLogger()

	args.Efs = &efs

	dbase.Connect()

	//查询转发配置数
	checkForward()

	crond.Daemon()
	plugin.CronjobPluginSetup()
	plugin.KeywordPluginSetup()

	robot.Start()

	httpd.Server()
}

func initLogger() {
	// 配置日志输出
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	})
}

func checkForward() {
	// 群转发消息
	count, err := forward.Count(&forward.CountParam{})
	if err != nil {
		return
	}
	log.Info().Msg("消息转发开启配置数：" + strconv.FormatInt(count, 10))
}
