package main

import (
	"fmt"
	"compress/gzip"
	"io"
	"os"
	"bytes"
)

// constant definitions

const END byte = 0
const BYTE byte = 1
const SHORT byte = 2
const COMPOUND byte = 10

// type defintions

type TAG_End struct { // 0
	tagType int8
}

type TAG_Byte struct { // 1
	tagType int8
	name *TAG_String
	theByte int8
}

type TAG_Short struct { // 2
	tagType int8
	name *TAG_String
	theShort int16
}

type TAG_Int struct { // 3
	tagType int8
	name *TAG_String
	theInt int32
}

type TAG_String struct { // 8
	length *TAG_Short
	text []int8
}

func toShort(a, b byte) int16 {
	aShort := (a << 8) | b
	return int16(aShort)
}

func main() {
	file, fiErr := os.Open("test.nbt")
	if fiErr != nil {
		panic(fiErr)
	}
	defer file.Close()
	data := make([]byte, 1024)
	gReader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	for {
		n, readErr := gReader.Read(data)
		if readErr != nil && readErr != io.EOF {
			panic(readErr)
		}
		if n == 0 {
			break
		}
	}
	byteReader := bytes.NewReader(data)
	for {
		tagType, _ := int(byteReader.ReadByte())
		var name string
		if typeB != END {
			high, _ := byteReader.ReadByte()
			low, _ := byteReader.ReadByte()
			nameLength := toShort(high, low)
			currentIndex := len(data) - int(byteReader.Len())
			name = string(data[currentIndex:int16(currentIndex) + nameLength])
			byteReader.Seek(int64(nameLength), 1)
			fmt.Println(name)
		}
		if byteReader.Len() == 0 {
			break
		}
	}
}