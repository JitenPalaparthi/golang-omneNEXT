package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello OmneNEXT")
	//io.Writer

	fw := New("data.txt")
	fmt.Fprintln(fw, "Hello World")
	fmt.Fprintln(fw, "Hello World, trying to write to the file")
}

type FileWriter struct {
	FileName string
}

func New(filename string) *FileWriter {
	return &FileWriter{FileName: filename}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	// n, err = f.Write(p)
	// return n, err
	defer f.Close()
	return f.Write(p)
}
