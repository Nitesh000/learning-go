package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ashfl4lka"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println("-------------------------------------")

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams) // the type will be url.Value, it is like the key, value pair
	// and we can get the result like a map.

	fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUlr := &url.URL{ // NOTE: we have to pass the reference not a copy
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=nitesh",
	}

	anotherURl := partsOfUlr.String()

	fmt.Println(anotherURl)
}
