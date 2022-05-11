package utils

//#cgo CFLAGS: -I./lib
//#cgo LDFLAGS: -L./lib -lsm4
//#include "sm4.h"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/tjfoc/gmsm/sm4"
)

// todo: cgo版本sm4加解密
//cgo版本的sm4加密
func Sm4Encode_cgo(plainText []byte, mkey []byte) []byte {
	length := len(plainText)
	en_output := make([]byte, 16)
	C.SM4_Encrypt((*C.uchar)(unsafe.Pointer(&mkey[0])), (*C.uchar)(unsafe.Pointer(&plainText[0])), (*C.uchar)(unsafe.Pointer(&en_output[0])))
	return en_output[0:length]
}

//cgo版本的sm4解密
func Sm4Decode_cgo(cipherText []byte, mkey []byte) []byte {
	length := len(cipherText)
	de_output := make([]byte, 16)
	C.SM4_Decrypt((*C.uchar)(unsafe.Pointer(&mkey[0])), (*C.uchar)(unsafe.Pointer(&cipherText[0])), (*C.uchar)(unsafe.Pointer(&de_output[0])))
	return de_output[0:length]
}

// todo: 原生go版本sm4加解密
// 原生go版本的sm4加密
func Sm4Encode(plainText []byte, mkey []byte) []byte {
	en_output, err := sm4.Sm4Ecb(mkey, plainText, true)
	if err != nil {
		fmt.Errorf("sm4 enc error: %s", err)
	}
	return en_output
}

// 原生go版本的sm4解密
func Sm4Decode(cipherText []byte, mkey []byte) []byte {
	de_output, err := sm4.Sm4Ecb(mkey, cipherText, false)
	if err != nil {
		fmt.Errorf("sm4 dec error: %s", err)
	}
	return de_output
}
