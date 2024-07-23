package wrobot

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/wrest-chat/dbase/forword"
)

type Forword struct{}

// @Summary 消息转发列表
// @Produce json
// @Tags BOT::消息转发
// @Param body body forword.FetchAllParam true "获取消息转发列表参数"
// @Success 200 {array} tables.Forword
// @Router /bot/forword/list [post]
func (*Forword) list(c *gin.Context) {
	var rq *forword.FetchAllParam
}
