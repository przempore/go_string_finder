package main

import (
	"fmt"
	"io/ioutil"
)

func GetFileData(file_path *string) ([]byte, error) {
	data, err := ioutil.ReadFile(*file_path)
	if err != nil {
		fmt.Println("Can't read file:", file_path)
		panic(err)
	}
	return data, nil
}

func PrintFileContent(data []byte) {
	fmt.Println("File content is:")
	fmt.Println(string(data))
}

func FindString(bytes []byte, string_to_find *string) string {

	return "empty"
}


func Hello() string {
	return "Hello, world"
}
