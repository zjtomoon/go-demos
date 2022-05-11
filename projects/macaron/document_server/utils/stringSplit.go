package utils

// todo: 字符串截取
// 参考: https://blog.csdn.net/psyuhen/article/details/51998223
func SubString(source string, start int, end int) string {
	var str = []rune(source)
	length := len(str)

	if start < 0 || end > length || start > end {
		return ""
	}
	if start == 0 && end == length {
		return source
	}
	return string(str[start:end])
}

func FindindexofColon(source string) int {
	var index int
	for i := 0; i < len(source); i++ {
		if source[i] == ':' {
			index = i
		}
	}
	return index
}

func FindindexofPlus(source string) int {
	var index int
	for i := 0; i < len(source); i++ {
		if source[i] == '+' {
			index = i
		}
	}
	return index
}

func FindindexofSymbol(source string, Symbol byte) int {
	var index int
	for i := 0; i < len(source); i++ {
		if source[i] == Symbol {
			index = i
		}
	}
	return index
}
