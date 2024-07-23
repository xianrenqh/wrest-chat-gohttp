package robot

import (
	"encoding/xml"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
	"log"
)

func receiver47(msg *wcferry.WxMsg) {
	// 石头剪刀布 | 表情图片
	var ret types.MsgContent47
	err := xml.Unmarshal([]byte(msg.Content), &ret)
	if err != nil {
		log.Printf("Error unmarshalling XML: %v", err)
		return
	}

	log.Printf("Emoji Type: %d", ret.Emoji.Type)

	if ret.Emoji.Type == 47 {
		// 处理类型为47的表情
		log.Println("Handling emoji type 47")
		// 在这里添加处理逻辑
	}

}
