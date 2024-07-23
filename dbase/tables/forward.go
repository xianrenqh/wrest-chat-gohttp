package tables

type Forword struct {
	Rd          uint   `json:"rd" gorm:"primaryKey"`      // 主键
	Roomid      string `json:"roomid" gorm:"uniqueIndex"` // 群聊 id
	Wxid        string `json:"wxid"`                      // 监听人（如果为空，监听全部发言）
	SendRoomids string `json:"send_roomids"`              // 转发到群聊id
	Status      int32  `json:"status"`                    // 状态：0:关闭，1:开启
	CreatedAt   int64  `json:"created_at"`                // 创建时间戳
	UpdatedAt   int64  `json:"updated_at"`                // 最后更新时间戳
}
