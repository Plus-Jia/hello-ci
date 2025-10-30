package main

import "fmt"

func main() {
	fmt.Println("Hello, GitHub CI!")
	result := Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", result)
}

// Add 函数用于测试
func Add(a, b int) int {
	return a + b
}
