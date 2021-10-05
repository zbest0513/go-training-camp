package main

import (
	"log"
	_ "notify/doc"
	"notify/pkg/config"
)

// @title notify-server
// @version 1.0
// @description notify-server api 文档

// @contact.name zbest
// @contact.url http://wiki.zbest.tech
// @contact.email zbest0513@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	app, err := initApp(config.FileConfig{
		Path:     "./configs",
		Name:     "config",
		FileType: "yaml",
	})
	if err != nil {
		log.Fatalf("配置文件加载失败:%+v", err)
	}
	err = app.Run()
	if err != nil {
		log.Fatalf("项目启动失败:%+v", err)
	}
}
