package lib

import (
	"io/fs"
	"math"
)

func IDF(word string, dir string, files []fs.DirEntry) float64 {

	var filesWithWord float64
	var filesInDir float64 = float64(len(files) - 1)

	for _, file := range files {
		if WordInDocument(dir + "/" + file.Name(), word) {
			filesWithWord++
		}
	}

	if filesWithWord == 0 {
		return 0
	}

	return math.Log(filesInDir / filesWithWord)
}
