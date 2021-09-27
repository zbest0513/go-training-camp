package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SignalHandle(f func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case msg := <-signals:
		log.Println(fmt.Sprintf("接受信号:%v", msg))
		f()
	}
}
