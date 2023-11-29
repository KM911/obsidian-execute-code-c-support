package util

import "os"

func ReadAll(src string) string {
	//读取文件内容并返回全部
	//data = ""
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func CreatFile(src string, data string) {
	//创建文件并写入内容
	//data = ""
	err := os.WriteFile(src, []byte(data), 0666)
	if err != nil {
		panic(err)
	}
}

func DeleteFile(src string) {
	//删除文件
	err := os.Remove(src)
	if err != nil {
		panic(err)
	}
}

func CopyFile(src string, dst string) {
	//复制文件
	data := ReadAll(src)
	CreatFile(dst, data)
}
