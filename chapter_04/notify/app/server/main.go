package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		log.Fatalf("项目启动失败:%+v", err)
	}

	engine := gin.New()

	_ = app.Register(engine)

	if err := engine.Run(":8899"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	} else {
		fmt.Println("startup service success\n")
	}

}
