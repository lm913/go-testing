package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type response struct {
	url   string
	err   error
	count int
}

func worker(urls <-chan string, resps chan<- *response, wg *sync.WaitGroup) {
	// grab urls until the urls channel is closed.
	for url := range urls {
		res, err := http.Get(url)
		if err != nil {
			resps <- &response{url, err, 0}
			continue
		}
		slurp, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			resps <- &response{url, err, 0}
			continue
		}
		resps <- &response{
			url,
			nil,
			bytes.Count(slurp, []byte("Go")),
		}
	}

	wg.Done() // decrement wg
}

func main() {
	maxWorkers := 5 // could be flag instead.

	urls := make(chan string)
	resps := make(chan *response)

	var wg sync.WaitGroup

	// start maxWorkers workers.
	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go worker(urls, resps, &wg)
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
		// input is done, close urls channel.
		close(urls)

		// wait for wg counter to be 0 and then close
		// the resps channel to exit the for loop below.
		wg.Wait()
		close(resps)
	}()

	var total int
	for r := range resps {
		if r.err != nil {
			log.Printf("%s: %s\n", r.url, r.err)
			continue
		}
		total += r.count
		fmt.Printf("Count for %s: %d\n", r.url, r.count)
	}
	fmt.Println(total)
}
