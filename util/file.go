package util

import "os"

func ReadAll(src string) string {
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func CreatFile(src string, data string) {
	err := os.WriteFile(src, []byte(data), 0666)
	if err != nil {
		panic(err)
	}
}

func DeleteFile(src string) {
	err := os.Remove(src)
	if err != nil {
		panic(err)
	}
}

func CopyFile(src string, dst string) {
	data := ReadAll(src)
	CreatFile(dst, data)
}
