package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/goland"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}

	// sql语句，如果没存在库表todo，则新建一个
	var sqlStr = `CREATE TABLE IF NOT EXISTS todo(
		id  int NOT NULL PRIMARY KEY AUTO_INCREMENT,
		ToAddress VARCHAR(100) NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	// 执行建表语句
	_, err = db.Exec(sqlStr)
	if err != nil {
		return err
	}
	var sqlStr1 = `CREATE TABLE IF NOT EXISTS  acount  (
		time date NOT NULL,
		count int(11) NOT NULL
	)`
	_, err = db.Exec(sqlStr1)
	if err != nil {
		return err
	}

	return nil
}

//向表todo中插入数据
func InsertTodo(address string) (err error) {
	var sqlStr = `INSERT INTO todo (ToAddress) VALUES (?)`

	_, err = db.Exec(sqlStr, address)
	if err != nil {
		return err
	}

	return
}

//关闭数据库
func CloseMysql() {
	_ = db.Close()
}

//向表acount插入数据
func InsertTodo1() (err error) {
	var count int
	var sqlStr = `SELECT COUNT(*) as b from acount where time = CURDATE()`
	var sqlStr1 = `INSERT INTO acount (time,count) VALUES (CURDATE(),1)`
	var sqlStr2 = `UPDATE acount SET COUNT = COUNT + 1 WHERE TIME =CURDATE()`
	err = db.QueryRow(sqlStr).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0{
		_, err = db.Exec(sqlStr1)
		if err != nil {
			return err
		}
	}else {
		_, err = db.Exec(sqlStr2)
		if err != nil {
			return err
		}
	}

	return
}
