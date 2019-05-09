package models

import (
  "github.com/jinzhu/gorm"

  . "github.com/tksasha/wp/config/db"
  _ "github.com/tksasha/wp/config"
)

func init() {
  DB.AutoMigrate(&Phrase{})
}

type Phrase struct {
  gorm.Model

  Text        string `gorm:"type:text;unique;not null"`
  Translation string `gorm:"type:text;unique;not null"`
}
