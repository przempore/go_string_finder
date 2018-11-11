package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func usage() string {
	path := os.Args[0]
	file := filepath.Base(path)
	return fmt.Sprintf("%s -path=<path/to/file> -find=<string to find>", file)
}

const unknown = "unknown"

const (
	success = iota
	to_few_arguments = iota
	path_not_specified = iota
	find_not_specified = iota
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage())
		os.Exit(to_few_arguments)
	}

	path := flag.String("path", unknown, "a file path")
	string_to_find := flag.String("find", unknown, "string which you want to find")
	flag.Parse()

	switch {
	case *path == unknown:
		fmt.Println("You have to specify flag -path")
		os.Exit(path_not_specified)
	case *string_to_find == unknown:
		fmt.Println("You have to specify flag -find")
		os.Exit(find_not_specified)
	}

	data, err := GetFileData(path)
	result := FindString(data, string_to_find)
	if err != nil {
		os.Exit(5)
	}

	PrintFileContent(data)
	fmt.Printf("'%s' found in:\n", *string_to_find)
	fmt.Println(result)
}
