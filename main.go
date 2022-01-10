package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	profinet "github.com/Kowiste/ProfinetServer"
)

func main() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGABRT)
	server := profinet.NewServer()
	server.SetOutput([]uint16{11, 22, 33, 44, 55, 66, 77, 88, 99, 100})
	server.SetInput([]uint16{11, 22, 33, 44, 55, 66, 77, 88, 99, 100})
	server.SetDB(13, []uint16{11, 22, 33, 44, 55, 66, 77, 88, 99, 100})
	err := server.Listen("0.0.0.0:1503", 0, 1)
	if err != nil {
		log.Println(err)
		return
	}
	signal := <-channel
	context.Background().Done()
	log.Println("Received exit signal: ", signal)
}
