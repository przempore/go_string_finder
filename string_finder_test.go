package main

import (
	"errors"
	"testing"
)

func Test_Hello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func Test_ToFewArguments(t *testing.T) {
	var args []string
	args = append(args, "dupa")
	_, got_err := GetFileData(args)
	want_err := errors.New("To few arguments.")

	if got_err != want_err {
		t.Errorf("got '%s' want '%s'", got_err, want_err)
	}
}

func Test_FindString(t *testing.T) {

}
