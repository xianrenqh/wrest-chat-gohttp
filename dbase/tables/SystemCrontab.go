package tables

type SystemCrontab struct {
	Id              uint   `json:"id" gorm:"primaryKey"` // 主键
	Title           string `json:"title"`
	Type            int    `json:"type"`                           // 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Rule            string `json:"rule"`                           // 任务执行表达式
	RuleParams      string `json:"rule_params"`                    // 任务执行表达式的字段
	Target          string `json:"target"`                         // 调用任务字符串
	RunningTimes    int64  `json:"running_times" gorm:"default:0"` // 已运行次数
	LastRunningTime string `json:"last_running_time"`              //上次运行时间
	Remark          string `json:"remark"`
	Sort            int64  `json:"sort"`
	Status          int    `json:"status"`
	Singleton       int    `json:"singleton"` // 是否单次执行 (0 是 1 不是)
	Deliver         string `json:"deliver"`   // 投递到：（多个用引文逗号隔开）
	Parameter       string `json:"parameter"`
	DeliverType     int8   `json:"deliver_type"` // 投递类型：【1=私聊；2=群聊】
	CreateTime      int64  `json:"create_time"`  // 创建时间戳
	UpdateTime      int64  `json:"update_time"`  // 最后更新时间戳
}
