#### Data Type

`Go`是一种静态的编程语言，`Go`中的变量必须要有特定的类型并且不能改变，使用关键词`var`用来定义特定的数据类型，具体的语法是  

```
var name type = expression
```

使用关键字 `var` 声明一个变量，并且赋值给它。也可以声明多个相同类型的变量，具体的语法

```
var fname,lname string
```

在函数体内也可以使用短声明来定义变量，具体语法 

```
country,state := "Chinese","Berlin"
```

```
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

```