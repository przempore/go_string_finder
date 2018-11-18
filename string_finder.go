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
	r.DataInString = reg.ReplaceAllString(string(data), " ")

	return nil
}

func (r *Reader) PrintFileContent() {
	//fmt.Println("File content is:")
	//fmt.Println(r.DataInRawString)
}

func (r *Reader) FindString(string_to_find *string) string {
	tmp := strings.Split(r.DataInString, "\n")
	fmt.Println(tmp)

	fmt.Printf("lines: %d\nWhich is:\n", len(r.DataInLines))
	for _, l := range r.DataInLines {
		fmt.Println(l)
	}

	fmt.Println("the end of lines")

	words := make([][]string, len(r.DataInLines))
	for ln, l := range r.DataInLines {
		wordsInLine := strings.Fields(l)
		words[ln] = make([]string, len(wordsInLine))
		for wn, w := range words[ln] {
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
