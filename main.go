package main

import "fmt"

func main() {
	stepBlock := 21.0 // 每21w个区块奖励减半
	curGet := 50.0    // 初识奖励 50
	total := 0.0
	for curGet > 0 {
		curNum := stepBlock * curGet
		curGet *= 0.5
		total += curNum
	}
	fmt.Println(total)
}
