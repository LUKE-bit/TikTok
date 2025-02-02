package main

import (
	"TikTok/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	dao.Init()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
