package main

import (
	"fmt"
	"os"
)


func showMenu() {
	fmt.Println("Welcome sms!")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

var smr studentMgr

func main() {
	smr=studentMgr{
		allStudents:make(map[int64]student,100),
		index:99,
	}
	for {
		showMenu()
		fmt.Print("请输入要执行的操作:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("您输入的是:", choice)
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:	
			smr.deleteStudents()
		case  5:
			os.Exit(1)
		default:
			fmt.Println("滚~")
		}
	}
}
