package util

import (
	"os"
	"path/filepath"
)

// 为了提升兼容性 我们使用的路径都是就是linux风格的
// 我们的文件换行都应该使用就是 LF
// 函数命名采用 驼峰命名法 首字母大写
// 变量命名采用 蛇形命名法 首字母小写
// 为了避免和系统包冲突 我们的变量名结尾都加上下划线 例如 file_ path_ type_ 等等 尽量不影响阅读

// AbsPath
// 返回文件的绝对路径
// 移除错误处理
// e.g.
// a.txt -> /home/xxx/a.txt
// Makefile -> D:/SOFT/CODE/Makefile
func AbsPath(file_ string) string {
	abs, _ := filepath.Abs(file_)
	return filepath.ToSlash(abs)
}

/*
获取当前命令窗口的路径
移除错误处理
e.g.
CmdPath() -> /home/xxx
CmdPath() -> D:/SOFT/CODE
*/
func CmdPath() string {
	dir, _ := os.Getwd()
	return filepath.ToSlash(dir)
}

/*
返回可执行文件的路径 不包含可执行文件 不会因为在哪里运行而改变
起到了类似于 ./ 的作用
*/
func ExecutePath() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(path)
}
