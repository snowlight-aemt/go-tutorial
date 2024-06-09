package main

import (
	"fmt"
	"github.com/snowlight-aemt/type-dict/mydict"
	"log"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("first", "nick")
	if err != nil {
		log.Fatalln(err)
	}

	def, err2 := dictionary.Search("first")
	if err2 != nil {
		log.Fatalln(err)
	}
	fmt.Println(def)
}
