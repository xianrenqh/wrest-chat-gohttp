package args

import (
	"fmt"
	"os"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
)

var Configer *Config

type ConfigData struct {
	Log             *ILog             `yaml:"Log"`             // 日志
	Wcf             *IWcf             `yaml:"Wcf"`             // Wcf 服务
	Web             *IWeb             `yaml:"Web"`             // Web 服务
	Others          *IOthers          `yaml:"Others"`          // 其他配置
	FunctionKeyWord *IFunctionKeyWord `yaml:"FunctionKeyWord"` //
	ApiServer       *IApiServer       `yaml:"ApiServer"`       //
}

func init() {

	fmt.Println(AppName, AppSummary)
	fmt.Println("Version:", Version, "build", BuildVersion)

	fmt.Println("┌───┐   ┌───┬───┬───┬───┐ ┌───┬───┬───┬───┐ ┌───┬───┬───┬───┐ ┌───┬───┬───┐\n│Esc│   │ F1│ F2│ F3│ F4│ │ F5│ F6│ F7│ F8│ │ F9│F10│F11│F12│ │P/S│S L│P/B│  ┌┐    ┌┐    ┌┐\n└───┘   └───┴───┴───┴───┘ └───┴───┴───┴───┘ └───┴───┴───┴───┘ └───┴───┴───┘  └┘    └┘    └┘\n┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───────┐ ┌───┬───┬───┐ ┌───┬───┬───┬───┐\n│~ `│! 1│@ 2│# 3│$ 4│% 5│^ 6│& 7│* 8│( 9│) 0│_ -│+ =│ BacSp │ │Ins│Hom│PUp│ │N L│ / │ * │ - │\n├───┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─────┤ ├───┼───┼───┤ ├───┼───┼───┼───┤\n│ Tab │ Q │ W │ E │ R │ T │ Y │ U │ I │ O │ P │{ [│} ]│ | \\ │ │Del│End│PDn│ │ 7 │ 8 │ 9 │   │\n├─────┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴─────┤ └───┴───┴───┘ ├───┼───┼───┤ + │\n│ Caps │ A │ S │ D │ F │ G │ H │ J │ K │ L │: ;│\" '│ Enter  │               │ 4 │ 5 │ 6 │   │\n├──────┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴────────┤     ┌───┐     ├───┼───┼───┼───┤\n│ Shift  │ Z │ X │ C │ V │ B │ N │ M │< ,│> .│? /│  Shift   │     │ ↑ │     │ 1 │ 2 │ 3 │   │\n├─────┬──┴─┬─┴──┬┴───┴───┴───┴───┴───┴──┬┴───┼───┴┬────┬────┤ ┌───┼───┼───┐ ├───┴───┼───┤ E││\n│ Ctrl│    │Alt │         Space         │ Alt│    │    │Ctrl│ │ ← │ ↓ │ → │ │   0   │ . │←─┘│\n└─────┴────┴────┴───────────────────────┴────┴────┴────┴────┘ └───┴───┴───┘ └───────┴───┴───┘\n")
	fmt.Println("\n\n  顶顶顶顶顶顶顶顶顶　顶顶顶顶顶顶顶顶顶\n  顶顶顶顶顶顶顶　　　　　顶顶　　　　　\n  　　　顶顶　　　顶顶顶顶顶顶顶顶顶顶顶\n  　　　顶顶　　　顶顶顶顶顶顶顶顶顶顶顶\n  　　　顶顶　　　顶顶　　　　　　　顶顶\n  　　　顶顶　　　顶顶　　顶顶顶　　顶顶\n  　　　顶顶　　　顶顶　　顶顶顶　　顶顶\n  　　　顶顶　　　顶顶　　顶顶顶　　顶顶\n  　　　顶顶　　　顶顶　　顶顶顶　　顶顶\n  　　　顶顶　　　　　　　顶顶顶　\n  　　　顶顶　　　　　　顶顶　顶顶　顶顶\n  　顶顶顶顶　　　顶顶顶顶顶　顶顶顶顶顶\n  　顶顶顶顶　　　顶顶顶顶　　　顶顶顶顶\n \n")
	// 调试模式

	de := os.Getenv("TDP_DEBUG")
	Debug = de == "1" || de == "true"

	// 初始化配置

	Configer = &Config{
		File: "config.yml",
		Data: &ConfigData{Log, Wcf, Web, Others, FunctionKeyWord, ApiServer},
	}

	if len(os.Args) > 1 {
		Configer.File = os.Args[1]
	}

	if err := Configer.Load(); err != nil {
		panic(err)
	}

	// 初始化存储

	if !filer.Exists(Web.Storage) {
		os.MkdirAll(Web.Storage, 0755)
	}

	// 初始化日志

	logman.SetDefault(&logman.Config{
		Level:    Log.Level,
		Target:   Log.Target,
		Storage:  Log.Dir,
		Filename: "common",
	})

}
