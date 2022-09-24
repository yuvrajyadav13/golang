package main

import (
	"encoding/json"
	"fmt"
)

// Normal struct
type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

// Go provides a support for Json
// Here we have created a separate struct for that
type aliascourse struct {
	// By using `` or double backticks we can provide alias to any attribute in struct.
	// here name at the time of json encode will convert Name to coursename
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`

	// Here in password using "-" as alias in json will completely remove that filled from json
	Password string `json:"-"`
	// Here in tags, if we simple give `json:tags`, whenever tags are nil json will convert them
	// into "null" and response you will be getting is tags: null
	// To remove null filleds from json we use "omitempty" or `json:tags,omitempty`
	// Field with omitempty won't show up in json if they are empty
	// Remember there should not be any space between alias and omitempty
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in GO")
	EncodeJson()
	EncodeJson2()
	EncodeJson3()
	decodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Mern", 99, "lco.in", "a454b3", []string{"web", "stack"}},
		{"ReactJs", 299, "lco.in", "acfd53", nil},
	}

	// packaging data as JSON

	finalJson, err := json.Marshal(courses)

	if err != nil {
		panic(err)
	}

	fmt.Println(finalJson)
	fmt.Printf("%s\n", finalJson)
}

func EncodeJson2() {
	courses := []course{
		{"ReactJs", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Mern", 99, "lco.in", "a454b3", []string{"web", "stack"}},
		{"ReactJs", 299, "lco.in", "acfd53", nil},
	}

	// To format JSON instead of simply using Marshal we use MarshelIndent
	// MarshalIndent takes 3 arguments. data, any prefixes to be added in every line and by what
	// values it should be indented with
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func EncodeJson3() {
	courses := []aliascourse{
		{"ReactJs", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Mern", 99, "lco.in", "a454b3", []string{"web", "stack"}},
		{"ReactJs", 299, "lco.in", "acfd53", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func decodeJson() {
	jsonData := []byte(`
		{
			"coursename": "ReactJs",
			"Price": 299,
			"website": "lco.in",
			"tags": [
			"web-dev",
			"js"
			]
		}
	`)

	var lcocourse aliascourse

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Printf("JSON is not valid")
	}

	// Json to key value or map format
	var jsonMapData map[string]interface{}
	json.Unmarshal(jsonData, &jsonMapData)
	fmt.Println(jsonMapData)

	for key, val := range jsonMapData {
		fmt.Printf("Key is %v and value is %v. Type = %T\n", key, val, val)
	}
}
