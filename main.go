package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	profinet "github.com/Kowiste/ProfinetServer"
)

var db = []uint16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var rack, slot = uint16(0), uint16(1)

func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGABRT)
	server := profinet.NewServer()
	server.SetDB(10, db)

	if err := server.Listen("0.0.0.0:1503", rack, slot); err != nil {
		log.Println(err)
		return
	}
	ticker := time.NewTicker(time.Duration(time.Second * 5))
	cctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		fmt.Print("DB: ")
		for i := range db {
			fmt.Printf(" %v ", db[i])
		}
		fmt.Printf("\r")
		for {
			<-ticker.C
			select {
			case <-cctx.Done():
				return
			default:
				{
				}
			}
			for i := range db {
				v := rand.Intn(256)
				db[i] = uint16(v)
			}
			fmt.Print("DB: ")
			for i := range db {
				fmt.Printf(" %v ", db[i])
			}
			fmt.Printf("\r")
		}

	}(cctx)
	signal := <-channel
	cancel()
	log.Println("Received exit signal: ", signal)
}
