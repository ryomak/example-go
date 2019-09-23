package main

import (
	"web-go/controller"
	"web-go/util"

	"github.com/gin-gonic/gin"
)

func main() {

	server := controller.Server{
		DB: util.InitDB(),
	}
	router := gin.Default()

	router.LoadHTMLGlob("view/*")
	//template(view) ページ
	router.GET("/", server.GetArticlePage)
	router.GET("/new", server.NewArticlePage)
	//form 処理
	router.POST("/new", server.PostArticleHandler)
	router.POST("/delete/:id", server.DeleteArticleHandler)

	router.Run(":8080")
}
