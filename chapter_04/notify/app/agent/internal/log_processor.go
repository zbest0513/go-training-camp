package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

type LogProcessor struct {
	path   string
	offset int64
	count  int
	isRun  bool
	addr   string
	key    string
}

func CreateLogProcessor(path string, max int, offset int64, addr string, key string) *LogProcessor {
	lp := &LogProcessor{
		path:   path,
		count:  max,
		offset: offset,
		addr:   addr,
		key:    key,
	}
	go lp.Start()
	return lp
}

func (lp *LogProcessor) Start() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprintf("日志监听异常:%v", err))
		}
	}()
	lp.isRun = true
	for lp.isRun {
		log.Println(fmt.Sprintf("%v准备采集数据%v", lp.path, lp.offset))
		newIdx, err, errStr := lp.readLines(lp.offset)
		lp.offset = newIdx
		if errStr != "" {
			log.Println(fmt.Sprintf("%v采集到日志:%v", lp.key, errStr))
			dto := SendMsgDto{
				Content: lp.key + ":" + errStr,
			}
			result := Result{}
			_, err2 := Post(lp.addr, dto, &result)
			if err2 != nil {
				log.Println(fmt.Sprintf("%v上报错误:%v", lp.key, err2))
			}
		}
		if errors.Is(err, eofError) {
			time.Sleep(time.Second * time.Duration(1))
			continue
		} else if err != nil {
			log.Println(fmt.Sprintf("读取文件错误:%v", err))
			lp.Destroy()
		}
	}
}

func (lp *LogProcessor) Destroy() {
	lp.isRun = false
}

func (lp *LogProcessor) readLines(idx int64) (int64, error, string) {
	f, err := os.OpenFile(lp.path, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if idx == 0 {
		stat, err := f.Stat()
		if err != nil {
			panic(err)
		}
		if size := stat.Size(); size > 0 {
			idx = size
		}
	}

	seek, err := f.Seek(idx, io.SeekStart)
	rd := bufio.NewReader(f)
	var step int64
	i := 0
	var flag = false
	var exceptionStr string
	for lp.count > i || flag {
		i++
		line, err := rd.ReadString('\n')
		step += int64(len(line))
		if err != nil || io.EOF == err {
			break
		}
		exception, _ := regexp.Compile(`\w+[.]\w+Exception:`)

		match := exception.MatchString(line)
		if match {
			flag = true
			exceptionStr = fmt.Sprintf("%v", line)
			continue
		}
		at, _ := regexp.Compile(`[\t]at `)
		cause, _ := regexp.Compile(`Caused by`)

		match = at.MatchString(line)
		match2 := cause.MatchString(line)
		if !match && !match2 {
			flag = false
		} else {
			exceptionStr = fmt.Sprintf("%v%v", exceptionStr, line)
		}
	}
	if step == 0 {
		return seek, eofError, exceptionStr
	}
	return seek + step, nil, exceptionStr
}
