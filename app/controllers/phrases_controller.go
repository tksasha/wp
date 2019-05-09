package controllers

import (
  "fmt"
  "os"

  . "github.com/tksasha/wp/config/db"
  . "github.com/tksasha/wp/app/models"
)

type PhrasesController struct {}

func (c PhrasesController) Create() {
  if len(os.Args) < 4 {
    fmt.Println("Use: ", os.Args[0], "add PHRASE TRANSLATION")

    return
  }

  DB.Create(&Phrase { Text: os.Args[2], Translation: os.Args[3] })
}
