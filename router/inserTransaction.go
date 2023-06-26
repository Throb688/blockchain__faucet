package router

import (
	"github.com/gin-gonic/gin"
	"test2/controller"
	"test2/db"
)

//进行转账
func HomePage(c *gin.Context) {
	//获取差数
	nick := c.PostForm("url")
	//执行合约函数
	_, err :=controller.SetQuantity()
	if err!=nil{
		respError(c, err)
		return
	}
	//进行判断
	_, err =controller.Withdraw(nick,GetRequestIP(c))
	if err!=nil{
		respError(c,err)
		return
	}

	//向表todo记录该交易
	if err := db.InsertTodo(nick); err != nil {
		respError(c, err)
		return
	}
	//向表进行统计
	if err := db.InsertTodo1(); err != nil {
		respError(c, err)
		return
	}
	//返回结果
	respOK(c,"转账成功")

}
