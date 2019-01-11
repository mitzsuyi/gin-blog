package models

import  (
    "time"
)

type Comment struct {
  Id  uint `gorm:"primary_key"`  
  Nickname     string `gorm:"not null" json:"nickname" binding:"required"`
  Content      string `gorm:"not null; size:999999" json:"content" binding:"required"`
  CreatedAt time.Time `gorm:"not null" json:"creatdAt"` 
  Comments []Comment `json:"comments" gorm:"foreignkey:CommentId"`
  ArticleId uint `json:"-"`
  CommentId uint `json:"-"`
}

type Article struct {
  Id  uint `gorm:"primary_key"`  
  Nickname     string `gorm:"not null" json:"nickname" binding:"required"`
  Title        string `gorm:"not null" json:"title" binding:"required"`
  Content      string `gorm:"not null;size:999999" json:"content;-" binding:"required"`
  CreatedAt time.Time `gorm:"not null" json:"createdAt"`
  Comments []Comment `json:"comments,omitempty" gorm:"PRELOAD:false;foreignkey:ArticleId"`
}

