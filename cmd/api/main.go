package main

import (
	"tokenize/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8000")
}
