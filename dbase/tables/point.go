package tables

type Point struct {
	Rd        uint   `json:"rd" gorm:"primaryKey;comment:'主键'"`        // 主键
	Roomid    string `json:"roomid" gorm:"index;comment:'群聊 id'"`      // 群聊 id
	Wxid      string `json:"wxid" gorm:"index;comment:'微信用户id'"`     // 微信用户id
	Point     int32  `json:"point" gorm:"index;comment:'积分'"`          // 总积分
	CreatedAt int64  `json:"created_at" gorm:"comment:'创建时间戳'"`     // 创建时间戳
	UpdatedAt int64  `json:"updated_at" gorm:"comment:'最后更新时间戳'"` // 最后更新时间戳
}
