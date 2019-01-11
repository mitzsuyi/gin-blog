package db

import (
"api/models"
 api_error "api/error"
 "net/http"
 "errors"
  "github.com/jinzhu/gorm"
  "github.com/mitzsuyi/gorm-paginator/pagination"
)

func paginate(model interface{}, page int) (* pagination.Paginator, error) {
  paginator, err := pagination.Pagging(&pagination.Param{
        DB:      db,
        Page:    page,
        Limit:   dbConfig.Paginate,
        OrderBy: []string{"id desc"},
        ShowSQL: true,
    }, model)  
  return paginator, err
}

func checkErrors(dberror error, resource string) * api_error.Error{
  if (gorm.IsRecordNotFoundError(dberror)){
    return api_error.ApiError(http.StatusNotFound, errors.New(resource + " could not be found"))
  } else {
    return api_error.InternalError(dberror)
  }
}

func GetResourceById(model interface {}, id int, resource string) * api_error.Error {
  return checkErrors(db.First(model, id).Error, resource)
}

func GetArticles(page int) (* pagination.Paginator,  * api_error.Error) {
    var articles  []models.Article 
    paginator, err := paginate(&articles, page)
    return paginator, api_error.InternalError(err)
}

func CreateArticle(article * models.Article) (* api_error.Error) {
    return api_error.InternalError(db.Create(article).Error)
}

func CreateCommentAssociation(model interface{}, comment * models.Comment, resource string) (* api_error.Error) {
    return checkErrors(db.Model(model).Association("Comments").Append(comment).Error, resource)
}

func GetCommentAssociation(model interface{}, resource string) (*[] models.Comment, * api_error.Error) {
    var comments []models.Comment
    return &comments, checkErrors(db.Set("gorm:auto_preload", true).Model(&model).Related(&comments).Error, resource)
}
