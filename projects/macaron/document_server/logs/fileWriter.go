package logs

import (
	"errors"
	"fmt"
	"os"
)

// 文件写入器
type fileWriter struct {
	file *os.File
}

// 设置文件写入器写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(filename)
	return err
}

// 实现LogWriter的Write()方法
func (f *fileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}
	str := fmt.Sprintf("%v\n", data)
	_, err := f.file.Write([]byte(str))
	return err
}

// 创建文件写入器实例
func NewFileWriter() *fileWriter {
	return &fileWriter{}
}
