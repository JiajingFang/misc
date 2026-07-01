package main

import (
	"fmt"
	ds "github.com/JiajingFang/misc/data_structure"
)

func main() {
	// init hashmap
	s := 5
	hm := ds.NewHashMap(&s)
	hm.Set("name", "Bob")
	hm.Set("age","10")
	hm.Set("location", "Berlin")

	fmt.Printf("value of key \"name\": %s\n", *hm.Get("name"))
	fmt.Printf("value of key \"age\": %s\n", *hm.Get("age"))
	fmt.Printf("value of key \"location\": %s\n", *hm.Get("location"))
	value := hm.Get("random")
	if value == nil {
		fmt.Println("value of key \"random\": nil")
	} else {
		fmt.Println("fail: no existing key has value")
	}
}