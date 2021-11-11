package main

import "chapter09/net"

//作业1在readme.md中
//作业2实现
func main() {
	var addr = "127.0.0.1:2880"
	newServer := net.NewServer(addr, 10)
	newServer.Start()
	client := net.NewClient(addr)
	client.Start()
}
