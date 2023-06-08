package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file system in golang")
	fmt.Println("---------------------------------")

	// here is the content whiich we have to put inside a file.
	content := "This needs to go in a file - niteshtudu.com"

	// the comma ok systax [because there may be some error while creating a file.]
	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) // this return length and err if present
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Length is : ", length)
	defer file.Close() // after done working with the file it is recommended to close the file.
	// and if we use the defer keyword then the file will be closing at the end of the scope wich is really good.

	readFile("./myfile.txt")
}

func readFile(filename string) { // here we have send the full path of the file too with the file name.
	// for reading and writing in a file it is ok to use (io) module/binary.
	// and for creating the file is the (os) operation.
	// and for some manipulation we use the (ioutils) cause there are like so many methos.
	databyte, err := ioutil.ReadFile(filename)
	// when we read some data using ReadFile it come out in the foramt of byte format not in the string format.
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte)) // by putting the byte format data inside the string() it will parse to a sting.
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
