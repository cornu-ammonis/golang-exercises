//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, ch chan<- *Tweet, wg *sync.WaitGroup) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(ch)
			wg.Done()
			return
		}

		ch <- tweet
	}
}

func consumer(ch <-chan *Tweet, wg *sync.WaitGroup) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}

	wg.Done()
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// buffer would not help because producer takes longer than consumer
	ch := make(chan *Tweet)

	var wg sync.WaitGroup

	wg.Add(2)

	// Producer
	go producer(stream, ch, &wg)

	// Consumer
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
