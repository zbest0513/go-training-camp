package internal

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
)

type ConfigProcessor struct {
	vip        *viper.Viper
	path       string
	watcher    *fsnotify.Watcher
	logPaths   map[string]string
	isRun      bool
	changeChan *chan bool
}

func NewConfigProcessor(changeChan *chan bool) *ConfigProcessor {
	v := viper.New()
	v.SetConfigName("fileconfig")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(fmt.Errorf("Fatal error watch config file: %s \n", err))
	}
	c := new(ConfigProcessor)
	c.vip = v
	c.watcher = watcher
	c.path = "configs/fileconfig.yaml"
	c.build()
	c.isRun = true
	c.changeChan = changeChan
	go c.start()
	return c
}

func (receiver *ConfigProcessor) build() {
	receiver.logPaths = receiver.vip.GetStringMapString("logs")
	log.Println(fmt.Sprintf("build path map: %v", receiver.logPaths))
}

func (receiver *ConfigProcessor) Destroy() {
	receiver.isRun = false
}

func (receiver *ConfigProcessor) start() error {
	defer receiver.watcher.Close()
	done := make(chan bool)
	go func() {
		defer close(done)
		//done <- true
		for receiver.isRun {
			select {
			case event, ok := <-receiver.watcher.Events:
				if !ok {
					return
				}
				if event.Op == fsnotify.Write || event.Op == fsnotify.Rename {
					receiver.vip.ReadInConfig()
					receiver.build()
					*receiver.changeChan <- true
				} else if event.Op == fsnotify.Remove {
					receiver.watcher.Add(receiver.path)
					log.Println("重新 add watcher ...")
				}
			}
		}
	}()
	err := receiver.watcher.Add(receiver.path)
	if err != nil {
		return errors.WithMessage(err, "watch file error")
	}
	<-done
	return nil
}

func (receiver *ConfigProcessor) QueryAddr() string {
	return receiver.vip.GetString("addr")
}
