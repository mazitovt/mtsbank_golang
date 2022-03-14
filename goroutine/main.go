package main

import (
	"fmt"
	"log"
	"mtsbank_golang/goroutine/pinger"
	"os"
	"sync"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("require at least one argument")
	}

	dataCh := make(chan string)

	wg := &sync.WaitGroup{}

	for _, addr := range args {
		wg.Add(1)

		go func(addr string, dataCh chan<- string, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 1; i < 4; i++ {
				dataCh <- fmt.Sprintf("Connecting to %s: try %d", addr, i)
				dst, dur, err := pinger.Ping(addr)
				if err != nil {
					continue
				}
				dataCh <- fmt.Sprintf("Ping %s (%s): %s", addr, dst, dur)
				return
			}

			dataCh <- fmt.Sprintf("Couldn't connect to %s", addr)

		}(addr, dataCh, wg)
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	for r := range dataCh {
		fmt.Println(r)
	}
}
