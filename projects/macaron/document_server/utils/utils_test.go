package utils

import (
	"fmt"
	"testing"
)

func TestGenerateIDS(t *testing.T) {
	num := GenerateIDS()
	fmt.Println(num)
}
