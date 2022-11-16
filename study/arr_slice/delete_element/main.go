package main

import "fmt"



func main()  {
	SliceDelete()
}

// 删除数组中的某一元素
func SliceDelete() {
	// 初始化一个新的切片 seq
	seq := []string{"a", "b", "c", "d", "e", "f", "g"}

	// 指定删除位置
	index := 3

	// 输出删除位置之前和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// seq[index+1:]... 表示将后段的整个添加到前段中
	// 将删除前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	// 输出链接后的切片
	fmt.Println(seq)
}