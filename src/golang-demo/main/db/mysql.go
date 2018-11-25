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
	rows, err := db.Query("select * from userinfo",3434,34)
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

func  InsertTest()  {
	db, err := sql.Open("mysql", "root:Wp123456!@tcp(47.104.15.216:3306)/credit_card?charset=utf8")
	fmt.Println(err)
	fmt.Println("insert test")
	//res, err :=db.Exec("insert into user_info (id,name ,age) values(?,?,?)",1,"finley",24)
	stmt,err:=db.Prepare("insert into user_info (id,name ,age) values(?,?,?)")
	stmt.Exec(1,"finleyy",24)
	fmt.Println(stmt,err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

