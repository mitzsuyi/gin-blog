package db

import  (
 "api/models"
 "api/config"
 "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var dbConfig *config.DbConfig
var err error

func Migrate(resetDb * bool){
  if *resetDb {
    db.Debug().DropTableIfExists(&models.Article{}, models.Comment{})   
  } 
  db.Debug().AutoMigrate(&models.Article{}, &models.Comment{})
}

func InitDB(_config *config.DbConfig) error{
  dbConfig = _config 
  db, err = gorm.Open("mysql", _config.Mysql.Url)
  return err
}

func Close(){
  db.Close()
}
