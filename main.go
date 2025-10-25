package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	quit := make(chan os.Signal, 1)

	go func() {
		defer wg.Done()
		fmt.Println("Starting 1")
		for len(quit) < 1 {
			time.Sleep(500)
		}
		fmt.Println("Ending 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Starting 2")
		for len(quit) < 1 {
			time.Sleep(500)
		}
		fmt.Println("Ending 2")
	}()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	wg.Wait()
	fmt.Println("Shutting Down")
}
