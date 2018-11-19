package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
  数据库应用测试
  1、数据库连接
  2、增删改查
  3、ORM映射
 */


func QueryTest() {
	db, err := sql.Open("mysql", "root:529186@/test?charset=utf8")
	CheckErr(err)
	rows, err := db.Query("select * from userinfo")
	CheckErr(err)
	for rows.Next() {
		var id int
		var gender string
		var age int
		var nick string
		err = rows.Scan(&id, &gender, &age,&nick)
		CheckErr(err)
		fmt.Println(id)
		fmt.Println(nick)
		fmt.Println(gender)
		fmt.Println(age)
	}
	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

