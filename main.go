package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Project2/server"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Connected on port 8000")
	srv := server.SetupRoutes()
	err := srv.Run(":8000")
	if err != nil {
		panic(err)
	}
}
