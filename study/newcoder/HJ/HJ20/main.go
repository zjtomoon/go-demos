package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		scan.Scan()
		if scan.Text() == "" {
			break
		}
		a := strings.Trim(scan.Text(), " ")
		if len(a) <= 8 {
			fmt.Println("NG")
			continue
		}
		is := false
		for i := 0; i < len(a)-3; i++ {
			if strings.Count(a, a[i:i+3]) != 1 {
				fmt.Println("NG")
				is = true
				break
			}
		}
		if is {
			continue
		}
		ac := 0
		uac := 0
		nc := 0
		oc := 0
		for i := 0; i < len(a); i++ {
			if a[i] <= '9' && a[i] >= '0' {
				nc = 1
			} else if a[i] <= 'Z' && a[i] >= 'A' {
				uac = 1
			} else if a[i] <= 'z' && a[i] >= 'a' {
				ac = 1
			} else {
				oc = 1
			}
		}
		if ac+uac+nc+oc >= 3 {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}
