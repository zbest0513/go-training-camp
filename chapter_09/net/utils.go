package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"unsafe"
)

func Int32ToBytes(num interface{}) []byte {
	data, ok := num.(int32)
	if !ok {
		opt, ok := num.(Opt)
		if ok {
			data = int32(opt)
		}
	}
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt32(data []byte) int32 {
	var num int32
	bytebuf := bytes.NewBuffer(data)
	binary.Read(bytebuf, binary.BigEndian, &num)
	return num
}

func Int16ToBytes(num interface{}) []byte {
	data, ok := num.(int16)
	if !ok {
		ver, ok := num.(Ver)
		if ok {
			data = int16(ver)
		}
	}
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt16(data []byte) int16 {
	var num int16
	bytebuf := bytes.NewBuffer(data)
	binary.Read(bytebuf, binary.BigEndian, &num)
	return num
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type sc struct {
	start  int
	length int
}

// BytesCombine 处理数组合并
func BytesCombine(pBytes ...[]byte) []byte {
	var total int
	m := make(map[int]sc)
	for i, pByte := range pBytes {
		tmp := total
		lh := len(pByte)
		total += lh
		m[i] = sc{
			start:  tmp,
			length: lh,
		}
	}
	result := make([]byte, total, total)
	for n, t := range m {
		for i := 0; i < t.length; i++ {
			result[i+t.start] = pBytes[n][i]
		}
	}
	return result
}

func print(err error) {
	if err == nil {
		return
	}
	log.Println(fmt.Sprintf("err:%v", err))
}
