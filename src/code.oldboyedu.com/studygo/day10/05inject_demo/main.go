package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u []user
	err := db.Select(&u, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	fmt.Printf("user:%#v\n", u)
}

type user struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn  %s invalid ,err :%v\n", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Open %s failed ,err :%v\n", err)
		return err
	}
	db.SetMaxOpenConns(10)
	return nil
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v", err)
		return
	}
	
	sqlInjectDemo("xxx' or 1=1 #")
	//select id, name, age from user where name='xxx' or 1=1 #'
}
