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
    mask := scan.Text()
    if mask == "" {
      break
    }
    scan.Scan()
    ip1 := scan.Text()
    scan.Scan()
    ip2 := scan.Text()

    if !ipValid(mask) || !ipValid(ip1) || !ipValid(ip2) || !isMask(mask) {
      fmt.Println(1)
      continue
    }
    newIP1 := processIp(ip1)
    newIP2 := processIp(ip2)
    newMask := processIp(mask)
    if newIP1&newMask == newIP2&newMask {
      fmt.Println(0)
      continue
    }
    fmt.Println(2)
  }
}


func ipValid(ip string) bool {
  ipList := strings.Split(ip,".")
  for _,v := range ipList {
    ipN,_ := strconv.Atoi(v)
    if ipN < 0 {
      return false
    }
    ip0,_ := strconv.ParseUint(v,10,32)
    if ip0 > uint64(255) || ip0 < uint64(0) {
      return false
    }
  }
  return true
}


func isMask(mask string) bool {
  ipList := strings.Split(mask,".")
  for k,v := range ipList {
    ipN,_ := strconv.Atoi(v)
    if ipN < 0 {
      return false
    }
    ip0,_ := strconv.ParseUint(v,10,32)
    if k == 0 {
      if ip0 != uint64(255) {
        return false
      }
    } else if k == 3 {
      if ip0 != uint64(0) {
        return false
      }
    }
  }
  return true
}


func processIp(ip string) uint32 {
  ipList := strings.Split(ip,".")
  newIp := uint32(0)
  ip0,_ := strconv.ParseUint(ipList[0],10,32)
  ip1,_ := strconv.ParseUint(ipList[1],10,32)
  ip2,_ := strconv.ParseUint(ipList[2],10,32)
  ip3,_ := strconv.ParseUint(ipList[3],10,32)
  newIp += uint32(ip0) << 24
  newIp += uint32(ip1) << 16
  newIp += uint32(ip2) << 8
  newIp += uint32(ip3)

  return newIp
}
