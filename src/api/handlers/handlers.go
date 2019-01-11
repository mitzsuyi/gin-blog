package handlers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api/db"
  "api/models"
  "strconv"
)

func GetArticles(c *gin.Context){
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    paginator, err := db.GetArticles(page)
    respondWithJSON(http.StatusOK, paginator, err, c)
}

func CreateArticle(c *gin.Context){
    var article models.Article
    if err := bindJSON(&article, c, "Article"); err != nil{
        return
    }
    err := db.CreateArticle(&article)
    respondWithJSON(http.StatusCreated, gin.H{"id":article.Id}, err, c)
}

func CreateArticleComment(c *gin.Context){
    if article, err := getArticle(c); err == nil {
      createCommentAssociation(&article, c, "Article")   
    }
}

func GetArticleComments(c *gin.Context){
  if article, _err := getArticle(c); _err == nil {
    comments, err := db.GetCommentAssociation(article, "Article") 
    respondWithJSON(http.StatusOK, gin.H{"data":comments}, err, c)
  }
}

func GetArticleContent(c *gin.Context){
  if article, err := getArticle(c); err == nil {
   respondWithJSON(http.StatusOK, gin.H{"data":article.Content}, nil, c)
  }
}

func CreateCommentComment(c *gin.Context) {
   if comment, err := getComment(c); err == nil{
     createCommentAssociation(&comment, c, "Comment")
   }
}