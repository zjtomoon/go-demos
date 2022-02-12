package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		scan.Scan()
		input1 := scan.Text()
		if input1 == "" {
			break
		}

		scan.Scan()
		input2 := scan.Text()
		input1_split := strings.Split(input1, " ")
		if len(input1_split) < 2 {
			continue
		}
		input2_split := strings.Split(strings.Trim(input2, " "), " ")
		n, _ := strconv.Atoi(input1_split[0])
		k, _ := strconv.Atoi(input1_split[1])
		if n == 0 || k == 0 {
			fmt.Println(0)
		}
		var arr []int
		for i := 0; i < len(input2_split); i++ {
			tmp, _ := strconv.Atoi(input2_split[i])
			arr = append(arr, tmp)
		}
		if n != len(arr) || k > n {
			break
		}
		sort.Ints(arr)
		//fmt.Println(arr[:k])
		//整型数组转为字符串数组并打印输出
		var arr2 []string
		for i := 0; i < len(arr); i++ {
			tmp := strconv.Itoa(arr[i])
			arr2 = append(arr2, tmp)
		}
		fmt.Println(strings.Join(arr2[:k], " "))
	}
}
