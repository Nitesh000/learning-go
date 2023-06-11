package main

import (
	"encoding/json"
	"fmt"
)

// we can add a third param to the struct variables with some extra configuration on the encoding.
type course struct {
	Name     string `json:"coursenmae"` // we can add alias for the json in the struct
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // the '-' shows that we don't want this field to be consumed.
	Tags     []string `json:"tags,omitempty"` // omitempty:= it means we don't want to show if there is nil or null.
}

// encoding: convetinga a raw data type to a valid json.

func main() {
	fmt.Println("Welcome to JSON in golang")
	fmt.Println("---------------------------")
	// EncodingJson()
	DecodeJson()
}

func EncodingJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Vue Bootcamp", 299, "LearnCodeOnline.in", "bdc123", []string{"full-stack", "js"}},
		{"Anguale Bootcamp", 299, "LearnCodeOnline.in", "nti124", nil},
	}
	// package this data as JSON dat
	// json.Marshal := it will just create a JSON file and then return it, also if an error appears
	// json.MarshalIndent := it will just create a JSON file, with a prittier version and then return it, also if an error appears
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

// how to consume json data in golang

func DecodeJson() {
	jsonDataFromWeb := []byte(`
  {
    "coursenmae": "ReactJS Bootcamp",
    "Price": 299,
    "website": "LearnCodeOnline.in", 
    "tags": ["web-dev","js"]
  }
  `)

	var lcoCourse course

	// json.Valid := checking if the json is valid or not.
	// it returns a boolean value. If valid true else false.
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		// by using the Unmarshal we are decoding it from the byte format to a new structure.
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) // %#v is required to print it.
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value.

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v type of value is %T\n", k, v, v)
	}
}
