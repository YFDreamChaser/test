package logger

import (
	"fmt"
	"os"
)

type consoleWriter struct {
}

//实现LogWriter的Write方法
func (f *consoleWriter) Write(data interface{}) error {
	//将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	//将数据以字节数组写入命令行中
	_, err := os.Stdout.Write([]byte(str))
	return err
}

//创建实例
func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
