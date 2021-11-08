package main

func main() {
	var addr = "127.0.0.1:2880"
	newServer := NewServer(addr, 10)
	newServer.start()
	client := NewClient(addr)
	client.start()
}
