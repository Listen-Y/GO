package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	userName = "root"           //登录Mysql的用户名
	password = "listen"         //对应用户名的密码
	ip       = "127.0.0.1"      //ip地址
	port     = "3306"           //端口
	dbName   = "go_mysql_study" //数据库名字
)

var Db *sqlx.DB

func init() {
	InitDB()
}

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	database, err := sqlx.Open("mysql", path)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	//验证连接
	if err := database.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	//设置数据库最大连接数
	database.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	database.SetMaxIdleConns(10)
	Db = database

	fmt.Println("connect success")
}

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func main() {
	//insertTest()

	//selectTest()

	//updateTest()

	//deleteTest()

	//transactionTest()
}

func insertTest() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert success:", id)
	defer Db.Close()
}

func selectTest() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select success:", person)
	defer Db.Close()

}

func updateTest() {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
	defer Db.Close()

}

func deleteTest() {
	res, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete succ: ", row)
	defer Db.Close()

}

/*
1） import (“github.com/jmoiron/sqlx")
2)  Db.Begin()        开始事务
3)  Db.Commit()        提交事务
4)  Db.Rollback()     回滚事务
*/

func transactionTest() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
	defer Db.Close()

}
