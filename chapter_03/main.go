package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func server(ctx context.Context, wg *sync.WaitGroup, port string) {
	defer wg.Done()
	go http.ListenAndServe(port, nil)
	log.Println(fmt.Sprintf("listen port :%s", port))
	select {
	case <-ctx.Done():
		log.Println(fmt.Sprintf("listen port :%s done ......", port))
	}
}
func signalHandle(f func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case msg := <-signals:
		log.Println(fmt.Sprintf("接受信号:%v", msg))
		f()
	}
}

func main() {
	WaitGroupVersion() //wait group 实现的版本
}

func WaitGroupVersion() {
	var wg sync.WaitGroup
	ctxA, cancelA := context.WithCancel(context.Background())
	defer cancelA()
	wg.Add(1)
	go server(ctxA, &wg, ":8001")
	wg.Add(1)
	go server(ctxA, &wg, ":8082")
	go signalHandle(func() {
		cancelA()
	})
	wg.Wait()
}
