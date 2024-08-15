package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入目录

var Efs *embed.FS

// 版本信息

const Version = "0.26.3"
const BuildVersion = "240106"

// 应用描述

const AppName = "Hui Robot"
const AppSummary = "智能聊天助手"
