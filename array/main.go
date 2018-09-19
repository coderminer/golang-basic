package main

import "fmt"

func main() {
	var intArray = [5]int{0: 10, 2: 30, 4: 50}
	fmt.Println(intArray)

	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	for index, item := range intArray {
		fmt.Println(index, item)
	}

	strArr1 := [3]string{"Beijing", "Shanghai", "Guangzhou"}
	fmt.Printf("strArr1: %v \n", strArr1)
	strArr2 := strArr1
	fmt.Printf("strArr2: %v\n", strArr2)
	strArr1[0] = "Hangzhou"
	fmt.Printf("strArr1: %v\n", strArr1)
	fmt.Printf("strArr2: %v\n", strArr2)

	strArr3 := &strArr1
	fmt.Printf("strArr3: %v\n", strArr3)
	fmt.Printf("&strArr3: %v\n", &strArr3)
	fmt.Printf("*strArr3: %v\n", *strArr3)
}
