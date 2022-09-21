package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

const (
	characterBytes  = "!@#$%^&*?"
	digitBytes      = "1234567890"
	lowLetterBytes  = "abcdefghijklmnopqrstuvwxyz"
	highLetterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits   = 6                    // 6bits to represent a leter index
	letterIdxMask   = 1<<letterIdxBits - 1 // All 1-bits,as many as letterIdxBits
	letterIdMax     = 63 / letterIdxBits   // of letter indices fitting in 63 bits
)

// NewPwd 复杂密码生成器
func NewPwd(length int) (string, error) {
	defer mainMenu()
	//check length
	if length < 4 || length > 25 {
		return "", errors.New("length is invalid")
	}
	// assign elements from 4 kinds of base elements
	r := rand.New(rand.NewSource(time.Now().Unix()))
	factor := length / 4
	b := make([]byte, length)
	b = assignElement(lowLetterBytes, length, 1, b)
	b = assignElement(highLetterBytes, length, factor*2, b)
	b = assignElement(characterBytes, length, factor*3, b)
	b = assignElement(digitBytes, length, factor*4, b)

	// shuffle
	rand.Shuffle(len(b), func(i, j int) {
		i = r.Intn(length)
		b[i], b[j] = b[j], b[i]
	})
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
	isInvalid := Judge(*(*string)(unsafe.Pointer(&b)))
	fmt.Println("是否满足复杂度", isInvalid)
	return *(*string)(unsafe.Pointer(&b)), nil
}

// assignMent 为密码分配元素
func assignElement(base string, length int, factor int, b []byte) []byte {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, cacheValue, remain := length-factor, r.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cacheValue, remain = r.Int63(), letterIdMax
		}
		if idx := int(cacheValue & letterIdMax); idx < len(base) {
			b[i] = base[idx]
			i--
		}
		cacheValue >>= letterIdxBits
		remain--
	}
	return b
}

// SimpleNewPwd 复杂密码生产，简易实现方式
func SimpleNewPwd(length int) (pwd string, err error) {
	defer mainMenu()
	if length < 4 || length > 25 {
		return "", errors.New("length is invalid")
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	factor := length / 4
	characters := []rune(characterBytes)
	digits := []rune(digitBytes)
	highLetters := []rune(highLetterBytes)
	lowLetters := []rune(lowLetterBytes)

	// get rand elements
	b := make([]rune, length)
	for i := range b[:factor] {
		b[i] = highLetters[r.Intn(len(highLetters))]
	}
	for i := range b[factor : factor*2] {
		b[i+factor] = lowLetters[r.Intn(len(characters))]
	}
	for i := range b[factor*2 : factor*3] {
		b[i+factor*2] = characters[r.Intn(len(characters))]
	}
	for i := range b[factor*3:] {
		b[i+factor*3] = digits[r.Intn(len(digits))]
	}

	// shuffle
	rand.Shuffle(len(b), func(i, j int) {
		i = r.Intn(length)
		b[i], b[j] = b[j], b[i]
	})
	fmt.Println(string(b))
	isInvalid := Judge(string(b))
	fmt.Println("是否满足复杂度", isInvalid)
	return string(b), nil
}

// MiddleNewPwd 复杂密码生成，中等实现方式
func MiddleNewPwd(length int) (pwd string, err error) {
	defer mainMenu()
	if length < 4 || length > 25 {
		return "", errors.New("length is invalid")
	}
	var pwdBase string
	factor := length / 4
	lastFactor := length - 3*factor
	characters := getRandStr(characterBytes, factor)
	pwdBase += characters
	digits := getRandStr(digitBytes, factor)
	pwdBase += digits
	highLetters := getRandStr(highLetterBytes, factor)
	pwdBase += highLetters
	lowLetters := getRandStr(lowLetterBytes, lastFactor)
	pwdBase += lowLetters

	//shuffle
	var temp []string
	for _, s := range pwdBase {
		temp = append(temp, string(s))
	}
	pwd, err = shuffle(temp)
	if err != nil {
		return
	}
	fmt.Println(pwd)
	isValid := Judge(pwd)
	fmt.Println("是否满足复杂度", isValid)
	return
}

// shuffle 打乱元素
func shuffle(slice []string) (str string, err error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	length := len(slice)
	// judge length
	if length < 1 {
		return "", errors.New("length is invalid")
	}
	// shuffle
	for i := 0; i < length; i++ {
		randIndex := r.Intn(length) // 随机数
		slice[length-1], slice[randIndex] = slice[randIndex], slice[length-1]
	}
	str = strings.Join(slice, "")
	return
}

// getRandstr 获取随机字符
func getRandStr(baseStr string, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}

// Judge 密码复杂度判断
func Judge(pwd string) bool {
	if len(pwd) < 8 {
		return false
	}
	// 检查字符串元素复杂度
	var flag []int
	for i := 0; i < len(pwd); i++ {
		if unicode.IsLower(rune(pwd[i])) {
			flag = append(flag, 1)
		} else if unicode.IsDigit(rune(pwd[i])) {
			flag = append(flag, 2)
		} else if strings.Contains(characterBytes, string(pwd[i])) {
			flag = append(flag, 3)
		} else if unicode.IsUpper(rune(pwd[i])) {
			flag = append(flag, 4)
		}
	}
	// 复杂度标记切片去重
	complexity := len(removeRepeatedElement(flag[:]))
	if complexity >= 3 {
		return true
	} else {
		return false
	}
}

// removeRepeatedElement 数组去重 通过map键的唯一性去重
func removeRepeatedElement(s []int) []int {
	result := make([]int, 0)
	m := make(map[int]bool) // map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

func exit() {
	// 保持窗口等待用户操作
	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func mainMenu() {
	var choice int
	fmt.Println("\n***\t主菜单\t\n***\t1 生成复杂密码 快速版本\t\n***\t2 生成复杂密码 简单版本\t\n***\t3 生成复杂密码 中等版本\t\n***\t4 退出\t\n***\t请输入菜单:\t")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		_, _ = NewPwd(8)
	case 2:
		_, _ = SimpleNewPwd(8)
	case 3:
		_, _ = MiddleNewPwd(8)
	case 4:
		exit()
	default:
		fmt.Println("抱歉，输入有误")
		mainMenu()
	}
}

func main() {
	mainMenu()
}
