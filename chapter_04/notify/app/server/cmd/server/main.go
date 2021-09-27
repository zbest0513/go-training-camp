package main

import (
	"log"
	"notify/pkg/config"
)

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
