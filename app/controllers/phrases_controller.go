package controllers

import (
  "fmt"
  "os"
  "encoding/json"

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

func (c PhrasesController) Index() {
  var collection []Phrase

  DB.Find(&collection)

  if err := json.NewEncoder(os.Stdout).Encode(collection); err != nil {
    panic(err)
  }
}
