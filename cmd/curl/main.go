package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)



func main() {

	URL := os.Args[1]

	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	filename := strings.Split(URL, "/")
	name := "data/" + filename[len(filename) - 1]
	
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(name, data, 0770)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("File " + name + " downloaded")
}