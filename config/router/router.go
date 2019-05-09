package config

import (
  "fmt"
  "os"

  . "github.com/tksasha/wp/app/controllers"
)

func Router() {
  if len(os.Args) < 2 {
    fmt.Println("Use: ", os.Args[0], " [add|]")

    return
  }

  switch os.Args[1] {
    case "add":
      new(PhrasesController).Create()
    case "list":
      new(PhrasesController).Index()
    default:
      fmt.Println("undefined command")
  }
}
