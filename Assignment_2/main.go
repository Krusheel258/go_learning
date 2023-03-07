package main

import (
	"ass_2/Cache"
	"ass_2/Cache/Lru"
	"fmt"
	"log"
)

func main() {

	l, err := Lru.CreateLruCache(1)
	if err != nil {
		log.Fatal(err)
	}
	var data = []Cache.Data{
		{"1", "Violet"},
		{"2", "Indigo"},
		{"3", "Blue"},
	}
	for _, d := range data {
		l.Put(d)
	}

	val, err := l.Get("2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The value for key : ", val)

}
