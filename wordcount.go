package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	sentense := strings.ReplaceAll(s, ",", "")
	sentense = strings.ReplaceAll(sentense, ".", "")

	words := strings.Fields(sentense)
	result := make(map[string]int)
	
	for _, i := range words {
		result[i] = result[i] + 1
	}
	
	return result
}


func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}