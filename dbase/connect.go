package dbase

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/keyword"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/dbase/tables"
	"github.com/rs/zerolog/log"
)

func Connect() {
	log.Info().Msg("正在连接数据库，请稍后...")
	dbname := "wrest.db3"
	if !filer.Exists(dbname) {
		dbname = args.Web.Storage + "/" + dbname
	}

	// 连接数据库
	db := dborm.Connect(&dborm.Config{
		Type:   "sqlite",
		DbName: dbname,
	})

	// 实施自动迁移
	db.AutoMigrate(
		&tables.Chatroom{},
		&tables.Cronjob{},
		&tables.Contact{},
		&tables.Keyword{},
		&tables.LLModel{},
		&tables.Message{},
		&tables.Profile{},
		&tables.Setting{},
		&tables.Webhook{},
		&tables.Forward{},
		&tables.Point{},
		&tables.PointList{},
		&tables.SystemCrontab{},
		&tables.SystemCrontabLog{},
		&tables.MpArticle{},
	)

	// 开启外键约束
	db.Exec("PRAGMA foreign_keys=ON;")

	// 加载全局配置
	setting.DataMigrate()
	setting.Laod()
	keyword.DataMigrate()
	log.Info().Msg("数据库连接成功")
}
