package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Sleeping...")
		time.Sleep(5 * time.Second)
	}
}
