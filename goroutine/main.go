package main

import (
	"fmt"
	"log"
	"mtsbank_golang/goroutine/pinger"
	"os"
	"sync"
	"time"
)

func main() {
	PingWithOneChahnel()
	//PingWithChannelPerAddress()
}

// one chan for all addr
func PingWithOneChahnel() {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("require at least one argument")
	}

	dataCh := make(chan string, 5)

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

// chan per addr
func PingWithChannelPerAddress() {

	args := os.Args[1:]
	l := len(args)

	if l < 1 {
		log.Fatal("require at least one argument")
	}

	channels := make([]chan string, l)

	for j, addr := range args {
		channels[j] = make(chan string)
		go func(addr string, dataCh chan<- string) {
			for i := 1; i < 4; i++ {
				dst, dur, err := pinger.Ping(addr)
				if err != nil {
					continue
				}
				dataCh <- fmt.Sprintf("Ping %s (%s): %s", addr, dst, dur)
				close(dataCh)
				return
			}

			dataCh <- fmt.Sprintf("Couldn't connect to %s", addr)
			close(dataCh)

		}(addr, channels[j])
	}

	for i := 0; len(channels) != 0; i++ {
		select {
		case v := <-channels[i]:
			fmt.Println(v)
			channels = remove(channels, i)
			i--
		case <-time.After(1 * time.Second):
			//fmt.Println("after 1 sec")
		}

		if i == len(channels)-1 {
			i = -1
		}
	}
}

func remove(slice []chan string, s int) []chan string {
	return append(slice[:s], slice[s+1:]...)
}
