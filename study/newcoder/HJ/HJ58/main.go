package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
  scan := bufio.NewScanner(os.Stdin)
  for {
    scan.Scan()
    input1 := scan.Text()
    if input1 == "" {
      break
    }

    scan.Scan()
    input2 := scan.Text()
    input1_split := strings.Split(input1," ")
    if len(input1_split) < 2 {
      continue
    }
    input2_split := strings.Split(strings.Trim(input2," ")," ")
    n,_ := strconv.Atoi(input1_split[0])
    k,_ := strconv.Atoi(input1_split[1])
    if n == 0 || k == 0 {
      fmt.Println(0)
    }
  }
}
