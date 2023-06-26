package main

import (
	"fmt"
	"log"
	"test2/controller"
	"test2/db"
	"test2/router"
)



func main() {
	//初始化数据库
	err := db.InitDB()
	if err!=nil{
		log.Fatal(err)
		fmt.Println(err)
		return
	}

	//路由进行初始化
	router.InitRouter()

	//执行每天重置函数
	controller.BoottimeTimingSettlement()

	//关闭数据库
	defer db.CloseMysql()

	//controller.Init()

}