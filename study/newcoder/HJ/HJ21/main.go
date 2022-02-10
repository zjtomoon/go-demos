package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//简单密码
func main() {
	var str []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str = append(str, input.Text())
	}
}

func deCrypt(str string) {
	dic := map[string]string{
		"1":    "1",
		"0":    "0",
		"abc":  "2",
		"def":  "3",
		"ghi":  "4",
		"jkl":  "5",
		"mno":  "6",
		"pqrs": "7",
		"wxyz": "9",
	}

	for i := 0; i < len(str); i++ {
		switch {
		case str[i] >= 'A' && str[i] <= 'Z':
			if str[i] == 'Z' {
				fmt.Println("a")
			} else {
				fmt.Print(strings.ToLower(string(str[i] + 1)))
			}
		case str[i] >= 'a' && str[i] <= 'z':
			for k, v := range dic {
				if strings.Contains(k, string(str[i])) {
					fmt.Print(v)
				}
			}
		default:
			fmt.Print(string(str[i]))
		}
	}
	fmt.Println()
}
