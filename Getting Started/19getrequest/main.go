package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the server!! Here we are going to make a GET request to a server")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the URL to make a GET request")
	geturl, err := reader.ReadString('\n')
	ErrorHandle(err)
	geturl = strings.TrimSpace(geturl)
	user := getString{getUrl: geturl}
	user.GetRequest()

}

type getString struct {
	getUrl string
}

func (r *getString) GetRequest() {
	response, err := http.Get(r.getUrl)
	ErrorHandle(err)
	defer response.Body.Close()
	fmt.Println("Status of the respose:", response.Status)
	// fmt.Println(response.Body) // This will print the address of response body

	// content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println("Content of the response:", string(content))
	// fmt.Printf("Content of the response: %T\n", string(content)) // Json of string

}

func ErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
