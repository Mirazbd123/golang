package main

import (
	"encoding/json"
	"fmt"
)

// First i have to make a struct for json
type course struct {
	Name     string `json:"Coursename"` // Veriable in the struct always start with capital letter , other wise it will show error
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"` // It will hide the password
	Tags     []string `json:"FullFrom"`
}

func main() {
	fmt.Println("Welcome to json!!!!")
	EncodeJson()

}

func EncodeJson() {
	mycourses := []course{
		{"Mern Stack Bootcamp", 3000, "Ostadh.com", "armans323", []string{"Mongodb", "Express Js", "React js", "Node js"}},
		{"Pern Stack Bootcamp", 4000, "Ostadh.com", "armadfns323", []string{"ProsgreSql", "Express Js ...", "React js...", "Node js..."}},
		{"Gern Stack Bootcamp", 7000, "Ostadh.com", "armanwws323", nil},
	}

	// Package these data in json
	finaljson, err := json.MarshalIndent(mycourses, "", "\t")
	ErrorHandle(err)
	fmt.Printf("%s\n", string(finaljson))
	// fmt.Println(mycourses)
}

func ErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
