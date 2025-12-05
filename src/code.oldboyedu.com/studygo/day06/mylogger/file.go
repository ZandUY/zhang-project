package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	errFileName string
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open err Log file failed, err:%v", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)

}
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)

}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize

}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.fileObj) {
				newFile, err := f.splitFile(f.fileObj)
				if err != nil {
					return
				}
				f.fileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}

}
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000")

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, f.fileName, nowStr)
	//1.关闭当前日志文件
	f.fileObj.Close()
	//2.备份一下rename  xx.log -> xx.log.bak201908031709
	os.Rename(logName, newLogName)
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v", err)
		return nil, err
	}

	//4.重新赋值给f.fileObj
	return fileObj, nil
}
