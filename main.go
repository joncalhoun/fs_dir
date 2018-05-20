package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	ex2()
}

func ex1() {
	d := http.Dir(".")
	f, err := d.Open("stuff/thing.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}

func ex2() {
	d := http.Dir("./stuff")
	f, err := d.Open("thing.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}
