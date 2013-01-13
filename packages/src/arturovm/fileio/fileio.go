package fileio

import (
	"os"
	//"bytes"
	"io"
)

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	size := file.Stat().Size()
	data := make([]byte, size)
	for {
		n, err := file.Read(data)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	return data
}
