# Wrest Chat

[![Wrest Chat Builder](https://github.com/opentdp/wrest-chat/actions/workflows/release.yml/badge.svg)](https://github.com/opentdp/wrest-chat/actions/workflows/release.yml)

fork地址：
[![Opentdp]](https://github.com/opentdp/wrest-chat)

智能聊天助手，是一个通用的聊天辅助程序，通过 **Nanomsg 协议** 与聊天软件互通，内置 WEB
管理界面，可接入GPT、Gemini、星火、文心、混元、通义千问等大语言模型。目前已适配 *PC微信*，更多聊天软件适配中，敬请期待！

> 本项目**未对微信程序进行任何破解或修改**
> ，与微信互操作的能力均基于开源项目 [WeChatFerry RPC](https://github.com/lich0821/WeChatFerry/tree/master/WeChatFerry)
> 实现，感谢各位开源贡献者。

## 功能特性

这里仅列举了一些主要的特性，其他信息请参阅[项目文档](https://docs.opentdp.org/#/wrest/)
（by [KincaidYang](https://github.com/KincaidYang)）

- 使用 Go 语言编写，无运行时依赖
- 提供 HTTP 接口，便于对接各类编程语言
- 提供 Websocket 接口，接收推送的新消息
- 支持 HTTP/WS 接口授权，参见 [配置文件解析](https://docs.opentdp.org/#/wrest/配置文件解析)
- 支持作为 SDK 使用，参见 [SDK模块说明](https://docs.opentdp.org/#/wrest/开发指南/SDK模块)
- 内置 AI 机器人，参见 [BOT模块说明](https://docs.opentdp.org/#/wrest/开发指南/BOT模块)
- 内置 Web 管理界面，可以管理机器人各项配置
- 内置 Api 调试工具，所有接口都可以在线调试
- 尽可能将消息中的 Xml 转为 Object，便于前端解析
- 支持计划任务、外部指令、指令插件等扩展功能，详见 [wrest-plugin](https://github.com/opentdp/wrest-plugin)
## wrest使用文档：
https://docs.opentdp.org/#/wrest/README

## 食用方法
### 运行：
```shell
go mod tidy
go run main.go
```

### 打包：
```shell
go build main.go
```
或者 双击 build.bat

默认微信版本：
3.9.2.23


## 代码提交

提交代码时请使用 `feat: something` 作为说明，支持的标识如下

- `feat` 新功能（feature）
- `fix` 错误修复
- `docs` 文档更改（documentation）
- `style` 格式（不影响代码含义的更改，空格、格式、缺少分号等）
- `refactor` 重构（即不是新功能，也不是修补bug的代码变动）
- `perf` 优化（提高性能的代码更改）
- `test` 测试（添加缺失的测试或更正现有测试）
- `chore` 构建过程或辅助工具的变动
- `revert` 还原以前的提交

## 其他
目前使用的是exe调用方式（旧版本）。dll调用方式（新版本）有部分功能没法使用，暂时关闭。

如果使用dll方式调用，请修改：

/wclient/client.go文件夹下：

1. 大概第38行，注释掉。
2. dll文件请放到：/wclient/libs/文件夹下
3. 微信版本使用：39.10.27


## 免责声明

[WrestChatGoHttp](https://github.com/xianrenqh/wrest-chat-gohttp) 、[WrestChat](https://github.com/opentdp/wrest-chat) 和 [WeChatFerry](https://github.com/lich0821/WeChatFerry)
是供学习交流的开源项目，代码及其制品仅供参考，不保证质量，不构成任何商业承诺或担保，不得用于商业或非法用途，使用者自行承担后果。


License [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.txt)

Copyright (c) 2022 - 2024 OpenTDP
