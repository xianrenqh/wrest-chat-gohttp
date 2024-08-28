//go:build !windows

package wcferry

import (
	"github.com/rs/zerolog/log"
)

// 调用 sdk.dll 中的函数
// return error 错误信息
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	log.Warn().Str("fn", fn, "a", a).Msg("跳过加载sdk.dll")
	return nil
}
