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
	ap.changeChan = c
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
	t := time.Tick(5 * time.Second)
	//100ms 生成一个新桶
	go func() {
		for {
			log.Println("保存offset....")
			<-t
			go func() {
				receiver.savePoint()
			}()
		}
	}()
}

func (receiver *AgentProcessor) Start() {
	receiver.build()
	receiver.startListener()
	<-receiver.isStart
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
	for _, v := range receiver.logsContainer {
		v.Destroy()
	}
	receiver.config.Destroy()
	receiver.savePoint()
	receiver.isStart <- false
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
