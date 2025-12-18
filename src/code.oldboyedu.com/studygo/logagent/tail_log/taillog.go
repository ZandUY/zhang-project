package taillog

import (
	"context"
	"fmt"

	"code.oldboyedu.com/studygo/logagent/kafka"
	"github.com/hpcloud/tail"
)

// 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	//为了能实现能够退出t.run()
	ctx context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx,cancel:=context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:  path,
		topic: topic,
		ctx:ctx,
		cancelFunc: cancel,
	}
	tailObj.init()
	return
}
func (t *TailTask) init() {
	// logEntry.Path  :
	config := tail.Config{
		Follow:    true,  //进行跟随
		ReOpen:    true,  //重新打开
		MustExist: false, //文件打开失败不报错
		Poll:      true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tailFile failed,err:", err)
		return
	}
	go t.run()
}
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n",t.path,t.topic)
			return
		case line := <-t.instance.Lines:
			// 	kafka.SendToKafka(t.topic, line.Text)
			//先把日志数据发到一个通道中
			kafka.SendToChan(t.topic, line.Text)
		}

	}

}

//	func Init(fileName string) (err error) {
//		// filename := "./my.log"
//		config := tail.Config{
//			Follow:    true,  //进行跟随
//			ReOpen:    true,  //重新打开
//			MustExist: false, //文件打开失败不报错
//			Poll:      true,
//			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
//		}
//		tailObj, err = tail.TailFile(fileName, config)
//		if err != nil {
//			fmt.Println("tail file failed,err:", err)
//			return
//		}
//		return
//	}
func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
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
