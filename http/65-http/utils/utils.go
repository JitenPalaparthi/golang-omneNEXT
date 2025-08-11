package utils

import "os"

func SaveToFile(fileName string, data []byte) (int, error) {

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	return f.Write(data)

}
