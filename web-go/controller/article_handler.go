package controller

import (
	"net/http"
	"strconv"
	"web-go/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) GetArticlePage(c *gin.Context) {
	errMsg := ""
	articles, err := model.GetAllArticles(s.DB)
	if err != nil {
		errMsg = "エラー発生"
		articles = []model.Article{}
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"articles": articles,
		"errMsg":   &errMsg,
	})
}

func (s *Server) NewArticlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "new.tmpl", gin.H{})
}

func (s *Server) PostArticleHandler(c *gin.Context) {
	c.Request.ParseForm()
	article := new(model.Article)
	article.Text = c.Request.Form["text"][0]
	if err := model.PostArticle(s.DB, article); err != nil {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "登録できませんでした",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) DeleteArticleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	if err := model.DeleteArticleByID(s.DB, uint(id)); err != nil {
		articles, _ := model.GetAllArticles(s.DB)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"articles": articles,
			"errMsg":   "エラー発生",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
