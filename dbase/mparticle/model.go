package mparticle

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

type CreateParam struct {
	Rd       uint   `json:"rd" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"index" binding:"required"`
	Desc     string `json:"desc"`
	Url      string `json:"url"`
	PubTime  string `json:"pub_time"`
	Cover    string `json:"cover"`
	Digest   string `json:"digest"`
	Username string `json:"username" gorm:"index"`
	Appname  string `json:"appname"`
}

func Create(data *CreateParam) (uint, error) {
	item := &tables.MpArticle{
		Title:    data.Title,
		Desc:     data.Desc,
		Url:      data.Url,
		PubTime:  data.PubTime,
		Cover:    data.Cover,
		Digest:   data.Digest,
		Username: data.Username,
		Appname:  data.Appname,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error
}
