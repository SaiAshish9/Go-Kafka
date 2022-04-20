package main

import (
	"fmt"
	"time"

	appKafka "./kafka"
)

func main() {
	fmt.Print("go kafka")
	go appKafka.StartKafka()
	fmt.Println("kafka has been started")
	time.Sleep(10 * time.Minute)
}
