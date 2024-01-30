package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"abc", "xyx"}
	fmt.Println(sort.SearchStrings(a, "xyz"))
}
