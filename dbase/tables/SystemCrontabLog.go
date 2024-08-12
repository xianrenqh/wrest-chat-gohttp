package tables

type SystemCrontabLog struct {
	Id          uint   `json:"id" gorm:"primaryKey"` // 主键
	CrontabId   int64  `json:"crontab_id"`
	Target      string `json:"target"`
	Parameter   string `json:"parameter"`
	Exception   string `json:"exception"`
	ReturnCode  int    `json:"return_code"`
	RunningTime string `json:"running_time"`
	CreateTime  int64  `json:"create_time"` // 创建时间戳
	UpdateTime  int64  `json:"update_time"` // 最后更新时间戳
}
