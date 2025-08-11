package utils

import (
	"log"
	"os"
)

var ChUser chan []byte
var ChErr chan *FileError

type FileError struct {
	Message string
	Code    string
	Data    []byte
}

func NewFileError(code string, message string, data []byte) *FileError {
	return &FileError{message, code, data}
}

func init() {
	ChUser = make(chan []byte, 10)
	ChErr = make(chan *FileError, 10)
	go SaveUsers("users.dat")
	go ProcessErrors()
}

func SaveUsers(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		ChErr <- NewFileError("1001", "file creation error", []byte(err.Error()))
		log.Println(err.Error())
		return
	}
	defer f.Close()
	for user := range ChUser {
		_, err := f.Write(user)
		if err != nil {
			ChErr <- NewFileError("1002", "user creation error", user)
			log.Println(err.Error() + string(user))
		}
	}
}

func ProcessErrors() {
	for fileError := range ChErr {
		log.Println(fileError.Code)
		// do the actual stuff
	}
}
func SaveToFile(fileName string, data []byte) (int, error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}
