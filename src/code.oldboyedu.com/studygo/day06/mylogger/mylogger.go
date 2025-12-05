package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint8

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}
type ConsoleLogger struct {
	Level LogLevel
}

func parseLoglevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}
func getLogString(lv LogLevel)string{
	switch lv{
		case DEBUG:
			return "DEBUG"
		case TRACE:
			return "TRACE"
		case INFO:
			return "INFO"
		case WARNING:
			return "WARNING"
		case ERROR:
			return "ERROR"
		case FATAL:
			return "FATAL"
		default:
			return "DEBUG"
	}
}
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}
func getInfo(skip int)(funcName,fileName string,lineNo int){
		pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("获取失败")
		return
	}
	funcName=runtime.FuncForPC(pc).Name()
	fileName=path.Base(file)
	return funcName,fileName,lineNo
}
func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}