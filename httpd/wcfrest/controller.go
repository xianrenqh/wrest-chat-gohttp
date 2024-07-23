package wcfrest

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"

	"github.com/opentdp/wrest-chat/wcferry"
)

type Controller struct {
	*wcferry.Client
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

// WS 协议升级
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Summary 登录二维码
// @Produce json
// @Tags WCF::其他
// @Success 200 {object} CommonPayload
// @Router /wcf/login_qr [post]
func (wc *Controller) loginQr(c *gin.Context) {

	url := wc.CmdClient.RefreshQrcode()

	c.Set("Payload", CommonPayload{
		Success: url != "",
		Result:  url,
	})

}

// @Summary 检查登录状态
// @Produce json
// @Tags WCF::其他
// @Success 200 {object} bool
// @Router /wcf/is_login [post]
func (wc *Controller) isLogin(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.IsLogin())

}

// @Summary 获取登录账号wxid
// @Produce json
// @Tags WCF::其他
// @Success 200 {object} string
// @Router /wcf/self_wxid [post]
func (wc *Controller) getSelfWxid(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetSelfWxid())

}

// @Summary 获取登录账号个人信息
// @Produce json
// @Tags WCF::其他
// @Success 200 {object} UserInfoPayload
// @Router /wcf/self_info [post]
func (wc *Controller) getSelfInfo(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetSelfInfo())

}

type UserInfoPayload struct {
	// 用户 id
	Wxid string `json:"wxid,omitempty"`
	// 昵称
	Name string `json:"name,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 文件/图片等父路径
	Home string `json:"home,omitempty"`
}

// @Summary 获取所有消息类型
// @Produce json
// @Tags WCF::其他
// @Success 200 {object} map[int32]string
// @Router /wcf/msg_types [post]
func (wc *Controller) getMsgTypes(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetMsgTypes())

}

// @Summary 获取数据库列表
// @Produce json
// @Tags WCF::数据库查询
// @Success 200 {object} []string
// @Router /wcf/db_names [post]
func (wc *Controller) getDbNames(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetDbNames())

}

// @Summary 获取数据库表列表
// @Produce json
// @Tags WCF::数据库查询
// @Param body body GetDbTablesRequest true "获取数据库表列表参数"
// @Success 200 {object} []DbTablePayload
// @Router /wcf/db_tables [post]
func (wc *Controller) getDbTables(c *gin.Context) {

	var req GetDbTablesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetDbTables(req.Db))

}

type DbTablePayload struct {
	// 表名
	Name string `json:"name,omitempty"`
	// 建表 SQL
	Sql string `json:"sql,omitempty"`
}

type GetDbTablesRequest struct {
	// 数据库名称
	Db string `json:"db"`
}

// @Summary 执行数据库查询
// @Produce json
// @Tags WCF::数据库查询
// @Param body body DbSqlQueryRequest true "数据库查询参数"
// @Success 200 {object} []map[string]any
// @Router /wcf/db_query_sql [post]
func (wc *Controller) dbSqlQuery(c *gin.Context) {

	var req DbSqlQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.DbSqlQuery(req.Db, req.Sql))

}

type DbSqlQueryRequest struct {
	// 数据库名称
	Db string `json:"db"`
	// 待执行的 SQL
	Sql string `json:"sql"`
}

// @Summary 获取群列表
// @Produce json
// @Tags WCF::群聊管理
// @Success 200 {object} []ContactPayload
// @Router /wcf/chatrooms [post]
func (wc *Controller) getChatRooms(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetChatRooms())

}

// @Summary 获取群成员列表
// @Produce json
// @Tags WCF::群聊管理
// @Param body body GetChatRoomMembersRequest true "获取群成员列表参数"
// @Success 200 {object} []ContactPayload
// @Router /wcf/chatroom_members [post]
func (wc *Controller) getChatRoomMembers(c *gin.Context) {

	var req GetChatRoomMembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetChatRoomMembers(req.Roomid))

}

type GetChatRoomMembersRequest struct {
	// 群聊 id
	Roomid string `json:"roomid"`
}

// @Summary 获取群成员昵称
// @Produce json
// @Tags WCF::群聊管理
// @Param body body GetAliasInChatRoomRequest true "获取群成员昵称参数"
// @Success 200 {object} string
// @Router /wcf/alias_in_chatroom [post]
func (wc *Controller) getAliasInChatRoom(c *gin.Context) {

	var req GetAliasInChatRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetAliasInChatRoom(req.Wxid, req.Roomid))

}

type GetAliasInChatRoomRequest struct {
	// 群聊 id
	Roomid string `json:"roomid"`
	// 用户 id
	Wxid string `json:"wxid"`
}

// @Summary 邀请群成员
// @Produce json
// @Tags WCF::群聊管理
// @Param body body ChatroomMembersRequest true "管理群成员参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/invite_chatroom_members [post]
func (wc *Controller) inviteChatroomMembers(c *gin.Context) {

	var req ChatroomMembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.InviteChatroomMembers(req.Roomid, strings.Join(req.Wxids, ","))

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type ChatroomMembersRequest struct {
	// 群聊 id
	Roomid string `json:"roomid"`
	// 用户 id 列表
	Wxids []string `json:"wxids"`
}

// @Summary 添加群成员
// @Produce json
// @Tags WCF::群聊管理
// @Param body body ChatroomMembersRequest true "管理群成员参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/add_chatroom_members [post]
func (wc *Controller) addChatRoomMembers(c *gin.Context) {

	var req ChatroomMembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.AddChatRoomMembers(req.Roomid, strings.Join(req.Wxids, ","))

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 删除群成员
// @Produce json
// @Tags WCF::群聊管理
// @Param body body ChatroomMembersRequest true "管理群成员参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/del_chatroom_members [post]
func (wc *Controller) delChatRoomMembers(c *gin.Context) {

	var req ChatroomMembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.DelChatRoomMembers(req.Roomid, strings.Join(req.Wxids, ","))

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

// @Summary 撤回消息
// @Produce json
// @Tags WCF::消息发送
// @Param body body RevokeMsgRequest true "撤回消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/revoke_msg [post]
func (wc *Controller) revokeMsg(c *gin.Context) {

	var req RevokeMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.RevokeMsg(req.Msgid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type RevokeMsgRequest struct {
	// 消息 id
	Msgid uint64 `json:"msgid"`
}

// @Summary 转发消息
// @Produce json
// @Tags WCF::消息发送
// @Param body body ForwardMsgRequest true "转发消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/forward_msg [post]
func (wc *Controller) forwardMsg(c *gin.Context) {

	var req ForwardMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	receiverList := strings.Join(req.Receiver, ",")

	var status bool
	if len(req.Receiver) > 1 {
		for _, receive := range req.Receiver {
			if wc.CmdClient.ForwardMsg(req.Id, receive) == 1 {
				status = true
			} else {
				status = false
				break // 如果有一个失败，可以提前退出
			}
		}
	} else {
		status = wc.CmdClient.ForwardMsg(req.Id, receiverList) == 1
	}

	c.Set("Payload", CommonPayload{
		Success: status,
	})

}

type ForwardMsgRequest struct {
	// 待转发消息 id
	Id uint64 `json:"id"`
	// 转发接收人或群的 id 列表
	Receiver []string `json:"receiver"`
}

// @Summary 发送文本消息
// @Produce json
// @Tags WCF::消息发送
// @Param body body SendTxtRequest true "发送文本消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/send_txt [post]
func (wc *Controller) sendTxt(c *gin.Context) {

	var req SendTxtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendTxt(req.Msg, req.Receiver, strings.Join(req.Aters, ","))

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

type SendTxtRequest struct {
	// 消息内容
	Msg string `json:"msg"`
	// 接收人或群的 id
	Receiver string `json:"receiver"`
	// 需要 At 的用户 id 列表
	Aters []string `json:"aters"`
}

// @Summary 发送图片消息
// @Produce json
// @Tags WCF::消息发送
// @Param body body SendFileRequest true "发送图片消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/send_img [post]
func (wc *Controller) sendImg(c *gin.Context) {

	var req SendImgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	// 将 Base64 写入文件
	if req.Base64 != "" {
		fileData, err := base64.StdEncoding.DecodeString(req.Base64)
		if err != nil {
			c.Set("Error", err)
			return
		}
		if err := os.WriteFile(req.Path, fileData, 0644); err != nil {
			c.Set("Error", err)
			return
		}
	}

	status := wc.CmdClient.SendImg(req.Path, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

type SendImgRequest = SendFileRequest

// @Summary 发送文件消息
// @Produce json
// @Tags WCF::消息发送
// @Param body body SendFileRequest true "发送文件消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/send_file [post]
func (wc *Controller) sendFile(c *gin.Context) {

	var req SendFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	// 将 Base64 数据写入文件
	if req.Base64 != "" {
		fileData, err := base64.StdEncoding.DecodeString(req.Base64)
		if err != nil {
			c.Set("Error", err)
			return
		}
		if err := os.WriteFile(req.Path, fileData, 0644); err != nil {
			c.Set("Error", err)
			return
		}
	}

	status := wc.CmdClient.SendFile(req.Path, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

type SendFileRequest struct {
	// 文件路径，若提供 base64 则写入此路径
	Path string `json:"path"`
	// 文件 base64 数据
	Base64 string `json:"base64"`
	// 接收人或群的 id
	Receiver string `json:"receiver"`
}

// @Summary 发送卡片消息
// @Produce json
// @Tags WCF::消息发送
// @Param body body SendRichTextRequest true "发送卡片消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/send_rich_text [post]
func (wc *Controller) sendRichText(c *gin.Context) {

	var req SendRichTextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendRichText(req.Name, req.Account, req.Title, req.Digest, req.Url, req.Thumburl, req.Receiver)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

type SendRichTextRequest struct {
	// 左下显示的名字
	Name string `json:"name"`
	// 填公众号 id 可以显示对应的头像（gh_ 开头的）
	Account string `json:"account"`
	// 标题，最多两行
	Title string `json:"title"`
	// 摘要，三行
	Digest string `json:"digest"`
	// 点击后跳转的链接
	Url string `json:"url"`
	// 缩略图的链接
	Thumburl string `json:"thumburl"`
	// 接收人或群的 id
	Receiver string `json:"receiver"`
}

// @Summary 拍一拍群友
// @Produce json
// @Tags WCF::消息发送
// @Param body body SendPatMsgRequest true "拍一拍群友参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/send_pat_msg [post]
func (wc *Controller) sendPatMsg(c *gin.Context) {

	var req SendPatMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.SendPatMsg(req.Roomid, req.Wxid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type SendPatMsgRequest struct {
	// 群 id
	Roomid string `json:"roomid"`
	// 用户 id
	Wxid string `json:"wxid"`
}

// @Summary 获取语音消息
// @Produce json
// @Tags WCF::消息收取
// @Param body body GetAudioMsgRequest true "获取语音消息参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/audio_msg [post]
func (wc *Controller) getAudioMsg(c *gin.Context) {

	var req GetAudioMsgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	if req.Timeout > 0 {
		resp, err := wc.CmdClient.GetAudioMsgTimeout(req.Msgid, req.Dir, req.Timeout)
		c.Set("Payload", CommonPayload{
			Success: resp != "",
			Result:  resp,
			Error:   err,
		})
	} else {
		resp := wc.CmdClient.GetAudioMsg(req.Msgid, req.Dir)
		c.Set("Payload", CommonPayload{
			Success: resp != "",
			Result:  resp,
		})
	}

}

type GetAudioMsgRequest struct {
	// 消息 id
	Msgid uint64 `json:"msgid"`
	// 存储路径
	Dir string `json:"path"`
	// 超时重试次数
	Timeout int `json:"timeout"`
}

// @Summary 获取OCR识别结果
// @Produce json
// @Tags WCF::消息收取
// @Param body body GetOcrRequest true "获取OCR识别结果参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/ocr_result [post]
func (wc *Controller) getOcrResult(c *gin.Context) {

	var req GetOcrRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	if req.Timeout > 0 {
		resp, err := wc.CmdClient.GetOcrResultTimeout(req.Extra, req.Timeout)
		c.Set("Payload", CommonPayload{
			Success: resp != "",
			Result:  resp,
			Error:   err,
		})
	} else {
		resp, stat := wc.CmdClient.GetOcrResult(req.Extra)
		c.Set("Payload", CommonPayload{
			Success: stat == 0,
			Result:  resp,
		})
	}

}

type GetOcrRequest struct {
	// 消息中的 extra 字段
	Extra string `json:"extra"`
	// 超时重试次数
	Timeout int `json:"timeout"`
}

// @Summary 下载图片
// @Produce json
// @Tags WCF::消息收取
// @Param body body DownloadImageRequest true "下载图片参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/download_image [post]
func (wc *Controller) downloadImage(c *gin.Context) {

	var req DownloadImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	resp, err := wc.CmdClient.DownloadImage(req.Msgid, req.Extra, req.Dir, req.Timeout)

	c.Set("Payload", CommonPayload{
		Success: resp != "",
		Result:  resp,
		Error:   err,
	})

}

type DownloadImageRequest struct {
	// 消息 id
	Msgid uint64 `json:"msgid"`
	// 消息中的 extra 字段
	Extra string `json:"extra"`
	// 存储路径
	Dir string `json:"dir"`
	// 超时重试次数
	Timeout int `json:"timeout"`
}

// @Summary 下载附件
// @Produce json
// @Tags WCF::消息收取
// @Param body body DownloadAttachRequest true "下载附件参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/download_attach [post]
func (wc *Controller) downloadAttach(c *gin.Context) {

	var req DownloadAttachRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.DownloadAttach(req.Msgid, req.Thumb, req.Extra)

	c.Set("Payload", CommonPayload{
		Success: status == 0,
	})

}

type DownloadAttachRequest struct {
	// 消息 id
	Msgid uint64 `json:"msgid"`
	// 消息中的 thumb 字段
	Thumb string `json:"thumb"`
	// 消息中的 extra 字段
	Extra string `json:"extra"`
}

// @Summary 获取头像列表
// @Produce json
// @Tags WCF::联系人管理
// @Param body body GetAvatarsRequest true "获取头像列表参数"
// @Success 200 {object} []AvatarPayload
// @Router /wcf/avatars [post]
func (wc *Controller) getAvatars(c *gin.Context) {

	var req GetAvatarsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	sql := "SELECT usrName as UsrName, bigHeadImgUrl as BigHeadImgUrl, smallHeadImgUrl as SmallHeadImgUrl FROM ContactHeadImgUrl"

	if len(req.Wxids) > 0 {
		for i, v := range req.Wxids {
			req.Wxids[i] = strings.ReplaceAll(v, "'", "''")
		}
		sql += " WHERE usrName IN ('" + strings.Join(req.Wxids, "','") + "')"
	}

	res := wc.CmdClient.DbSqlQuery("MicroMsg.db", sql)

	var result []AvatarPayload
	if mapstructure.Decode(res, &result) == nil {
		c.Set("Payload", result)
	} else {
		c.Set("Payload", res)
	}

}

type GetAvatarsRequest struct {
	// 用户 id 列表
	Wxids []string `json:"wxids"`
}

type AvatarPayload struct {
	// 用户 id
	UsrName string `json:"usr_name,omitempty"`
	// 大头像 url
	BigHeadImgUrl string `json:"big_head_img_url,omitempty"`
	// 小头像 url
	SmallHeadImgUrl string `json:"small_head_img_url,omitempty"`
}

// @Summary 获取完整通讯录
// @Produce json
// @Tags WCF::联系人管理
// @Success 200 {object} []ContactPayload
// @Router /wcf/contacts [post]
func (wc *Controller) getContacts(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetContacts())

}

type ContactPayload struct {
	// 用户 id
	Wxid string `json:"wxid,omitempty"`
	// 微信号
	Code string `json:"code,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 昵称
	Name string `json:"name,omitempty"`
	// 国家
	Country string `json:"country,omitempty"`
	// 省/州
	Province string `json:"province,omitempty"`
	// 城市
	City string `json:"city,omitempty"`
	// 性别
	Gender int32 `json:"gender,omitempty"`
}

// @Summary 获取好友列表
// @Produce json
// @Tags WCF::联系人管理
// @Success 200 {object} []ContactPayload
// @Router /wcf/friends [post]
func (wc *Controller) getFriends(c *gin.Context) {

	c.Set("Payload", wc.CmdClient.GetFriends())

}

// @Summary 根据wxid获取个人信息
// @Produce json
// @Tags WCF::联系人管理
// @Param body body GetInfoByWxidRequest true "根据wxid获取个人信息参数"
// @Success 200 {object} ContactPayload
// @Router /wcf/user_info [post]
func (wc *Controller) getInfoByWxid(c *gin.Context) {

	var req GetInfoByWxidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", wc.CmdClient.GetInfoByWxid(req.Wxid))

}

type GetInfoByWxidRequest struct {
	// 用户 id
	Wxid string `json:"wxid"`
}

// @Summary 刷新朋友圈
// @Produce json
// @Tags WCF::其他
// @Param body body RefreshPyqRequest true "刷新朋友圈参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/refresh_pyq [post]
func (wc *Controller) refreshPyq(c *gin.Context) {

	var req RefreshPyqRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.RefreshPyq(req.Id)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type RefreshPyqRequest struct {
	// 分页 id
	Id uint64 `json:"id"`
}

// @Summary 接受好友请求
// @Produce json
// @Tags WCF::联系人管理
// @Param body body AcceptNewFriendRequest true "接受好友参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/accept_new_friend [post]
func (wc *Controller) acceptNewFriend(c *gin.Context) {

	var req AcceptNewFriendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.AcceptNewFriend(req.V3, req.V4, req.Scene)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type AcceptNewFriendRequest struct {
	// 加密的用户名
	V3 string `json:"v3"`
	// 验证信息 Ticket
	V4 string `json:"v4"`
	// 添加方式：17 名片，30 扫码
	Scene int32 `json:"scene"`
}

// @Summary 接受转账
// @Produce json
// @Tags WCF::消息收取
// @Param body body ReceiveTransferRequest true "接受转账参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/receive_transfer [post]
func (wc *Controller) receiveTransfer(c *gin.Context) {

	var req ReceiveTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	status := wc.CmdClient.ReceiveTransfer(req.Wxid, req.Tfid, req.Taid)

	c.Set("Payload", CommonPayload{
		Success: status == 1,
	})

}

type ReceiveTransferRequest struct {
	// 转账人
	Wxid string `json:"wxid,omitempty"`
	// 转账id transferid
	Tfid string `json:"tfid,omitempty"`
	// Transaction id
	Taid string `json:"taid,omitempty"`
}
