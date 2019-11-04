package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	allIndex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allIndex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchIndex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchIndex)

	submatchAll := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchAll)

	submatchAllIndex := re2.FindAllSubmatchIndex([]byte(a), - 1)
	fmt.Println(submatchAllIndex)
}
