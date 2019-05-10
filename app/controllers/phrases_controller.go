package controllers

import (
  "fmt"
  "os"
  "os/exec"
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

func (c PhrasesController) Show() {
  resource := new(Phrase)

  switch len(os.Args) {
    case 2:
      DB.Order("random()").Limit(1).Find(resource)

      fmt.Println("Phrase ID =", resource.ID)

      exec.Command("say", resource.Text).Run()
    case 3:
      DB.First(resource, os.Args[2])

      if err := json.NewEncoder(os.Stdout).Encode(resource); err != nil {
        panic(err)
      }
    default:
      fmt.Println("Do Nothing!")
  }

}
