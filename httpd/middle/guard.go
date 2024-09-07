package middle

import (
	"github.com/opentdp/wrest-chat/wclient/deliver"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/args"
)

func ApiGuard(c *gin.Context) {

	token := ""

	// 取回 Token
	authcode := c.GetHeader("Authorization")
	parts := strings.SplitN(authcode, " ", 2)
	if len(parts) == 2 && parts[0] == "Bearer" {
		token = parts[1]
	} else {
		token = c.Query("token")
	}

	//token增加md5验证
	md5Token := deliver.DoubleMD5(args.Web.Token)
	// 校验 Token
	if token != md5Token {
		c.Set("Error", gin.H{"Code": 401, "Message": "操作未授权"})
		c.Set("ExitCode", 401)
		c.Abort()
	}

}

func SwaggerGuard(c *gin.Context) {

	if !args.Web.Swagger && strings.HasPrefix(c.Request.URL.Path, "/swagger") {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, "功能已禁用")
		c.Abort()
	}

}
