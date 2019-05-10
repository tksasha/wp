package config

import (
  "os"
  "bytes"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"

  . "github.com/tksasha/wp/config/db"
)

func init() {
  prepareWorkDir()

  prepareDB()
}

func prepareWorkDir() {
  var workDir bytes.Buffer

  if dir, err := os.UserHomeDir(); err == nil {
    workDir.WriteString(dir)
  } else {
    panic(err)
  }

  workDir.WriteString("/.wp")

  if err := os.MkdirAll(workDir.String(), 0755); err != nil {
    panic(err)
  }
}

func prepareDB() {
  var err error

  var database bytes.Buffer

  if dir, err := os.UserHomeDir(); err == nil {
    database.WriteString(dir)
  } else {
    panic(err)
  }

  database.WriteString("/.wp/database.sqlite3")

  if DB, err = gorm.Open("sqlite3", database.String()); err != nil {
    panic("failed to connect database")
  }
}
