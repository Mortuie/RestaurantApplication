package main

import (
	"log"
	"restapp/order-api/bootstrap"
)

func main() {
	err := bootstrap.Setup()

	if err != nil {
		log.Fatal(err)
	}
}
