package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	Id     int            `json:"id" db:"id"`
	Cover  sql.NullString `json:"cover" db:"cover"`
	Title  string         `json:"title" db:"title"`
	Author string         `json:"author" db:"author"`
	Date   string         `json:"date" db:"date"`
	Press  string         `json:"press" db:"press"`
	Abs    string         `json:"abs" db:"abs"`
	Cid    int            `json:"cid" db:"cid"`
}

func (f *Book) Find(c *gin.Context) {
	//使用sqlx进行查询
	db, err := sqlx.Open("mysql", "root:root@tcp(192.168.43.1:3306)/white_jotter?charset=utf8")
	checkErr(err)

	book := []Book{}
	db.Select(&book, "select * from book")
	fmt.Printf("%#v\n%#v", book[0], book[1])
	r, err := json.Marshal(book)
	c.JSON(200, string(r))
	//c.JSON(200, gin.H{
	//	"id":     "1",
	//	"title":  "中文",
	//	"author": "测试",
	//	"date":   "date",
	//	"press":  "press",
	//	"abs":    "abs",
	//	"cid":    "cid",
	//})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
