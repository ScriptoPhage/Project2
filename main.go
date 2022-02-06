package main

import (
	"Project2/wordfreq"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1]
	userInput, err := ioutil.ReadFile(args)
	if err != nil {
		panic(err)
	}

	wordfreq.WordCountService(string(userInput))
	//
}
