package tables

type MpArticle struct {
	Rd        uint   `json:"rd" gorm:"primaryKey;comment:'主键'"` // 主键
	Title     string `json:"title" gorm:"index"`                // 名称
	Desc      string `json:"desc"`
	Url       string `json:"url"`
	PubTime   string `json:"pub_time"`
	Cover     string `json:"cover"`
	Digest    string `json:"digest"`
	Username  string `json:"username" gorm:"index"`               // 公众号id
	Appname   string `json:"appname" gorm:"index"`                // 公众号名称
	CreatedAt int64  `json:"created_at" gorm:"comment:'创建时间戳'"`   // 创建时间戳
	UpdatedAt int64  `json:"updated_at" gorm:"comment:'最后更新时间戳'"` // 最后更新时间戳
}
