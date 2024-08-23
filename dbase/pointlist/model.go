package pointlist

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

type CreateParam struct {
	Rd     uint   `json:"rd" gorm:"primaryKey;comment:'主键'"`
	Roomid string `json:"roomid" gorm:"index;comment:'群聊 id'"` // 群聊 id
	Wxid   string `json:"wxid" gorm:"comment:'微信用户id'"`        // 微信用户id
	Point  int    `json:"point" gorm:"comment:'积分'"`           // 积分
	Type   int32  `json:"type" gorm:""`                        // 积分类型 [1=看图猜成语；2=签到]
	Sign   int32  `json:"sign" gorm:"index"`                   // 积分增加减少：[1=增加，2=减少]
	Desc   string `json:"desc"`
}

func Create(data *CreateParam) (uint, error) {
	item := &tables.PointList{
		Roomid: data.Roomid,
		Wxid:   data.Wxid,
		Point:  data.Point,
		Type:   data.Type,
		Sign:   data.Sign,
		Desc:   data.Desc,
	}

	result := dborm.Db.Create(item)
	return item.Rd, result.Error
}

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {
	result := dborm.Db.
		Where(&tables.PointList{
			Rd: data.Rd,
		}).
		Updates(tables.PointList{
			Roomid: data.Roomid,
			Wxid:   data.Wxid,
			Point:  data.Point,
			Type:   data.Type,
			Sign:   data.Sign,
			Desc:   data.Desc,
		})

	return result.Error
}

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Roomid string `json:"roomid"`
	Wxid   string `json:"wxid"`
	Type   int32  `json:"type"`
	Sign   int32  `json:"sign"`
}

func Fetch(data *FetchParam) (*tables.PointList, error) {
	var item *tables.PointList

	result := dborm.Db.
		Where(&tables.PointList{
			Roomid: data.Roomid,
			Wxid:   data.Wxid,
			Type:   data.Type,
			Sign:   data.Sign,
		}).Order("rd desc").Find(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	if item == nil {
		item = &tables.PointList{Roomid: data.Roomid}
	}

	return item, nil
}
