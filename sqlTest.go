package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	sql2 "hello/sql"
	"log"
)

/*
@Time : 2024/8/14 15:39
@Author : echo
@File : sqlTest
@Software: GoLand
@Description:
*/
type UserMul struct {
	ID        int
	UserName  string
	LoginName string
	Password  string
}

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql2.CreateDB("root", "123456", "127.0.0.1", "3306", "test")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功连接到数据库！")
	users, err := queryUer(db)
	if err != nil {
		return
	}
	for _, user := range users {
		fmt.Printf("ID: %d, UserName: %s, LoginName: %s, Password: %s\n", user.ID, user.UserName, user.LoginName, user.Password)
	}
	fmt.Println("查询成功！")
	err = insertUser(db, "echo", "echoLogin", "123456")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("插入成功！")
	err = updateUser(db, 2, "echoUpdate")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("更新成功！")
}

func queryUer(db *sql.DB) ([]UserMul, error) {
	query := "select * from user"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []UserMul
	for rows.Next() {
		var user UserMul
		if err := rows.Scan(&user.ID, &user.UserName, &user.LoginName, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}
func insertUser(db *sql.DB, userName, loginName, password string) error {
	query := "insert into user (user_name,login_name,password) values (?,?,?)"
	_, err := db.Exec(query, userName, loginName, password)
	if err != nil {
		return err

	}
	fmt.Printf("用户 %s 插入成功！\n", userName)
	return nil
}

func updateUser(db *sql.DB, id int, userName string) error {
	query := "update user set user_name = ? where  id = ?"
	_, err := db.Exec(query, userName, id)
	if err != nil {
		return err
	}
	fmt.Printf("用户 %s 更新成功！\n", userName)
	return nil
}
