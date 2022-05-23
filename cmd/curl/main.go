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

	URLSplit := strings.Split(URL, "/")
	path := "data/" + URLSplit[len(URLSplit) - 1]

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(path, data, 0770)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("File " + path + " downloaded")
}