package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	//sli := make([]int,5,5)
	//fmt.Printf("adress = %p\n",sli)
	//fmt.Println(sli)
	sli := arr[:3] // 切片遵循左闭右开的规则
	fmt.Println("arr = ", arr)
	fmt.Println("sli = ", sli)
	fmt.Println("length of sli = ", len(sli))
	fmt.Println("cap of sli = ", cap(sli))

	// 扩容未超过cap
	sli = append(sli, 7, 8, 9)
	fmt.Println("arr = ", arr)
	//fmt.Printf("address of arr = %p\n",arr)
	fmt.Println("sli = ", sli)
	fmt.Printf("address of sli = %p\n", sli)

	//扩容超过cap
	sli = append(sli, 10)
	fmt.Println("arr = ", arr)
	fmt.Println("sli = ", sli)
	fmt.Printf("address of sli = %p\n", sli)
}
