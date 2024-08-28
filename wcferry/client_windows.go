//go:build windows

package wcferry

import (
	"errors"
	"github.com/rs/zerolog/log"
	"syscall"

	"github.com/opentdp/go-helper/filer"
)

// 调用 sdk.dll 中的函数
// return error 错误信息
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	if c.SdkLibrary == "" {
		log.Warn().Msg("跳过加载 sdk.dll")
		return nil
	}
	// 查找 sdk.dll
	dll := c.SdkLibrary
	if !filer.Exists(dll) {
		dll = "wcferry/" + dll
		if !filer.Exists(dll) {
			return errors.New(dll + " not found")
		}
	}
	// 加载 sdk.dll
	sdk, err := syscall.LoadDLL(dll)
	if err != nil {
		log.Warn().Err(err).Msg("sdk.dll加载失败，错误：")
		return err
	}
	defer sdk.Release()
	// 查找 fn 函数
	proc, err := sdk.FindProc(fn)
	if err != nil {
		log.Warn().Err(err).Msg("failed to call " + fn)
		return err
	}
	// 执行 fn(a...)
	r1, r2, err := proc.Call(a...)
	log.Warn().Msgf("%s(%v) -> %v %v %v", fn, a, r1, r2, err)
	if err.Error() == "Attempt to access invalid address." {
		err = nil // 忽略已知问题
	}
	return err
}
