package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
	// filename := "./my.log"
	config := tail.Config{
		Follow:    true,  //进行跟随
		ReOpen:    true,  //重新打开
		MustExist: false, //文件打开失败不报错
		Poll:      true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	return
}
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
	// for {
	// 	line, ok := <-tailObj.Lines
	// 	if !ok {
	// 		fmt.Println("tail file close reopen, filename: ", tailObj.Filename)
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// 	fmt.Println("line:", line.Text)
	// }
}
