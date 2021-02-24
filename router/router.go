package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"wj-go/dao"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	apiVersionOne := router.Group("/api/")
	apiVersionOne.GET("books", Books)
	return router
}

func Books(c *gin.Context) {
	(&dao.Book{}).Find(c)
}

func Demos(c *gin.Context) {
	//使用go-sql-driver查询
	db, err := sql.Open("mysql", "root:root@tcp(192.168.43.1:3306)/white_jotter?charset=utf8")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT id,title,author,date,press,abs,cid FROM book")
	checkErr(err)

	var id int
	//var cover string
	var title string
	var author string
	var date string
	var press string
	var abs string
	var cid int

	for rows.Next() {
		err = rows.Scan(&id, &title, &author, &date, &press, &abs, &cid)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(title)
		fmt.Println(author)
		fmt.Println(press)
	}
	c.JSON(200, gin.H{
		"id":     id,
		"title":  title,
		"author": author,
		"date":   date,
		"press":  press,
		"abs":    abs,
		"cid":    cid,
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
