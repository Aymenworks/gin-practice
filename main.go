package main

import (
  "gin-practice/models"
  "net/http"
  "github.com/gin-gonic/gin"
)

var router *gin.Engine

var articleList = []models.Article{
  models.Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
  models.Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func indexPage(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html",
    gin.H{
    "title": "Pepito",
    },
  )
}
func initializeRoute() {
  router.GET("/", indexPage)
}

func main() {
  router = gin.Default()
  router.LoadHTMLGlob("templates/*")

  initializeRoute()

  router.Run()
}
