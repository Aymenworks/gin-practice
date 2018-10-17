package main

import (
  "gin-practice/models"
  "net/http"
  "github.com/gin-gonic/gin"
  "errors"
  "strconv"
)

var router *gin.Engine

var articleList = []models.Article{
  models.Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
  models.Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getArticleByID(id int) (*models.Article, error) {
   for _, article := range articleList {
     if article.ID == id {
       return &article, nil
     }
   }

   return nil, errors.New("Article not found")
}

func indexPage(c *gin.Context) {
  data :=  gin.H{
  "title": "Aymen training",
  "articles": articleList,
  }
  render(c, data, "index.html")
}

func articleDetailPage(c *gin.Context) {
  if id, errorParameter := strconv.Atoi(c.Param("id")); errorParameter == nil {
    article, errorGetArticle := getArticleByID(id)
    if errorGetArticle != nil {
      c.AbortWithError(http.StatusNotFound, errorGetArticle)
    } else {
      data :=  gin.H{"article": article,}
      render(c, data, "articleDetail.html")
    }
  } else {
    c.AbortWithError(http.StatusBadRequest, errorParameter)
  }
}

func initializeRoute() {
  router.GET("/", indexPage)
  router.GET("/articles/view/:id", articleDetailPage)
}

func render(c *gin.Context, data gin.H, template string) {
  contentType := c.Request.Header.Get("Accept")

  switch contentType {
  case "application/json":
    c.JSON(http.StatusOK, data)
  case "application/xml":
    c.XML(http.StatusOK, data)
  default:
    c.HTML(http.StatusOK, template, data)
  }
}

func main() {
  router = gin.Default()
  router.LoadHTMLGlob("templates/*")

  initializeRoute()

  router.Run("localhost:8181")
}
