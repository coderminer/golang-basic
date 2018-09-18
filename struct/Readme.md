### Golang Struct 声明和使用

`Go`可以声明自定义的数据类型，组合一个或多个类型，可以包含内置类型和用户自定义的类型，可以像内置类型一样使用`struct`类型

#### Struct 声明

具体的语法  

```
type identifier struct{
    field1 data_type
    field2 data_type
    field3 data_type
}
```

例子  

```
package main

import (
	"fmt"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	fmt.Println(rectangle{10.4, 25.10, "red"})
}

```

#### struct 实例化的方法

1. 点运算符

可以使用点运算符访问结构体中的数据值  

```
type rectangle struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

func main() {
	var rec rectangle
	rec.breadth = 19
	rec.length = 23
	rec.color = "Green"

	rec.geometry.area = rec.length * rec.breadth
	rec.geometry.perimeter = 2 * (rec.length + rec.breadth)
	fmt.Println(rec)
	fmt.Println("Area:\t", rec.geometry.area)
	fmt.Println("Perimeter:", rec.geometry.perimeter)
}
```

2. 使用 `var`关键词和 `:=`运算符

如果初始化时，指定了特定的名称，那么有些字段是可以省略的

```
type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	var rect1 = rectangle{10, 20, "Red"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 10, color: "Red"}
	fmt.Println(rect2)

	rect3 := rectangle{10, 20, "Green"}
	fmt.Println(rect3)

	rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

	rect5 := rectangle{breadth: 20, color: "Green"}
	fmt.Println(rect5)
}
```

3. 使用 `new` 关键字

```
type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	rect1 := new(rectangle)
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	rect2 := new(rectangle)
	rect2.breadth = 20
	rect2.color = "Red"
	fmt.Println(rect2)
}
```

4. 使用 `&` 运算符

```
type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	var rect1 = &rectangle{10, 20, "Red"} //此时不能省略任何值
	fmt.Println(rect1)

	var rect2 = &rectangle{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)

	var rect3 = &rectangle{}
	(*rect3).breadth = 20
	(*rect3).color = "Blue"
	fmt.Println(rect3)
}
```

#### struct 练习

1. `struct`中的`tag`标签

```
type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func main() {
	json_str := `
	{
		"firstname":"Kevin",
		"lastname":"Woo",
		"city":"Beijing"
	}`

	emp1 := new(Employee)
	err := json.Unmarshal([]byte(json_str), emp1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emp1)

	emp2 := new(Employee)
	emp2.FirstName = "John"
	emp2.LastName = "Lee"
	emp2.City = "Shanghai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s", jsonStr)
}
```
输出结果  

```
&{Kevin Woo Beijing}
{"firstname":"John","lastname":"Lee","city":"Shanghai"}
```

2. 内嵌的 `struct` 类型

```
func main() {
	type Salary struct {
		Basic, HRA, TA float64
	}
	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}

	e := Employee{
		FirstName: "Kevin",
		LastName:  "Woo",
		Email:     "test@mail.com",
		Age:       12,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.0,
				TA:    2000.0,
			},
			Salary{
				Basic: 16000.0,
				HRA:   6000.0,
				TA:    2100.0,
			},
		},
	}

	fmt.Println(e.FirstName,e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
}
```

输出结果

```
Kevin Woo
12
test@mail.com
{15000 5000 2000}
{16000 6000 2100}
```

3. 为 `struct` 添加方法

```
type Salary struct {
	Basic, HRA, TA float64
}
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("=================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "-------------------"
}

func main() {

	e := Employee{
		FirstName: "Kevin",
		LastName:  "Woo",
		Email:     "test@mail.com",
		Age:       12,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.0,
				TA:    2000.0,
			},
			Salary{
				Basic: 16000.0,
				HRA:   6000.0,
				TA:    2100.0,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}

```

输出结果

```
Kevin Woo
12
test@mail.com
=================
15000
5000
2000
=================
16000
6000
2100
-------------------
```