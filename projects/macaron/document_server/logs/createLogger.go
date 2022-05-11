package logs

import "fmt"

func CreateLogger(filename string) *Logger {
	l := NewLogger()

	cw := NewConsoleWriter()
	l.RegisterWriter(cw)

	fw := NewFileWriter()

	if err := fw.SetFile(filename); err != nil {
		fmt.Println(err)
	}

	l.RegisterWriter(fw)
	return l
}
