package main

import (
	"github.com/snowlight-aemt/type-dict/mydict"
	"log"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("first", "nick")
	if err != nil {
		log.Fatalln(err)
	}

	err2 := dictionary.Update("first", "TEST")
	if err2 != nil {
		log.Fatalln(err)
	}

	_, err3 := dictionary.Search("first")
	if err3 != nil {
		log.Fatalln(err)
	}

	err4 := dictionary.Delete("first")
	if err4 != nil {
		log.Fatalln(err)
	}
}
