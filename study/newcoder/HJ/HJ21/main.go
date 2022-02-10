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
	for loop := 0; loop < len(str); loop++ {
		deCrypt(str[loop])
	}
}

func deCrypt(str string) {
	dic := map[string]string{
		"1":    "1",
		"abc":  "2",
		"def":  "3",
		"ghi":  "4",
		"jkl":  "5",
		"mno":  "6",
		"pqrs": "7",
		"tuv":  "8",
		"wxyz": "9",
                "0":    "0",
	}

	for loop := 0; loop < len(str); loop++ {
		switch {
		case str[loop] >= 'A' && str[loop] <= 'Z':
			if str[loop] == 'Z' {
				fmt.Print("a")
			} else {
				fmt.Print(strings.ToLower(string(str[loop] + 1)))
			}
		case str[loop] >= 'a' && str[loop] <= 'z':
			for k, v := range dic {
				if strings.Contains(k, string(str[loop])) {
					fmt.Print(v)
				}
			}
		default:
			fmt.Print(string(str[loop]))
		}
	}
	fmt.Print("\n")
}
