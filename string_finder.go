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

func (r *Reader) FindString(stringToFind *string) []string {
	maxLength := getMaxLength(&r.DataInLines)

	words := make([][]string, maxLength)
	var retSlice []string
	for ln, l := range r.DataInLines {
		var modifiedLine string
		x := r.findInLine(&l, stringToFind, &modifiedLine)
		whatsLeft := l[x+len(*stringToFind):]
		for x != -1 {
			if x < len(whatsLeft) {
				whatsLeft = whatsLeft[x+len(*stringToFind):]
				x = r.findInLine(&whatsLeft, stringToFind, &modifiedLine)
			} else {
				x = -1
			}
		}
		retSlice = append(retSlice, modifiedLine)
		line := strings.Fields(l)
		words[ln] = make([]string, len(line))
		for wn, w := range line {
			words[ln][wn] = w
		}
	}

	return retSlice
}

func (r *Reader) findInLine(l *string, stringToFind *string, modifiedLine *string) int {
	x := -1
	if strings.Contains(strings.ToLower(*l), strings.ToLower(*stringToFind)) {
		x = strings.Index(strings.ToLower(*l), *stringToFind)
		stringToFindLength := len(*stringToFind)
		newLine := (*l)[:x] + " >>>> " + *stringToFind + " <<<<" + (*l)[x+stringToFindLength:]
		fmt.Println(newLine)
		//fmt.Fprintf(os.Stderr, "%s\n", newLine)
		*modifiedLine = *modifiedLine + newLine
	}
	return x
}

func getMaxLength(dataInLines *[]string) int {
	maxLength := 0
	for _, l := range *dataInLines {
		words := strings.Fields(l)
		length := len(words)
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func Hello() string {
	return "Hello, world"
}
