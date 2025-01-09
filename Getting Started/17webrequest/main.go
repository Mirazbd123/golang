package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myurl = "https://ediann.xyz"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(myurl)
	errorHandle(err)

	fmt.Printf("Response is of type: %T\n", response)    // response is of type *http.Response
	fmt.Printf("Response is: %v\n", response.StatusCode) // Print the status code of the response
	body, err := ioutil.ReadAll(response.Body)           // Read the body as datatype --> byte of the response
	errorHandle(err)
	fmt.Printf("Body is: %v\n", string(body)) // Convert the byte to string

	response.Body.Close() // Close the response body always

}

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
