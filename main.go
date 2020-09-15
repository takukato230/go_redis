package main

import (
	"go_redis/injector"
	"log"
)

func main() {
	if err := injector.Run(); err != nil {
		log.Fatal("can not run server...")
	}
}
