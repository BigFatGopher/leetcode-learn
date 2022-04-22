package main

import "fmt"

func main() {
	s1 := "how du you do it a b d d d d d"
	s2 := string.split(s1, " ")
	m := make(map[string]int, 10)
	for _, k := range s2 {
		fmt.Printf(k)
	}
}
