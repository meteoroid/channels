package main

import (
	"fmt"
	"time"
)

func main() {
	timeChan := make(chan time.Time)
	go func() {
		time.Sleep(time.Second)
		timeChan <- time.Now()
	}()
	fmt.Println("waiting for chan...")
	timeFromChan := <-timeChan
	fmt.Println("time from channel:", timeFromChan)
}
