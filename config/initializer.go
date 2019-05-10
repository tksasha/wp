package config

import (
  "os"
  "bytes"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"

  . "github.com/tksasha/wp/config/db"
)

var workDir bytes.Buffer

func init() {
  prepareWorkDir()

  prepareFiles()

  prepareDB()
}

func prepareWorkDir() {
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

func prepareFiles() {
  files := [2]string { "/PHRASE", "/ANSWER" }

  for _, file := range files {
    path := workDir

    path.WriteString(file)

    if fd, err := os.OpenFile(path.String(), os.O_CREATE, 0644); err != nil {
      panic(err)
    } else {
      defer fd.Close()
    }
  }
}

func prepareDB() {
  var err error

  database := workDir

  database.WriteString("/database.sqlite3")

  if DB, err = gorm.Open("sqlite3", database.String()); err != nil {
    panic("failed to connect database")
  }
}
