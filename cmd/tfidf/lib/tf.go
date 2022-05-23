package lib

import "strings"

func TF(word string, document string) float64 {

	timesWordInDocument := strings.Count(document, word)
	wordsInDocument := len(strings.Fields(document))

	return float64(timesWordInDocument) / float64(wordsInDocument)
}
