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

const (
	unknown                = "unknown"
	success                = iota
	to_few_arguments       = iota
	path_arg_not_specified = iota
	find_arg_not_specified = iota
	cant_read_from_file    = iota
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
		os.Exit(path_arg_not_specified)
	case *string_to_find == unknown:
		fmt.Println("You have to specify flag -find")
		os.Exit(find_arg_not_specified)
	}

	r := Reader{}

	err := r.GetFileData(path)
	if err != nil {
		os.Exit(cant_read_from_file)
	}

	result := r.FindString(string_to_find)
	fmt.Printf("'%s' found in:\n", *string_to_find)
	for _, r := range result {
		fmt.Println(r)
	}

	os.Exit(success)
}
