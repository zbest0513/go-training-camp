package main

import "notify-agent/internal"

func main() {
	processor := internal.NewAgentProcessor()
	processor.Start()
}
