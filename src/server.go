package main

import (
  "local.package/infrastructure"
)

func main() {
	err := infrastructure.Router.Run()
	if err != nil {
		return
	}
}
