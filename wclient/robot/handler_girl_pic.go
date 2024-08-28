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

func girlPicHandler() []*Handler {
	cmds := []*Handler{}
	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    400,
		Roomid:   "*",
		Command:  "美女图片",
		Emoij:    "🖼",
		Describe: "输入指定指令，如：'图片','妹子','美女'等，即可获取小姐姐图片",
		Callback: func(msg *wcferry.WxMsg) string {
			selfWxid := wc.CmdClient.GetSelfWxid()
			wxId := msg.Sender
			roomId := msg.Roomid

			if !msg.IsGroup || wxId == selfWxid {
				//非群聊或者是自己的消息，过滤
				return ""
			}

			picPoint := args.Others.PicPoint

			//判断权限级别
			up, _ := profile.Fetch(&profile.FetchParam{Wxid: wxId, Roomid: roomId})
			p, _ := point.Fetch(&point.FetchParam{Wxid: wxId, Roomid: roomId})
			if up.Level < 7 {
				// 判断群聊积分是否满足
				if p.Point < picPoint {
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
			// 获取图片
			picApi := args.ApiServer.PicApi
			dir, err := os.Getwd()
			if err != nil {
				log.Error().Err(err).Msg("获取项目目录失败:")
				return ""
			}
			if len(picApi) > 0 {
				randomIndex := rand.Intn(len(picApi))
				selectedPic := picApi[randomIndex]
				imgSrc, err := deliver.GetFiles(selectedPic, 1)
				if err != nil {
					log.Error().Err(err).Msg("请求失败:")
					return fmt.Sprintf("%s", err)
				}
				if imgSrc == "" {
					return "图片获取失败，请稍后再试"
				}

				if up.Level < 7 {
					// 扣减积分
					deliver.UpdateOrCreatePoints(wxId, roomId, 3, picPoint, 2, "查询美女图片")
					// 扣减积分
					msg := "请求资源成功，扣除积分：" + strconv.FormatInt(int64(picPoint), 10) + "\n"
					msg += "当前剩余积分：" + strconv.FormatInt(int64(p.Point-picPoint), 10)
					if u := wc.CmdClient.GetInfoByWxid(wxId); u != nil && u.Name != "" {
						msg = "@" + u.Name + "\n" + msg
					}
					wc.CmdClient.SendTxt(msg, roomId, wxId)
				}

				imagePath := dir + "/" + imgSrc
				wc.CmdClient.SendImg(dir+"/"+imgSrc, roomId)

				//5秒后删除资源
				time.Sleep(5 * time.Second)
				err = os.Remove(imagePath)
				if err != nil {
					return ""
				}
			} else {
				return "请先配置图片Api接口"
			}
			return ""
		},
	})
	return cmds
}
