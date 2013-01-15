package fileio

import (
	"os"
	"io"
)

// ReadFile opens a given with the given filename and returns its
// contents
func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stat, _ := file.Stat()
	size := stat.Size()
	data := make([]byte, size)
	for {
		n, err := file.Read(data)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
	}
	return data, nil
}
