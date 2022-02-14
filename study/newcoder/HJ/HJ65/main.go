package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var params []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		params = append(params, input.Text())
	}
	for i := 0; i < len(params); i += 2 {
		res := longest_sameStr(params[i], params[i+1])
		fmt.Println(res)
	}
}

func longest_sameStr(str1, str2 string) string {
	var a, b string
	if len(str1) <= len(str2) {
		a, b = str1, str2
	} else {
		a, b = str2, str1
	}
	var start1, start2, max = 0, 0, 0
	var tmpStr string
	for i := 0; i < len(a); i++ {
		tmpStr = string(a[start1:i] + string(a[i]))
		for !strings.Contains(b, tmpStr) && start1 < i {
			start1++
			tmpStr = string(a[start1:i] + string(a[i]))
		}
		if strings.Contains(b, tmpStr) && max < i-start1+1 {
			max = i - start1 + 1
			start2 = start1
		}
	}
	return a[start2 : start2+max]
}
