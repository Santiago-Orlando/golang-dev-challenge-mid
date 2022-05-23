package lib

import (
	"fmt"
	"os"
	"strings"
)

func WordInDocument(file string, word string) bool {

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	document := strings.ToLower(string(data))

	wordsInDocument := strings.Split(document, " ")

	for _, wordInDocument := range wordsInDocument {
		if strings.Contains(wordInDocument, word) {
			return true
		}
	}

	return false

}
