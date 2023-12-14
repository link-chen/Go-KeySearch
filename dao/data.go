package dao

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
)
import (
	_ "github.com/go-sql-driver/mysql"
)

var user struct {
	id       int32
	name     string
	password string
}
var count struct {
	Num int32
}
var DB *sql.DB
var err error

func InitDataBase() {
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	if DB == nil {
		fmt.Println("Mysql Init failed")
	} else {

	}
}

// 注册
func Regist(password string, name string) int32 {
	hashtable := make(map[int32]int32)
	rows, qerr := DB.Query("select * from user")
	if qerr != nil {
		log.Println(err)
		return -1
	}
	defer rows.Close()
	for rows.Next() {
		qerr := rows.Scan(&user.id, &user.name, &user.password)
		if qerr != nil {
			log.Println(err)
			return -1
		}
		hashtable[user.id]++
	}
	var id = rand.Int31n(1000000)
	flag := false
	for !flag {
		if hashtable[id] != 0 {
			id = rand.Int31n(100000)
		} else {
			DB.Exec("insert into user values(?,?,?)", id, name, password)

			flag = true
		}
	}

	return id
}

// 登录检测
func CheckExist(id int32, Password string) bool {
	rows, qerr := DB.Query("select * from user where id= ?", id)
	if qerr != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()
	var index int
	for rows.Next() {
		qerr := rows.Scan(&user.id, &user.name, &user.password)
		if qerr != nil {
			log.Println(err)
			return false
		}
		index++
	}
	return index != 0
}

// 注销
func RemoveFromList(id int32, name string, password string) bool {
	rows, qerr := DB.Query("select * from user where id= ?", id)
	if qerr != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		qerr := rows.Scan(&user.id, &user.name, &user.password)
		if qerr != nil {
			log.Println(qerr)
			return false
		}
		if user.id == id && user.name == name && user.password == password {
			DB.Exec("delete from user where id=?", id)
			return true
		} else {
			return false
		}
	}
	return false
}

// 改密
func ChangePassword(id int32, oldpassword string, newpassword string) bool {
	rows, qerr := DB.Query("select * from user where id= ?", id)
	if qerr != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		qerr := rows.Scan(&user.id, &user.name, &user.password)
		if qerr != nil {
			log.Println(qerr)
			return false
		}
		if user.password != oldpassword {
			return false
		} else {
			DB.Exec("update user set password=? where id=?", newpassword, id)
			return true
		}
	}
	return false
}

func LeaveMessage(id int, message string) bool {
	DB.Exec("insert into wordboard values(?,?)", id, message)
	return true
}
func RemoveMessage(id int, message string) bool {
	DB.Exec("delete from wordboard where id=? and message=?", id, message)
	return true
}

func DownLoadCount() bool {
	fmt.Println("count")
	rows, _ := DB.Query("select * from Count")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count.Num)
	}
	fmt.Println("count1")
	DB.Exec("update Count set Num= ?", count.Num+1)
	fmt.Println("addcount")
	return true
}
