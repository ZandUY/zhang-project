package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
func query(){

}
func insert(){
	
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goday10"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn  %s invalid ,err :%v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Open %s failed ,err :%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
}
