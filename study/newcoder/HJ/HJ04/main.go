package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		result := input.Text()

		for len(result) > 8 {
			fmt.Println(result[:8])
			result = result[8:]
		}
		if len(result) == 0 {
			continue
		}
		//fmt.Println(result)
		fmt.Print(result)
		for i := 0; i < 8-len(result); i++ {
			fmt.Print("0")
		}
		fmt.Println()
	}
}
