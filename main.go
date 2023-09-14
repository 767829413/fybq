package main

import "fmt"

func main() {
	var m = make(map[string][]int)

	m["wewewe"] = append(m["wewewe"], 454544545)
	m["wewewe"] = append(m["wewewe"], 343434)

	fmt.Println(m)
}
