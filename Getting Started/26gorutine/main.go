package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("world")
	website := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://youtube.com",
		"https://amazon.com",
	}
	for _, web := range website {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

}

// func greeter(s string) {
// 	for i := 0; i < 20; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
