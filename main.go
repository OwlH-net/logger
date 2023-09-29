package logger

import "fmt"

// Logger module to create log trace on different formats.
type Logger struct {
}

// LogConsole sends to console the string in ss
func (*Logger) LogConsole(ss string, other ...interface{}) {
	fmt.Println(ss)
}
