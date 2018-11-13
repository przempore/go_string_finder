package main

import (
	"fmt"
	"io/ioutil"
)

type Reader struct {
	DataInString []string
}

func (r Reader) GetFileData(file_path *string) error {
	data, err := ioutil.ReadFile(*file_path)
	if err != nil {
		fmt.Println("Can't read file:", file_path)
		return err
	}

	for _, element := range data {
		r.DataInString = append(r.DataInString, string(element))
	}

	return nil
}

func (r Reader) PrintFileContent() {
	fmt.Println("File content is:")
	fmt.Println(r.DataInString)
}

func (r Reader) FindString(string_to_find *string) string {

	return "empty"
}

func Hello() string {
	return "Hello, world"
}
