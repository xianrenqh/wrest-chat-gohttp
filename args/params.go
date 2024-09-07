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
	SdkLibrary   string `yaml:"SdkLibrary"`   // 留空则不注入微信
}

var Wcf = &IWcf{
	Address: "127.0.0.1:7601",
}

// Web 服务

type IWeb struct {
	Address  string `yaml:"Address"`  // Web 监听地址
	PushUrl  string `yaml:"PushUrl"`  // 消息推送地址
	Storage  string `yaml:"Storage"`  // 附件存储路径
	Swagger  bool   `yaml:"Swagger"`  // 是否启用 Api 文档
	Token    string `yaml:"Token"`    // 使用 Token 验证请求
	UserName string `yaml:"UserName"` // 使用 UserName 验证登录
	PassWord string `yaml:"PassWord"` // 使用 PassWord 验证登录
}

var Web = &IWeb{
	Address: "127.0.0.1:7600",
	Storage: "storage",
	Swagger: true,
}

type IOthers struct {
	SignInPoint     int `yaml:"SignInPoint"`
	PicPoint        int `yaml:"PicPoint"`
	VideoPoint      int `yaml:"VideoPoint"`
	AddPointManager int `yaml:"AddPointManager"`
}

var Others = &IOthers{
	SignInPoint:     10,
	PicPoint:        1,
	VideoPoint:      2,
	AddPointManager: 100,
}

type IFunctionKeyWord struct {
	VideoWord []string `yaml:"videoWord"`
	DogWord   []string `yaml:"dogWord"`
}

var FunctionKeyWord = &IFunctionKeyWord{}

type json struct {
	Code string `yaml:"code"`
	Msg  string `yaml:"msg"`
	Data string `yaml:"data"`
}

type IApiServer struct {
	FishApi          string   `yaml:"fishApi"`
	FishVideoApi     string   `yaml:"fishVideoApi"`
	DogApi           string   `yaml:"dogApi"`
	JokeApi          string   `yaml:"jokeApi"`
	ConstellationApi string   `yaml:"constellationApi"`
	PicApi           []string `yaml:"picApi"`
	VideosApi        []string `yaml:"videosApi"`
}

var ApiServer = &IApiServer{}
