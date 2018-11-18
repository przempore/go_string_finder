package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Reader struct {
	DataInString    string
	DataInRawString string
	DataInLines     []string
}

func (r *Reader) GetFileData(file_path *string) error {
	data, err := ioutil.ReadFile(*file_path)
	if err != nil {
		fmt.Println("Can't read file:", file_path)
		return err
	}

	r.DataInRawString = string(data)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(r.DataInRawString, "\n")
	for _, l := range lines {
		if l == "" {
			continue
		}
		r.DataInLines = append(r.DataInLines, l)
	}

	r.DataInString = reg.ReplaceAllString(string(data), " ")

	return nil
}

func (r *Reader) PrintFileContent() {
	fmt.Println("File content is:")
	fmt.Println(r.DataInRawString)
}

func (r *Reader) FindString(string_to_find *string) string {
	maxLenght := 0
	for _, l := range r.DataInLines {
		words := strings.Fields(l)
		length := len(words)
		if length > maxLenght {
			maxLenght = length
		}
	}
	words := make([][]string, maxLenght)
	for ln, l := range r.DataInLines {
		wordsInLine := strings.Fields(l)
		words[ln] = make([]string, len(wordsInLine))
		for wn, w := range wordsInLine {
			words[ln][wn] = w
		}
	}

	fmt.Println("-------------------------------- words --------------------------------")
	fmt.Println(words)
	fmt.Println("-------------------------------- words --------------------------------")

	for _, w := range words {
		fmt.Println(w)
	}

	return "Hello, world"
}

func Hello() string {
	return "Hello, world"
}
