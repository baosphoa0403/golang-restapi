package main

import (
	"example/todogolang/controller"
	"example/todogolang/db"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	db := db.ConnectDb()
	controller.RootRouter(db);
}
