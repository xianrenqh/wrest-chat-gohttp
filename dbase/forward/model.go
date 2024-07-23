package forward

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建消息转发
type CreateParam struct {
	Rd          uint   `json:"rd" gorm:"primaryKey"`
	Roomid      string `json:"roomid" gorm:"uniqueIndex"`
	Wxid        string `json:"wxid"`
	SendRoomids string `json:"send_roomids"`
	Status      int32  `json:"status"`
}

func Create(data *CreateParam) (uint, error) {
	item := &tables.Forword{
		Roomid:      data.Roomid,
		Wxid:        data.Wxid,
		SendRoomids: data.SendRoomids,
		Status:      data.Status,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error
}

// 更新消息转发
type UpdateParam = CreateParam

func Update(data *UpdateParam) error {
	result := dborm.Db.
		Where(&tables.Forword{
			Rd: data.Rd,
		}).
		Updates(tables.Forword{
			Roomid:      data.Roomid,
			Wxid:        data.Wxid,
			SendRoomids: data.SendRoomids,
			Status:      data.Status,
		})

	return result.Error
}

// 获取消息转发

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Roomid string `json:"roomid"`
}

// 删除消息转发

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Forword

	result := dborm.Db.
		Where(&tables.Forword{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取消息转发列表

type FetchAllParam struct {
}

func FetchAll(data *FetchAllParam) ([]*tables.Forword, error) {
	var items []*tables.Forword

	result := dborm.Db.Find(&items)

	return items, result.Error
}
