package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/robinson/gos7"
)

func Test_s7(t *testing.T) {
	handler := gos7.NewTCPClientHandler("127.0.0.1:1800", 0, 1)

	defer handler.Close()
	if err := handler.Connect(); err != nil {
		t.Error(err)
		return
	}
	client := gos7.NewClient(handler)
	dataBuf := make([]byte, 10)
	if err := client.AGReadDB(10, 0, 10, dataBuf); err != nil {
		t.Error(err)
		return
	}
	t.Log("client.AGReadDB =>", dataBuf)

}
