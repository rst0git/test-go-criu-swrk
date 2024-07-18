package main

import (
	"log"

	"github.com/checkpoint-restore/go-criu/v7"
)

func main() {
	c := criu.MakeCriu()
	version, err := c.GetCriuVersion()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(version)
}
