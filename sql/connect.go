package sql

import (
	"database/sql"
	"time"
)

/*
@Time : 2024/8/14 16:10
@Author : echo
@File : connect
@Software: GoLand
@Description:
*/

func CreateDB(userName, password, url, port, dataBase string) (*sql.DB, error) {
	//dsn := "use:123456@tcp(127.0.0.1:3306)/test"
	dbInfo := userName + ":" + password + "@tcp(" + url + ":" + port + ")/" + dataBase
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	// 设置连接池参数
	db.SetMaxOpenConns(25)                 // 最大打开连接数
	db.SetMaxIdleConns(25)                 // 最大空闲连接数
	db.SetConnMaxLifetime(time.Minute * 5) // 连接最大生命周期
	return db, nil
}
