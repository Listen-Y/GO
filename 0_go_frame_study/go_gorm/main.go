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
	//selectTest()
	//updateTest()
	//deleteTest()
	softDelete()
}

/**
如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！

拥有软删除能力的模型调用 Delete 时，记录不会被从数据库中真正删除。但 GORM 会将 DeletedAt 置为当前时间，
并且你不能再通过正常的查询方法找到该记录。

永久删除
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
*/
func softDelete() {
	user := User{
		Name:         "",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
		ActivedAt:    sql.NullTime{},
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	// user 的 ID 是 `111`
	db.Delete(&user)
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("age = ?", 20).Delete(&User{})
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	// 在查询时会忽略被软删除的记录
	users := make([]User, 0)
	// 获取全部记录
	db.Find(&users)
	// SELECT * FROM users
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

}

func deleteTest() {
	user := User{
		Name:         "",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
		ActivedAt:    sql.NullTime{},
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	// Email 的 ID 是 `10`
	db.Delete(&user)
	// DELETE from emails where id = 10;

	// 带额外条件的删除
	db.Where("name = ?", "jinzhu").Delete(&user)
	// DELETE from emails where id = 10 AND name = "jinzhu";
}

func updateTest() {
	// 条件更新
	db.Model(&User{}).Where("age = ?", 20).Update("name", "1111")
	// UPDATE users SET name='1111', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User 的 ID 是 `111`
	user := User{
		Name:         "",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
		ActivedAt:    sql.NullTime{},
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	db.Model(&user).Update("name", "2222")
	// UPDATE users SET name='2222', updated_at='2013-11-17 21:34:10' WHERE id=111;

	user.ID++
	// 根据条件和 model 的值进行更新
	db.Model(&user).Where("age = ?", 20).Update("name", "3333")
	// UPDATE users SET name='3333', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
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
