### Array

在`Go`中数组就是包含一种固定类型的数据的集合，数组的长度是固定的不能改变

#### 数组声明和实例化

1. Array的基本的声明的语法是 

```
var name [length] element_type
```
```
var intArr [5]int
```

2. 声明时初始化

```
var name [length] element_type {comma-separated list of element values}
```
如：

```
var intArray = [5]int{10,20,30}
```

可以利用索引初始化元素

```
var intArray = [5]int{0:10,2:30,4:50}
```

3. 在函数中可以使用 `:=` 进行初始化

```
intArray := [5]int{10, 20, 30, 40, 50}
```

4. 也可以使用 `...`初始化

```
intArray := [...]int{10, 20, 30, 40, 50}
```

数组的长度会根据初始化的元素的数量决定

#### 练习

1. 遍历数组

```
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}
```

```
	for index, item := range intArray {
		fmt.Println(index, item)
	}
```

2. 通过值或引用给数组赋值

```
	strArr1 := [3]string{"Beijing", "Shanghai", "Guangzhou"}
	fmt.Printf("strArr1: %v \n", strArr1)
	strArr2 := strArr1 // 数组值copy
	fmt.Printf("strArr2: %v\n", strArr2)
	strArr1[0] = "Hangzhou" //strArr1中值的改变，不会影响strArr2
	fmt.Printf("strArr1: %v\n", strArr1)
	fmt.Printf("strArr2: %v\n", strArr2)

	strArr3 := &strArr1
	fmt.Printf("strArr3: %v\n", strArr3)
	fmt.Printf("&strArr3: %v\n", &strArr3)
	fmt.Printf("*strArr3: %v\n", *strArr3)
```