package handlers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  api_error "api/error"
  "api/db"
  "errors"
  "strconv"
  "api/models"
)

func abortWithJSON(err *api_error.Error, c *gin.Context){
    c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Err.Error()})
}

func respondWithJSON(status int, h interface{}, err *api_error.Error, c *gin.Context){
  if err != nil {
    abortWithJSON(err, c)
  } else {
    c.JSON(status, h)
  }
}
 
func getParamId(c *gin.Context) (int, * api_error.Error) {
  _id := c.Param("id")
  id, _err := strconv.Atoi(_id)  
  if (_err != nil){
    err := api_error.BadRequest(errors.New("Invalid id"))
    abortWithJSON(err, c)
    return id, err
  }
  return id, nil
}

func getArticle(c * gin.Context) (* models.Article, * api_error.Error){
  var article models.Article
  err := getResourceById(&article, c, "Article")
  return &article, err
}

func getComment(c * gin.Context) (* models.Comment, * api_error.Error){
  var comment models.Comment
  err := getResourceById(&comment, c, "Comment")
  return &comment, err
}

func getResourceById(model interface{}, c *gin.Context, resource string) * api_error.Error {
   id, err := getParamId(c)
   if (err != nil){
     return err
   }
   if err = db.GetResourceById(model, id, resource); err != nil{
     abortWithJSON(err, c)
   }
   return err
}

func bindJSON(model  interface{}, c * gin.Context, resource string) * api_error.Error {
   if _err := c.BindJSON(model); _err != nil{
      err := api_error.BadRequest(errors.New("Invalid " + resource))
      abortWithJSON(err, c)
      return err
    }
    return nil
}

func createCommentAssociation(model interface{}, c *gin.Context, resource string) {
    var comment models.Comment
    if err := bindJSON(&comment, c, "Comment"); err == nil{
      err = db.CreateCommentAssociation(model, &comment, resource)
      respondWithJSON(http.StatusCreated, gin.H{"id":comment.Id}, err, c)
    }
}