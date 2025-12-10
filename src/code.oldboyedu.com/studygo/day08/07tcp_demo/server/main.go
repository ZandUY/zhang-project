package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ProcessConn(conn net.Conn) {
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			return
		}
		fmt.Println("read from client:", string(tmp[:n]))
		fmt.Println("请回复")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

}
func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed,err:", err)
		return
	}

	//2.等待客户端连接

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed ,err :", err)
			return
		}
		//3.与客户端通信
		go ProcessConn(conn)
	}

}
