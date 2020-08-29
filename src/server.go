package main

import (
  "github.com/masakazu1967/office-library_book-service/infrastructure"
)

func main() {
	err := infrastructure.Router.Run()
	if err != nil {
		return
	}
}
