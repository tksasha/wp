package config

import (
  "os"
  "bytes"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"

  . "github.com/tksasha/wp/config/db"
)

func init() {
  var err error

  var database bytes.Buffer

  database.WriteString(os.Getenv("HOME"))

  database.WriteString("/.wp.sqlite3")

  DB, err = gorm.Open("sqlite3", database.String())

  if err != nil {
    panic("failed to connect database")
  }
}
