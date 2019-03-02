package main

import (
	"log"

	"github.com/linehk/go-douban-movies/model"
	"github.com/linehk/go-douban-movies/router"
)

func main() {
	model.Init()
	r := router.InitRouter()
	if err := r.Run(":8888"); err != nil {
		log.Fatalf("run server err: %v", err)
	}
}
