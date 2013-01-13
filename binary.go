package main

import (
	"os"
	"compress/gzip"
	//"fmt"
)

type TestStruct struct {
	test string
	test2 string
	ha int8
}

func main() {
	fo, err := os.Create("output.dat")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(fo)
	defer fo.Close()
	message := "Hello!"
	mBytes := []byte(message)
	_, wErr := writer.Write(mBytes)
	if wErr != nil {
		panic(wErr)
	}
	writer.Close()
	fo.Close()
}