package main

import "fmt"

//学生管理系统
//有一个物件
//1.它保存了一些数据 --->结构体的字段
//2.它有三个功能     --->结构体内的方法
type student struct {
	id   int64
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudents map[int64]student
	index       int64
}

func (s studentMgr) showStudents() {
	for _, stu := range s.allStudents {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
	// fmt.Println(s, s.index)

}
func (s studentMgr) addStudents() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号:")
	fmt.Scan(&id)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)
	newStu := student{
		id:   id,
		name: name,
	}
	s.index = 100
	s.allStudents[newStu.id] = newStu
	fmt.Println("添加成功")
}
func (s studentMgr) editStudents() {
	//用户输入学号
	//展示该学号对应的学生信息，如果没有提示查无此人
	var stuID int64
	fmt.Println("请输入学号:")
	fmt.Scan(&stuID)
	stuObj, ok := s.allStudents[stuID]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Printf("学号:%d 姓名:%s\n", stuObj.id, stuObj.name)
	}
	var newName string
	fmt.Println("请输入新的姓名:")
	fmt.Scan(&newName)
	var newStu = student{
		id:   stuObj.id,
		name: newName,
	}
	s.allStudents[stuObj.id] = newStu
	fmt.Println("修改成功")

}
func (s studentMgr) deleteStudents() {
	var stuID int64
	fmt.Println("请用户输入想要删除的学生学号")
	fmt.Scanln(&stuID)
	_, ok := s.allStudents[stuID]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Printf("学号:%d 姓名:%s\n", s.allStudents[stuID].id, s.allStudents[stuID].name)
		delete(s.allStudents, stuID)
		fmt.Println("删除成功!")
	}

}
