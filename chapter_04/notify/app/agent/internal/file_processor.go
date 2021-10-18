package internal

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

type FileProcessor struct {
	path   string
	offset int
}

type Listener struct {
	watcher *fsnotify.Watcher
}

func (receiver *Listener) start(path string) error {
	defer receiver.watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-receiver.watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-receiver.watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := receiver.watcher.Add(path)
	if err != nil {
		return errors.WithMessage(err, "watch file error")
	}
	<-done
	return nil
}

type Parser struct {
}

func readLine() {
	f, err := os.OpenFile("./abc.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	seek, err := f.Seek(0, io.SeekEnd)

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}

	stat, _ := f.Stat()
	wr, err := f.WriteString("9977\n")
	println(fmt.Sprintf("=====:%v,%v", wr, err))
	ret, _ := f.Seek(seek, io.SeekStart)
	println("=========", seek, stat.Size(), ret)
	rd = bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}

}
