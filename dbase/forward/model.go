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
	Type        int32  `json:"type"`
	SendRoomids string `json:"send_roomids"`
	Status      int32  `json:"status"`
}

func Create(data *CreateParam) (uint, error) {
	item := &tables.Forward{
		Roomid:      data.Roomid,
		Wxid:        data.Wxid,
		Type:        data.Type,
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
		Where(&tables.Forward{
			Rd: data.Rd,
		}).
		Updates(tables.Forward{
			Roomid:      data.Roomid,
			Wxid:        data.Wxid,
			Type:        data.Type,
			SendRoomids: data.SendRoomids,
			Status:      data.Status,
		})

	return result.Error
}

// 获取消息转发

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Roomid string `json:"roomid"`
	Type   int32  `json:"type"`
}

func Fetch(data *FetchParam) (*tables.Forward, error) {

	var item *tables.Forward

	result := dborm.Db.
		Where(&tables.Forward{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.Forward{Roomid: data.Roomid}
	}

	return item, result.Error

}

// 删除消息转发

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Forward

	result := dborm.Db.
		Where(&tables.Forward{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取消息转发列表

type FetchAllParam struct {
	Type int32 `json:"type"`
}

func FetchAll(data *FetchAllParam) ([]*tables.Forward, error) {
	var items []*tables.Forward

	result := dborm.Db.Find(&items)

	return items, result.Error
}

// 获取转发配置总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Forward{}).
		Where(&tables.Forward{
			Type: data.Type,
		}).
		Count(&count)

	return count, result.Error

}
