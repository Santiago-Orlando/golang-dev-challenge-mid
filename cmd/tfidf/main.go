package main

import (
	"fmt"
	"os"
	"strings"

	"tfidf/lib"
)



func main() {

	args := os.Args[:]

	word := strings.ToLower(args[1])
	file := args[2]

	dir := lib.Path(file)

	if dir == "" {
		fmt.Println("The files cannot be in the root of the directory.")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	document := strings.ToLower(string(data))

	tf := lib.TF(word, document)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	idf := lib.IDF(word, dir, files)

	fmt.Println(tf*idf)
}
