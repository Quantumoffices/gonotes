package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

//加密后:rLyZug0MCEF2TBcJdhMyjg==
func TestAes(t *testing.T) {
	var aeskey = []byte("321423u9y8d2fwfl")
	fmt.Println(len(aeskey))
	pass := []byte("hello,ase")
	//pass2 := []byte("hello,ase2fafafaa")
	//xpass2 := []byte("hello,asefjkajkjfkaljklJLKfafafffffadfaffffff")
	xpass, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%s\n", pass64)
	//
	//bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//	return
	//	//}
	//bytesPass = append(bytesPass, byte(52))

	tpass, err := AesDecrypt([]byte(pass64), aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)

}
