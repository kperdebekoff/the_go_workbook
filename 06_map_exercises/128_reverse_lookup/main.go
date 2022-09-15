package main

import "fmt"

func main() {

	var frEn map[string]string = map[string]string{
		"le": "the",
		"la": "the",
		"livre": "book",
		"pomme": "apple",
	}

	fmt.Println("List of keys:", reverseLookup(frEn, "the"))
	fmt.Println("List of keys:", reverseLookup(frEn, "apple"))
	fmt.Println("List of keys:", reverseLookup(frEn, "iphone"))

}

func reverseLookup(dict map[string]string, value string) []string {

	var keys []string
	for k, v := range dict {
		if v == value {
			keys = append(keys, k)
		}
	}
	return keys

}