package config

type DbConfig struct {
   Mysql struct {
     Url string `json:"Url"`
   } `json:"Mysql"`
  Paginate int `json:"Paginate"`
}

type ApiConfig struct {
    Db DbConfig `json:"Db"`
}