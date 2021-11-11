package net

import (
	"log"
	"net"
	"sync/atomic"
)

type Server struct {
	addr string //监听ip:port
	max  int32  //最大连接数
}

func NewServer(addr string, max int32) *Server {
	return &Server{
		addr: addr,
		max:  max,
	}
}

func (receiver *Server) Start() {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", receiver.addr)
	print(err)
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	print(err2)
	go func() {
		leftover := receiver.max
		for {
			if atomic.AddInt32(&leftover, -1) > 0 {
				conn, err3 := tcplisten.Accept()
				if err3 != nil {
					atomic.AddInt32(&leftover, 1)
					continue
				}
				go func() {
					defer func() {
						conn.Close()
						atomic.AddInt32(&leftover, 1)
					}()
					length := make([]byte, 4)
					for {
						//读取package length
						n1, errx := conn.Read(length)
						if n1 == 0 || errx != nil {
							return
						}
						strBytes := length
						protocol := NewGoimProtocol()
						//解析package length
						count := protocol.DecodePackageLength(strBytes)
						//读取除了package length 外，整包的内容（head+body）
						data := make([]byte, count)
						n2, errx := conn.Read(data)
						if n2 == 0 || errx != nil {
							return
						}
						strBytes = data
						//解析包内容
						dStr := protocol.Decode(strBytes)
						s := dStr.(string)
						println("服务端接受:", s)
					}
				}()
			} else {
				atomic.AddInt32(&leftover, 1)
				log.Println("连接已经满了")
			}
		}
	}()
}
