package main

import (
	"fmt"
	"leadgen/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("qwerty")
	app := internal.NewApp()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	app.Run(sigterm)
}
