package robot

import (
	"encoding/xml"
	"github.com/opentdp/wrest-chat/dbase/mparticle"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"github.com/rs/zerolog/log"
	"strconv"
)

// 处理混合类消息
func receiver49(msg *wcferry.WxMsg) {

	ret := types.MsgContent49{}
	err := xml.Unmarshal([]byte(msg.Content), &ret)
	if err != nil {
		return
	}

	//公众号消息
	if ret.AppMsg.Type == 5 && setting.MpArticleEnable {
		mpItems := ret.AppMsg.MmRreader.Category.Item
		pubUsername := ret.AppMsg.MmRreader.Publisher.Username
		pubNickname := ret.AppMsg.MmRreader.Publisher.Nickname

		for _, item := range mpItems {
			createParam := mparticle.CreateParam{
				Title:    item.Title,
				Desc:     item.Summary,
				Url:      item.URL,
				PubTime:  item.PubTime,
				Cover:    item.Cover,
				Digest:   item.Digest,
				Username: pubUsername,
				Appname:  pubNickname,
			}

			result, _ := mparticle.Create(&createParam)
			if result == 0 {
				log.Info().Str("Status Code：", strconv.Itoa(int(result))).Msg("公众号文章错误：" + item.Title)
				continue
			}

			log.Info().Msg("公众号文章添加成功：" + item.Title)
			continue
		}
	}

	// 引用消息
	if ret.AppMsg.Type == 57 {
		msg.Extra = msg.Content
		msg.Content = ret.AppMsg.Title
		msg.Id = ret.AppMsg.ReferMsg.Svrid
		msg.Type = ret.AppMsg.ReferMsg.Type
		msg.Sign = "refer-msg"
		receiver1(msg)
		return
	}

}
