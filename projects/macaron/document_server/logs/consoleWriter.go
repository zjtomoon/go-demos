package logs

import (
	"fmt"
	"os"
)

// 命令行写入器
type consoleWriter struct {
}

// 实现LogWriter的Write()方法
func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// 创建命令行写入器实例
func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
