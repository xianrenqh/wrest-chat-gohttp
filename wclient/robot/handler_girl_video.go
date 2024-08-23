package robot

import (
	"fmt"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/point"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient/deliver"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func girlVideoHandler() []*Handler {
	cmds := []*Handler{}
	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    400,
		Roomid:   "*",
		Command:  "美女视频",
		Describe: "输入指定指令，如：'视频','美女视频'等，即可获取小姐姐视频",
		Callback: func(msg *wcferry.WxMsg) string {
			if !msg.IsGroup {
				return ""
			}
			wxId := msg.Sender
			roomId := msg.Roomid
			videoPoint := args.Others.VideoPoint

			//判断权限级别
			up, _ := profile.Fetch(&profile.FetchParam{Wxid: wxId, Roomid: roomId})
			p, _ := point.Fetch(&point.FetchParam{Wxid: wxId, Roomid: roomId})
			if up.Level < 7 {
				// 判断群聊积分是否满足
				if p.Point < videoPoint {
					return "积分不足，请先努力赚积分吧~\n\n偷偷告诉你：\n平时多舔舔管理员，让他给你施舍点~~"
				}
			} else {
				//不做限制
				msg := "您是尊贵的管理员/超级管理员，本次查询不扣除积分"
				if u := wc.CmdClient.GetInfoByWxid(wxId); u != nil && u.Name != "" {
					msg = "@" + u.Name + "\n" + msg
				}
				wc.CmdClient.SendTxt(msg, roomId, wxId)
			}
			// 获取视频
			videosApi := args.ApiServer.VideosApi
			dir, err := os.Getwd()
			if err != nil {
				log.Error().Err(err).Msg("获取项目目录失败:")
				return ""
			}

			if len(videosApi) > 0 {
				randomIndex := rand.Intn(len(videosApi))
				selectedVideo := videosApi[randomIndex]
				videoSrc, err := deliver.GetFiles(selectedVideo, 2)

				if err != nil {
					log.Error().Err(err).Msg("请求失败:")
					return fmt.Sprintf("%s", err)
				}
				if videoSrc == "" {
					return "图片获取失败，请稍后再试"
				}

				if up.Level < 7 {
					// 扣减积分
					deliver.UpdateOrCreatePoints(wxId, roomId, 3, int(videoPoint), 2, "查询美女视频")
					// 扣减积分
					msg := "请求资源成功，扣除积分：" + strconv.FormatInt(int64(videoPoint), 10) + "\n"
					msg += "当前剩余积分：" + strconv.FormatInt(int64(p.Point-videoPoint), 10)
					if u := wc.CmdClient.GetInfoByWxid(wxId); u != nil && u.Name != "" {
						msg = "@" + u.Name + "\n" + msg
					}
					wc.CmdClient.SendTxt(msg, roomId, wxId)
				}

				videoPath := dir + "/" + videoSrc
				//fmt.Println(videoPath)
				wc.CmdClient.SendFile(dir+"/"+videoSrc, roomId)

				//5秒后删除资源
				time.Sleep(5 * time.Second)
				err = os.Remove(videoPath)
				if err != nil {
					return ""
				}

			} else {
				return "请先配置视频Api接口"
			}

			return ""
		},
	})
	return cmds
}
