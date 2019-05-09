package main

import (
  . "github.com/tksasha/wp/config/db"
  _ "github.com/tksasha/wp/config"

  . "github.com/tksasha/wp/config/router"
)

func main() {
  defer DB.Close()

  Router()
}
