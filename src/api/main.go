package main

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "api/handlers"
    "api/config"
    "flag"
    "api/db"
    "utils"
    "log"
)

var configDir = "api/"

func main() {
   
   configFile := utils.ConfigFile(configDir)

   apiConfig := config.ApiConfig{}

   utils.GetConfigFatal(&apiConfig, &configFile)
   dbConfig := apiConfig.Db  
   err := db.InitDB(&dbConfig)

   if err != nil {
     log.Fatalln("error connecting to db", err)
   }

   var port = flag.Int("p", 8080, "port 8080")
   var resetDb = flag.Bool("reset-db", false, "reset database false")
   
   flag.Parse()

   db.Migrate(resetDb)

   defer db.Close()

   router := gin.Default()

   articles := router.Group("/articles")

   {
    articles.GET("", handlers.GetArticles)
    articles.POST("", handlers.CreateArticle)
    articles.POST(":id/comments", handlers.CreateArticleComment)
    articles.GET(":id/content", handlers.GetArticleContent)
    articles.GET(":id/comments", handlers.GetArticleComments)
   }

   comments := router.Group("/comments")
   
   {
     comments.POST(":id/comment", handlers.CreateCommentComment)

   }

   router.Run(":" + strconv.Itoa(*port))
}