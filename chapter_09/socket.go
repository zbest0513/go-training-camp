package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync/atomic"
)

type server struct {
	addr        string //监听ip:port
	max         int32  //最大连接数
	connTimeOut int    //超时未有请求，关闭连接
}

func (receiver *server) start() {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", receiver.addr)
	checkErr(err)
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	checkErr(err2)
	//死循环的处理客户端请求
	go func() {
		leftover := receiver.max
		for {
			if atomic.AddInt32(&leftover, -1) > 0 {
				conn, err3 := tcplisten.Accept()
				//如果有错误直接跳过
				if err3 != nil {
					atomic.AddInt32(&leftover, 1)
					continue
				}
				go func() {
					defer func() {
						conn.Close()
						atomic.AddInt32(&leftover, 1)
					}()
					data := make([]byte, 256)
					for {
						n, errx := conn.Read(data)
						if n == 0 || errx != nil {
							return
						}
						cmd := strings.TrimSpace(string(data[0:n]))
						//向客户端发送数据，并关闭连接
						conn.Write([]byte(cmd + "\r\n"))
					}
				}()
			} else {
				atomic.AddInt32(&leftover, 1)
				log.Println("连接已经满了")
			}
		}
	}()
}

func checkErr(err error) {
	log.Fatalf(fmt.Sprintf("error:%v", err))
}
