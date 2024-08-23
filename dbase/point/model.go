package point

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

type CreateParam struct {
	Rd     uint   `json:"rd" gorm:"primaryKey;comment:'主键'"`   // 主键
	Roomid string `json:"roomid" gorm:"index;comment:'群聊 id'"` // 群聊 id
	Wxid   string `json:"wxid" gorm:"index;comment:'微信用户id'"`  // 微信用户id
	Point  int    `json:"point" gorm:"comment:'积分'"`           // 积分
}

func Create(data *CreateParam) (uint, error) {
	item := &tables.Point{
		Roomid: data.Roomid,
		Wxid:   data.Wxid,
		Point:  data.Point,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error
}

// 更新积分
type UpdateParam struct {
	Rd     uint   `json:"rd" gorm:"primaryKey;comment:'主键'"`   // 主键
	Roomid string `json:"roomid" gorm:"index;comment:'群聊 id'"` // 群聊 id
	Wxid   string `json:"wxid" gorm:"index;comment:'微信用户id'"`  // 微信用户id
	Point  int
}

func Update(data *UpdateParam) error {
	result := dborm.Db.
		Where(&tables.Point{
			Rd: data.Rd,
		}).
		Updates(&tables.Point{
			Roomid: data.Roomid,
			Wxid:   data.Wxid,
			Point:  data.Point,
		})
	return result.Error
}

// 获取积分

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Wxid   string `json:"wxid"`
	Roomid string `json:"roomid"`
}

func Fetch(data *FetchParam) (*tables.Point, error) {

	var item *tables.Point

	result := dborm.Db.
		Where(&tables.Point{
			Rd:     data.Rd,
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		Find(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	if item == nil {
		item = &tables.Point{Roomid: data.Roomid}
	}

	return item, nil

}

// 删除积分

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Point

	result := dborm.Db.
		Where(&tables.Point{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取积分列表

type FetchAllParam struct {
	Roomid string `json:"roomid"`
	Type   int32  `json:"type"`
	Limit  int
	Offset int
}

func FetchAll(data *FetchAllParam) ([]*tables.Point, error) {
	var items []*tables.Point

	if data.Limit <= 0 {
		data.Limit = 10 // 默认值为 10
	}

	result := dborm.Db.
		Where(&tables.Point{
			Roomid: data.Roomid,
		}).
		Limit(data.Limit).
		Offset(data.Offset).
		Find(&items)

	return items, result.Error
}

// 获取转发配置总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Point{}).
		Where(&tables.Point{
			Roomid: data.Roomid,
		}).
		Count(&count)

	return count, result.Error

}
