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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() { //whenever any data comes from the web it is in the byte format

	jsonDataFromWeb := []byte(`
			{
		"Coursename": "React js Bootcamp",
		"Price": 4648,
		"Website": "lolo.com",
		"Password": "3434dff",
		"FullFrom": ["web","app"]
	}
	`)

	var yourCourses course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("Json is valid!!")
		json.Unmarshal(jsonDataFromWeb, &yourCourses)
		fmt.Printf("%#v\n", yourCourses)
	} else {
		fmt.Println("Json is not valid")
	}

	// There is another mathod where we use map instead of struct to make json --> key value pair
	var yourMapCourse map[string]interface{} //here interface means any kind od valu int or string or bool etc
	json.Unmarshal(jsonDataFromWeb, &yourMapCourse)
	for key, value := range yourMapCourse {
		fmt.Printf("Key is: %v and value is: %v\n", key, value)
	}
}

func ErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
