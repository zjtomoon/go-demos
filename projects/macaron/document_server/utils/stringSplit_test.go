package utils

import (
	"fmt"
	"testing"
)

func TestSubString(t *testing.T) {
	source := "login:szf+123"
	fmt.Println(len(source))
	str1 := SubString(source, 6, len(source))
	fmt.Println(str1)
}

func TestFindindexofColon(t *testing.T) {
	source := "login:szf+123"
	index := FindindexofColon(source)
	fmt.Println(index)
}

func TestFindindexofPlus(t *testing.T) {
	source := "login:szf+123"
	index := FindindexofPlus(source)
	fmt.Println(index)
}

func TestFindindexofSymbol(t *testing.T) {
	source := "login:szf+123"
	index1 := FindindexofSymbol(source, ':')
	fmt.Println(index1)
	index2 := FindindexofSymbol(source, '+')
	fmt.Println(index2)
}
