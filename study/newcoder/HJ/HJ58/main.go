package main

import "bufio"

func main()  {
  scan := bufio.NewScanner()
  for {
    scan.Scan()
    input1 := scan.Text()
    if input1 == "" {
      break
    }

    scan.Scan()
    input2 := scan.Scan()

  }
}
