package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var x int
var y int

func main() {
	x, y = 0, 0
	reader := bufio.NewReader(os.Stdin)
	data,_,_ := reader.ReadLine()
	array := strings.Split(string(data),";")
	length := len(array)
	for i := 0 ; i < length; i++ {
		itemLen := len(array[i])
		if itemLen < 2 || itemLen > 3 {
			continue
		}
		ch := array[i][0] // 字符串也是数组
		str := array[i][1:]
		span,err := strconv.Atoi(str)
		if err == nil {
			move(ch,span)
		}
	}
	fmt.Printf("%d,%d",x,y)
}

func move(ch byte, span int) {
	switch ch {
	case 'A':
		x -= span

	case 'D':
		x += span

	case 'W':
		y += span

	case 'S':
		y -= span
	default:
	}
}