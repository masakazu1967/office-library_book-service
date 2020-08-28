package main

import (
	"local.package/infrastructure"
		_ "github.com/lib/pq"
)

func main() {
	infrastructure.Router.Run()
}
