package args

// 日志配置

type ILog struct {
	Dir    string `yaml:"Dir"`    // 存储目录
	Level  string `yaml:"Level"`  // 记录级别 debug|info|warn|error
	Target string `yaml:"Target"` // 输出方式 both|file|null|stdout|stderr
}

var Log = &ILog{
	Dir:    "logs",
	Level:  "info",
	Target: "stdout",
}

// Wcf 服务

type IWcf struct {
	Address      string `yaml:"Address"`      // Rpc 监听地址
	MsgPrint     bool   `yaml:"MsgPrint"`     // 是否打印收到的消息
	MsgStore     bool   `yaml:"MsgStore"`     // 是否存储收到的消息
	MsgStoreDays int    `yaml:"MsgStoreDays"` // 消息留存天数
	WcfBinary    string `yaml:"WcfBinary"`    // 留空则不注入微信
	WeChatAuto   bool   `yaml:"wechatAuto"`   // 是否自动启动微信
}

var Wcf = &IWcf{
	Address:    "127.0.0.1:7601",
	WeChatAuto: true,
}

// Web 服务

type IWeb struct {
	Address string `yaml:"Address"` // Web 监听地址
	PushUrl string `yaml:"PushUrl"` // 消息推送地址
	Storage string `yaml:"Storage"` // 附件存储路径
	Swagger bool   `yaml:"Swagger"` // 是否启用 Api 文档
	Token   string `yaml:"Token"`   // 使用 Token 验证请求
}

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Storage: "storage",
	Swagger: true,
}

type IOthers struct {
	SignInPoint int   `yaml:"SignInPoint"`
	PicPoint    int   `yaml:"PicPoint"`
	VideoPoint  int32 `yaml:"VideoPoint"`
}

var Others = &IOthers{
	SignInPoint: 10,
	PicPoint:    1,
	VideoPoint:  2,
}

type IFunctionKeyWord struct {
	VideoWord []string `yaml:"videoWord"`
	FishWord  []string `yaml:"fishWord"`
	KfcWord   []string `yaml:"kfcWord"`
	DogWord   []string `yaml:"dogWord"`
}

var FunctionKeyWord = &IFunctionKeyWord{}

type IApiServer struct {
	FishApi string   `yaml:"fishApi"`
	PicApi  []string `yaml:"picApi"`
}

var ApiServer = &IApiServer{}
