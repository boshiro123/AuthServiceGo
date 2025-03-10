package main

import (
	"fmt"
	"urlshortener/pkg/config"
)

func main() {
	fmt.Println("its start working")

	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: config

	// TODO: logger

	// TODO: initialize app

	// TODO: run gr
}
