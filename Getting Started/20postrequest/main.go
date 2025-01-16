package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the server!! Here we are going to make a POST request to a server")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the URL to make a POST request")
	posturl, err := reader.ReadString('\n')
	ErrorHandle(err)
	posturl = strings.TrimSpace(posturl)
	user := postString{postUrl: posturl}
	user.PostRequest()

}

type postString struct {
	postUrl string
}

func ErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func (r *postString) PostRequest() {

	responsebody := strings.NewReader(`
			   {
			"name": "John",
			"age":30	
	}

	`)

	response, err := http.Post(r.postUrl, "application/json", responsebody)
	ErrorHandle(err)
	defer response.Body.Close()

	fmt.Println("Status :", response.Status)
	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
}
