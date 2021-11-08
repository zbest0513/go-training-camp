package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) start() {
	tcpaddr, _ := net.ResolveTCPAddr("tcp4", c.addr)
	tcpconn, _ := net.DialTCP("tcp", nil, tcpaddr)

	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please type in something:\n")
	n := 0
	for input.Scan() {
		str := input.Text()
		if str == "exit" {
			break
		}
		n++
		protocol := NewGoimProtocol()
		encode := protocol.Encode(Ver1, OptMsg, int32(1), str)
		//向tcpconn中写入数据
		tcpconn.Write(encode)
	}
}
