package test

import (
	"os"
	"strings"
	"testing"

	"tfidf/lib"
)

func TestTF(t *testing.T) {

	data, err := os.ReadFile("data_for_test/document_1.txt")
	if err != nil {
		t.FailNow()
	}

	document := strings.ToLower(string(data))
	words := []string{
		"nunc",
		"test",
	}

	tf1 := lib.TF(words[0], document)
	tf2 := lib.TF(words[1], document)

	if tf1 != 0.016483516483516484 {
		t.Fail()
	}

	if tf2 != 0 {
		t.Fail()
	}
}

func TestPath(t *testing.T) {

	files := []string{
		"/this/is/a/file",
		"this/is/also/a/file",
		"thisIsAFileToo",
	}

	path1 := lib.Path(files[0])
	path2 := lib.Path(files[1])
	path3 := lib.Path(files[2])

	if path1 != "this/is/a" {
		t.Fail()
	}

	if path2 != "this/is/also/a" {
		t.Fail()
	}

	if path3 != "" {
		t.Fail()
	}

}

func TestWordInDocument(t *testing.T) {

	words := []string{
		"nunc",
		"test",
	}
	file := "data_for_test/document_1.txt"

	word1 := lib.WordInDocument(file, words[0])
	word2 := lib.WordInDocument(file, words[1])

	if !word1 {
		t.Fail()
	}

	if word2 {
		t.Fail()
	}

}

func TestIDF(t *testing.T) {

	words := []string{
		"nunc",
		"test",
	}
	dir := "data_for_test"

	files, err := os.ReadDir(dir)
	if err != nil {
		t.FailNow()
	}

	idf1 := lib.IDF(words[0], dir, files)
	idf2 := lib.IDF(words[1], dir, files)

	if idf1 != -0.40546510810816444 {
		t.Fail()
	}

	if idf2 != 0 {
		t.Fail()
	}
}