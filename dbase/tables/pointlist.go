package tables

// 积分记录
type PointList struct {
	Rd        uint   `json:"rd" gorm:"primaryKey;comment:'主键'"`         // 主键
	Roomid    string `json:"roomid" gorm:"uniqueIndex;comment:'群聊 id'"` // 群聊 id
	Wxid      string `json:"wxid" gorm:"comment:'微信用户id'"`              // 微信用户id
	Point     int32  `json:"point" gorm:"comment:'积分'"`                 // 积分
	Type      int32  `json:"type" gorm:"comment:'积分类型 [1=看图猜成语]'"`      // 积分类型 [1=看图猜成语]
	Decs      int32  `json:"decs"`                                      // 积分增加减少：[1=增加，2=减少]
	CreatedAt int64  `json:"created_at" gorm:"comment:'创建时间戳'"`         // 创建时间戳
	UpdatedAt int64  `json:"updated_at" gorm:"comment:'最后更新时间戳'"`       // 最后更新时间戳
}
