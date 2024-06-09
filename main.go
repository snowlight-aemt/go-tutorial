package main

import (
	"fmt"
	"github.com/snowlight-aemt/type-dict/mydict"
	"log"
)

func main() {
	dictionary := mydict.Dictionary{"first": "niko"}
	value, err := dictionary.Search("second")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(value)
}
