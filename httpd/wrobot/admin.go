package wrobot

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/wclient/deliver"
	"net/http"
)

type Admin struct {
}

type AdminParam struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

// 通用结果
type CommonPayload struct {
	// 是否成功
	Success bool `json:"success,omitempty"`
	// 返回结果
	Result string `json:"result,omitempty"`
	// 错误信息
	Error error `json:"error,omitempty"`
}

// @Summary 用户密码登录
// @Produce json
// @Tags WCF::其他
// @Success 200 {object} CommonPayload
// @Router /login/login_admin [post]
func (*Admin) loginAdmin(c *gin.Context) {
	userName := args.Web.UserName
	passWord := args.Web.PassWord
	token := args.Web.Token

	var rq *AdminParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", gin.H{"code": 0, "message": "用户名或密码不能为空"})
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "message": "用户名或密码不能为空"})
		return
	}

	if userName != rq.UserName {
		c.Set("Error", gin.H{"code": 0, "message": "用户名或密码错误"})
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "message": "用户名或密码错误"})
		return
	}
	if passWord != rq.PassWord {
		c.Set("Error", gin.H{"code": 0, "message": "用户名或密码错误"})
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "message": "用户名或密码错误"})
		return
	}
	if token != "" {
		token = deliver.DoubleMD5(token)
	}

	c.Set("Success", gin.H{"code": 200, "message": "登录成功", "data": ""})
	c.JSON(http.StatusBadRequest, gin.H{"code": 200, "message": "登录成功", "token": token})
	return
}
