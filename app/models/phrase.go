package models

import (
  "time"

  . "github.com/tksasha/wp/config/db"
  _ "github.com/tksasha/wp/config"
)

func init() {
  DB.AutoMigrate(&Phrase{})
}

type Phrase struct {
  ID          uint        `gorm:"primary_key" json:"id"`
  CreatedAt   time.Time   `json:"-"`
  UpdatedAt   time.Time   `json:"-"`
  DeletedAt   *time.Time  `sql:"index" json:"-"`

  Text        string      `gorm:"type:text;unique;not null" json:"text"`
  Translation string      `gorm:"type:text;unique;not null" json:"translation"`
}
