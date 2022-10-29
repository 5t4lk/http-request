package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	starter := time.Now()
	urls := []string{
		"https://google.com/",
		"https://youtube.com/",
		"https://faceit.com",
		"https://medium.com/",
		"https://github.com/",
		"https://hltv.org/",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		go func(url string) {
			wg.Add(1)
			doHTTP(url, starter)
			wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Printf("< < < Parsing completed. Time elapsed: %.2f seconds > > >", time.Since(starter).Seconds())
}

func doHTTP(url string, starter time.Time) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("<%s> ~ Status Code [%d] ~ No response ~ Ping: %d ms\n", url, r.StatusCode, time.Since(starter).Milliseconds())
	}

	defer r.Body.Close()

	fmt.Printf("<%s> ~ Status Code [%d] ~ Response OK ~ Ping: %d ms\n", url, r.StatusCode, time.Since(starter).Milliseconds())
}
