package main

import (
	"fmt"
	"notify-agent/internal"
)

func main() {

	var idx int64 = 0
	var flag = true
	for flag {
		seek, err := internal.ReadLines("/Users/wbh/Desktop/test.txt", 1, idx)
		idx = seek
		if err != nil {
			println(fmt.Sprintf("空了%v,%v", err, seek))
			flag = false
		}
	}
}
