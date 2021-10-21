package internal

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"regexp"
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

func ReadLines(path string, count int, idx int64) (int64, error) {
	f, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	seek, err := f.Seek(idx, io.SeekStart)
	rd := bufio.NewReader(f)
	var step int64
	i := 0
	var flag = false
	for count > i || flag {
		i++
		line, err := rd.ReadString('\n')
		step += int64(len(line))
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
		exception, _ := regexp.Compile(`\w+[.]\w+Exception:`)
		match := exception.MatchString(line)
		if match {
			flag = true
			continue
		}
		at, _ := regexp.Compile(`[\t]at `)
		match = at.MatchString(line)
		if !match {
			flag = false
		}
	}
	if step == 0 {
		return seek, errors.New("null")
	}
	return seek + step, nil
}
