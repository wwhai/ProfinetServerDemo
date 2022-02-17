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

//
// Db是10个 short ，实际上是 20个字节
//
var db = []uint16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var rack, slot = uint16(0), uint16(1)

func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGABRT)
	server := profinet.NewServer()
	if err := server.Listen("0.0.0.0:1800", rack, slot); err != nil {
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
			server.SetDB(10, db)
			fmt.Printf("\r")
		}

	}(cctx)
	signal := <-channel
	cancel()
	log.Println("Received exit signal: ", signal)
}
