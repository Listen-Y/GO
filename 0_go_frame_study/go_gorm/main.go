package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	gorm.Model
}

const (
	userName = "root"           //登录Mysql的用户名
	password = "listen"         //对应用户名的密码
	ip       = "127.0.0.1"      //ip地址
	port     = "3306"           //端口
	dbName   = "go_mysql_study" //数据库名字
)

var db *gorm.DB

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8", "&parseTime=True"}, "")

	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       path,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = database
	fmt.Println("connect success")
}

func main() {
	//gormInsert()
	selectTest()
}

func selectTest() {

	users := make([]User, 0)
	// 获取全部记录
	result := db.Find(&users)
	// SELECT * FROM users

	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	fmt.Println(result.Error)        // returns error
}

func gormInsert() {

	email := "a.com"
	birthday := time.Now()

	var user = User{
		Name:         "listen",
		Email:        &email,
		Age:          20,
		Birthday:     &birthday,
		MemberNumber: sql.NullString{},
		ActivedAt:    sql.NullTime{},
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	result := db.Create(&user) // 通过数据的指针来创建

	fmt.Println(user.ID)             // 返回插入数据的主键
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
}
