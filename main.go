package main

import (
	"github.com/takutakukatokatojapan/go_redis/injector"
	"log"
)

func main() {
	if err := injector.Run(); err != nil {
		log.Fatal("can not run server...")
	}
}
