package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
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
func queryOne(id int) {
	var u1 user
	sqlStr := `select id,name,age from user where id=?;`
	rowObj := db.QueryRow(sqlStr, id)
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("id:%d  name:%s  age:%d", u1.id, u1.name, u1.age)
}
func queryMore(n int) {
	sqlStr := `select id,name,age from user where id>?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err :%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan rows failed, err:%v\n", err)
			return
		}
		fmt.Printf("%#v\n", u)

	}

}
func update(newAge int, id int) {
	sqlStr := `update user set age=? where id=?`
	res, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("%s failed,err:%v", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Printf("更新影响了%d条记录", n)
}
func insert() {
	sqlStr := `insert into user(name,age) values("聂灿",19)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("%s failed, err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	num, err1 := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id:%d failed,err:%v\n", err)
		return
	}
	if err1 != nil {
		fmt.Printf("get id:%d failed,err:%v\n", err1)
		return
	}
	fmt.Printf("共插入%d行,插入的最后一行:%d\n", num, id)
}
func transcation() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v", err)
		return
	}
	//执行多个sql操作
	sqlStr1 := `update user set age=age-2 where id=1`
	sqlStr2 := `update xxx set age=age+2 where id=5`
	_, err1 := tx.Exec(sqlStr1)
	if err1 != nil {
		tx.Rollback()
		fmt.Printf("执行SQL1出错了")
		return
	}
	_, err2 := tx.Exec(sqlStr2)
	if err2 != nil {
		tx.Rollback()
		fmt.Printf("执行SQL2出错了")
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("执行Commit出错了")
		return
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Open db failed,err:%v", err)
		return
	}
	fmt.Println("连接数据库成功")

	transcation()
}
