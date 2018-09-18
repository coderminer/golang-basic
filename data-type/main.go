package main

import (
	"fmt"
)

// 全局变量声明
var (
	m int
	n int
)

func main() {
	var x int = 1 // 声明并赋值
	var y int     // 声明，默认是0值
	fmt.Println(x, y)

	var a, b, c = 5.25, 23.3, 3.4
	fmt.Println(a, b, c)

	city := "Berlin"
	country := "Germany"
	fmt.Println(country, city)

	food, drink, price := "Pizza", "Pepsi", 234
	m, n = 1, 2
	fmt.Println(m, n)
	fmt.Println(food, drink, price)
}
