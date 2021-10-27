package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type AgentProcessor struct {
	logsContainer map[string]*LogProcessor
	changeChan    chan bool
	config        *ConfigProcessor
	isStart       chan bool
}

func NewAgentProcessor() *AgentProcessor {
	ap := new(AgentProcessor)
	c := make(chan bool)
	c2 := make(chan bool)
	ap.changeChan = c
	ap.isStart = c2
	ap.config = NewConfigProcessor(&c)
	ap.logsContainer = make(map[string]*LogProcessor)
	go func() {
		for {
			<-c
			ap.build()
		}
	}()
	return ap
}

func (receiver *AgentProcessor) startListener() {
	ticker := time.NewTicker(5 * time.Second)
	stopChan := make(chan bool)
	defer close(stopChan)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:

				receiver.savePoint()
			case stop := <-stopChan:
				if stop {
					log.Println("退出offset 保存")
					return
				}
			}
		}
	}(ticker)
	<-receiver.isStart
	stopChan <- true
}

func (receiver *AgentProcessor) Start() {
	receiver.build()
	receiver.startListener()
}

func (receiver *AgentProcessor) readPoint(key string) int64 {
	vip := viper.New()
	vip.SetConfigName("pointconfig")
	vip.SetConfigType("yaml")
	vip.AddConfigPath("configs")
	err := vip.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Sprintf("读取offset失败:%v", err))
	}
	getInt64 := vip.GetInt64(key)
	log.Println(fmt.Sprintf("读取配置:%v:%v", key, getInt64))
	return getInt64
}

func (receiver *AgentProcessor) savePoint() {
	log.Println("保存offset....")
	vip := viper.New()
	vip.SetConfigName("pointconfig")
	vip.SetConfigType("yaml")
	vip.AddConfigPath("configs")

	for k, v := range receiver.logsContainer {
		vip.Set(k, v.offset)
	}
	vip.WriteConfig()
}

func (receiver *AgentProcessor) Destroy() {
	defer close(receiver.isStart)
	for _, v := range receiver.logsContainer {
		v.Destroy()
	}
	receiver.config.Destroy()
	receiver.savePoint()
	receiver.isStart <- true
}

func (receiver *AgentProcessor) build() {
	pathmap := receiver.config.logPaths
	for k, v := range pathmap {
		if receiver.logsContainer[k] == nil {
			log.Println(fmt.Sprintf("加入监听:%v", v))
			receiver.logsContainer[k] = CreateLogProcessor(v, 40, receiver.readPoint(k), receiver.config.QueryAddr(), k)
		}
	}
}
