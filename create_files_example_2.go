package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	message := []byte("Look!")
	err := ioutil.WriteFile("test.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
