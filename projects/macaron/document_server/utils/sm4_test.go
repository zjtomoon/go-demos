package utils

import (
	"fmt"
	"testing"
)

func TestSm4Encode(t *testing.T) {
	mkey := []byte("1234567890abcdef")
	plainText := []byte("hello,sm4!!!!!!!!!!!!!!!!")
	en_output := Sm4Encode(plainText, mkey)
	fmt.Println(en_output)
	de_output := Sm4Decode(en_output, mkey)
	fmt.Println(string(de_output))
}

// todo :cgo版本加密后的[]byte与go版本不一致
func TestSm4Encode_cgo(t *testing.T) {
	mkey := []byte("1234567890abcdef")
	plainText := []byte("hello,sm4")
	en_output := Sm4Encode_cgo(plainText, mkey)
	fmt.Println(en_output)
	de_output := Sm4Decode_cgo(en_output, mkey)
	fmt.Println(string(de_output))
}
