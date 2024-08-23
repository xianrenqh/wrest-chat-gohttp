package keyword

// 迁移数据
func DataMigrate() {
	settings := []*CreateParam{
		{0, "handler", "*", "菜单", -1, "/help", ""},
		{0, "handler", "*", "美女", -1, "美女图片", ""},
		{0, "handler", "*", "图片", -1, "美女图片", ""},
		{0, "handler", "*", "妹子", -1, "美女图片", ""},
		{0, "handler", "*", "视频", -1, "美女视频", ""},
	}
	for _, item := range settings {
		if _, err := Fetch(&FetchParam{Group: item.Group, Phrase: item.Phrase, Roomid: item.Roomid}); err != nil {
			Create(item)
		}
	}
}
