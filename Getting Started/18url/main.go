package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=1"

func main() {
	fmt.Println("URL:", myurl)

	result, err := url.Parse(myurl)
	errorHandle(err)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Query:", result.RawQuery)
	fmt.Println("Fragment:", result.Fragment)

	qparams := result.Query()
	fmt.Printf("Type of Query params: %T\n", qparams)
	fmt.Println("Query params:", qparams)

	for key, value := range qparams {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

}

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
