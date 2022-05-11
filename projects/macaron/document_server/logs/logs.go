package logs

// 参考： https://github.com/hyahm/golog
// 参考： https://blog.csdn.net/jiohfgj/article/details/120606777
// logger https://github.com/wonderivan/logger.git

// 日志写入器
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// 将一个Data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

// 创建日志器实例
func NewLogger() *Logger {
	return &Logger{}
}
