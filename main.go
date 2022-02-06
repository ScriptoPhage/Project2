package main

import (
	"Project2/wordfreq"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1]
	userInput, err := ioutil.ReadFile(args)
	if err != nil {
		panic(err)
	}

	jsonOutput := wordfreq.WordCountService(string(userInput))
	fmt.Println(string(jsonOutput))
	//
}
