package main

import (
	"fmt"
	"golang/builtin"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)

	r, closer, err := builtin.BuildZipReader("./assets/my_data.txt.gz")
	if err != nil {
		panic(err)
	}
	defer closer()
	counts, err := builtin.CountLetters(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(counts)
}
