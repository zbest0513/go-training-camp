package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"unsafe"
)

func main() {
	server1()

	time.Sleep(1)
	//
	go cli1()
	select {}

	println(fmt.Sprintf("%v", convertToBin(1)))
	println(fmt.Sprintf("%v", bytes2str(str2bytes("a"))))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func convertToBin(num int) []byte {
	data := int16(num)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func server1() {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:2880")
	print(err)
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	print(err2)
	//死循环的处理客户端请求
	go func() {
		for {
			//等待客户的连接
			//注意这里是无法并发处理多个请求的
			conn, err3 := tcplisten.Accept()
			//如果有错误直接跳过
			if err3 != nil {
			}
			go func() {
				defer conn.Close()
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
		}
	}()
}

func cli1() {
	//我们模拟请求网易的服务器
	//ResolveTCPAddr用于获取一个TCPAddr
	//net参数是"tcp4"、"tcp6"、"tcp"
	//addr表示域名或IP地址加端口号
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:2880")
	print(err)
	//DialTCP建立一个TCP连接
	//net参数是"tcp4"、"tcp6"、"tcp"
	//laddr表示本机地址，一般设为nil
	//raddr表示远程地址
	tcpconn, err2 := net.DialTCP("tcp", nil, tcpaddr)
	print(err2)
	//向tcpconn中写入数据
	_, err3 := tcpconn.Write([]byte("hello \r\n\r\n"))

	print(err3)
	////读取tcpconn中的所有数据
	for {
		data := make([]byte, 256)
		n, errx := tcpconn.Read(data)
		if n == 0 || errx != nil {
			return
		}
		cmd := strings.TrimSpace(string(data[0:n]))
		//向客户端发送数据，并关闭连接
		println(cmd)
	}
}

func print(err error) {
	if err == nil {
		return
	}
	log.Println(fmt.Sprintf("err:%v", err))
}
