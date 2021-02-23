package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Books() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM book")
	checkErr(err)

	for rows.Next() {
		var id int
		var cover string
		var title string
		var author string
		var date string
		var press string
		var abs string
		var cid int
		err = rows.Scan(&id, &cover, &title, &author, &date, &press, &abs, &cid)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(title)
		fmt.Println(author)
		fmt.Println(press)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
