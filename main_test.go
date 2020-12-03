package main

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	t.Fatal("_________________________error_____________________________________")
}
func TestMain2(t *testing.T) {
	fmt.Println("::error is ignore from github log:: _________________________error_____________________________________")
	t.Fatal("::is ignore from github log:: _________________________error_____________________________________")
}
func TestMain3(t *testing.T) {
	t.Fatal("_________________________error_____________________________________")
}
func TestMain4(t *testing.T) {
	t.Fatal("_________________________error_____________________________________")
}
