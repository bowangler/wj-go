package main

import (
	"fmt"
	"wj-go/router"
)

func main() {
	fmt.Println("This works.")
	r := router.InitRouter()
	r.Run()
}
