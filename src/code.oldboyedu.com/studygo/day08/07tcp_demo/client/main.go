package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	tcp "code.oldboyedu.com/studygo/day08/07tcp_demo/server"
)

func postMsg(conn net.Conn) {
	//2.发送数据
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("请说话")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}
func main() {
	//1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed,err:")
		return
	}
	//2.发送消息
		postMsg(conn)

	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed,err:", err)
		return
	}
	defer listener.Close()
	for{
		conntmp, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go tcp.ProcessConn(conntmp)
	}


	conn.Close()
}
