package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//HJ18 识别有效的IP地址和掩码并进行分类统计
func main() {
	var errNum int

	m := make(map[string]int)
	m = map[string]int{
		"a":   0,
		"b":   0,
		"c":   0,
		"d":   0,
		"e":   0,
		"pri": 0,
	}
	// 初始化一个对象
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// 从输入读取
		scanner.Scan()
		var str = scanner.Text() //fmt.Println("输入的是："+str)
		// 如果读取到空行就退出
		if str == "" {
			break
		}
		ip, mask := ipToInt(str)
		if ip[0] == 0 || ip[0] == 127 {
			continue
		}
		if !ipRight(ip) {
			errNum++
			continue
		}
		if !ipMask(mask) {
			errNum++
			continue
		}
		ipTypeNum(ip, m)

	}
	fmt.Printf("%d %d %d %d %d %d %d", m["a"], m["b"], m["c"], m["d"], m["e"], errNum, m["pri"])
}

func ipMask(mask []int) bool {
	sum := mask[0]<<24 + mask[1]<<16 + mask[2]<<8 + mask[3]
	if sum == 0xFFFFFFFF || sum == 0 {
		return false
	}
	res := int64(sum ^ 0xFFFFFFFF + 1)
	bitSum := strconv.FormatInt(res, 2)
	if len(bitSum) > 32 {
		return false
	}
	var count int
	for _, v := range []byte(bitSum) {
		if v == '1' {
			count++
		}
	}
	if count == 1 {
		return true
	} else {
		return false
	}
}
func ipRight(ip []int) bool {
	if len(ip) != 4 {
		return false
	} else if ip[0] > 255 || ip[1] > 255 || ip[2] > 255 || ip[3] > 255 {
		return false
	} else {
		return true
	}
}
func ipTypeNum(ip []int, m map[string]int) {
	if ip[0] == 10 || (ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31) || (ip[0] == 192 && ip[1] == 168) {
		m["pri"]++
	}
	if ip[0] >= 1 && ip[0] <= 126 {
		m["a"]++
	}
	if ip[0] >= 128 && ip[0] <= 191 {
		m["b"]++
	}
	if ip[0] >= 192 && ip[0] <= 223 {
		m["c"]++
	}
	if ip[0] >= 224 && ip[0] <= 239 {
		m["d"]++
	}
	if ip[0] >= 240 && ip[0] <= 255 {
		m["e"]++
	}

}

func ipToInt(str string) (ip, mask []int) {
	word := strings.Split(str, "~")
	ipStr := word[0]
	maskStr := word[1]
	ipSection := strings.Split(ipStr, ".")
	maskSection := strings.Split(maskStr, ".")
	ip = section(ipSection)
	mask = section(maskSection)
	return
}

func section(str []string) (ip []int) {
	for i := 0; i < 4; i++ {
		num, err := strconv.Atoi(str[i])
		if err == nil {
			ip = append(ip, num)
		}
	}
	return ip

}
