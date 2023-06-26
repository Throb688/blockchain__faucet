package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test2/db"
)

//进行页面渲染
func LoginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}


func InitRouter()  {
	//gin框架的实列化
	router := gin.Default()

	//数据库的连接与关闭
	err := db.InitDB()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer db.CloseMysql()

	//加载HTML文件
	router.LoadHTMLGlob("./static/html/*")
	//加载静态文件
	router.StaticFS("/static", http.Dir("./static"))

	//打开页面
	router.GET("/", LoginPage)

	//提交申请
	router.POST("/home", HomePage)

	router.Run(":8080")
}
// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{}) {
	c.HTML(http.StatusOK, "success.html", gin.H{
		"data":data,
	})
	//c.JSON(200, gin.H{
	//	"code": 0,
	//	"data": data,
	//})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.HTML(http.StatusOK, "error.html", gin.H{
		"data":msg,
	})
	//c.JSON(200, gin.H{
	//	"code":    1,
	//	"message": msg,
	//})
}

func GetRequestIP(c *gin.Context)string{
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

