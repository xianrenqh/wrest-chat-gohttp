package wrobot

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/wrest-chat/dbase/forward"
	"github.com/opentdp/wrest-chat/wclient/robot"
)

type Forward struct{}

// @Summary 消息转发列表
// @Produce json
// @Tags BOT::消息转发
// @Param body body forward.FetchAllParam true "获取消息转发列表参数"
// @Success 200 {array} tables.Forward
// @Router /bot/forward/list [post]
func (*Forward) list(c *gin.Context) {
	var rq *forward.FetchAllParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}
	if lst, err := forward.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}
}

// @Summary 获取消息转发配置
// @Produce json
// @Tags BOT::消息转发
// @Param body body forward.FetchParam true "获取消息转发参数"
// @Success 200 {object} tables.Forward
// @Router /bot/forward/detail [post]
func (*Forward) detail(c *gin.Context) {

	var rq *forward.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := forward.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加消息转发配置
// @Produce json
// @Tags BOT::消息转发
// @Param body body forward.CreateParam true "添加消息转发配置参数"
// @Success 200
// @Router /bot/forward/create [post]
func (*Forward) create(c *gin.Context) {
	var rq *forward.CreateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := forward.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
		robot.Reset()
	} else {
		c.Set("Error", err)
	}
}

// @Summary 修改消息转发配置
// @Produce json
// @Tags BOT::消息转发
// @Param body body forward.UpdateParam true "修改消息转发参数"
// @Success 200
// @Router /bot/forward/update [post]
func (*Forward) update(c *gin.Context) {

	var rq *forward.UpdateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := forward.Update(rq); err == nil {
		c.Set("Message", "更新成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除消息转发配置
// @Produce json
// @Tags BOT::消息转发
// @Param body body forward.DeleteParam true "删除消息转发参数"
// @Success 200
// @Router /bot/forward/delete [post]
func (*Forward) delete(c *gin.Context) {

	var rq *forward.DeleteParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := forward.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}
