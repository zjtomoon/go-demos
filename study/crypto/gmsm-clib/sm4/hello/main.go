package main

//#cgo CFLAGS: -I./lib
//#cgo LDFLAGS: -L./lib -lsm4
//#include "sm4.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	key := []byte("1234567890abcdef")
	fmt.Printf("key = %v\n", key)
	plainText := []byte("hello,sm4")
	fmt.Println(plainText)
	en_output := make([]byte, len(plainText))
	fmt.Println(en_output)
	de_output := make([]byte, len(plainText))
	fmt.Println(de_output)
	C.SM4_Encrypt((*C.uchar)(unsafe.Pointer(&key[0])),
		(*C.uchar)(unsafe.Pointer(&plainText[0])),
		(*C.uchar)(unsafe.Pointer(&en_output[0])))
	fmt.Println(en_output)
	fmt.Println(string(en_output))
	fmt.Println([]byte(en_output))
	C.SM4_Decrypt((*C.uchar)(unsafe.Pointer(&key[0])), (*C.uchar)(unsafe.Pointer(&en_output[0])), (*C.uchar)(unsafe.Pointer(&de_output[0])))
	fmt.Println(de_output)
	fmt.Println(string(de_output))
	fmt.Println([]byte(de_output))
}

// 参考: https://blog.csdn.net/woailp___2005/article/details/106144388
//func main() {
//	ret := 0
//
//	key := []byte("1234567890abcdef")
//	iv := []byte("1234567890abcdef")
//	orig := []byte("hello world!")
//	cipherText := make([]byte,0)
//	plainText := make([]byte,0)
//
//	ret = Sm4CbcEncrypt(orig, cipherText, key, iv)
//	ret = Sm4CbcDecrypt(cipherText[0:ret], plainText, key, iv)
//	fmt.Println(string(plainText[0:ret]))
//}

//func Sm4CbcEncrypt(source, cipherText, key, iv []byte) (int) {
//
//	clen := (C.int)(0)
//
//	C.Sm4CbcEncrypt((*C.uchar)(unsafe.Pointer(&source[0])), (C.int)(len(source)),
//		(*C.uchar)(unsafe.Pointer(&cipherText[0])), &clen,
//		(*C.uchar)(unsafe.Pointer(&key[0])),
//		(*C.uchar)(unsafe.Pointer(&iv[0])));
//	return int(clen)
//}
//
//func Sm4CbcDecrypt(source, plainText, key, iv []byte) (int) {
//
//	clen := (C.int)(0)
//
//	C.Sm4CbcDecrypt((*C.uchar)(unsafe.Pointer(&source[0])), (C.int)(len(source)),
//		(*C.uchar)(unsafe.Pointer(&plainText[0])), &clen,
//		(*C.uchar)(unsafe.Pointer(&key[0])),
//		(*C.uchar)(unsafe.Pointer(&iv[0])));
//	return int(clen)
//}
