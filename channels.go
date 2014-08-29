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
	completedAt := <-timeChan
	fmt.Println("completedAt:", completedAt)
}
