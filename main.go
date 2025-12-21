package main

import (
	"fmt"
	"sync"

	"github.com/LiminalFish/kempt/internal/api"
)

var wg sync.WaitGroup

func main() {
	wg.Go(func() {
		api.StartServer()
	})

	fmt.Println("squid")

	wg.Wait()
}
