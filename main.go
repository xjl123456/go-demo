package main

//
//import (
//	"bytes"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"encoding/base64"
//	"encoding/pem"
//	"errors"
//	"log"
//)
//
//const (
//	privateKeyStr string = "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAJ27/M0S27kTd/JL6wYOZcK8VmmSpkqVFpcDF+0/381VsMFrdb3N9Jmac7Vs2vuOFrasQ6u5ARJ9X+9AqTlsQgqF5t7fnc6mU6UEQwwGGGfpWQACttfLmQf36zrlgmQvxbAIGRMMQkAsAQ9I1IhceiZV73oDcHV27ihd7HmZqt3fAgMBAAECgYEAmWgqz1yG/DJWFv4FH0fDaqj3tgfd2W20obxted4EkVTE6ujTg30aZoXUAWBUfhHMP2+9BPeSdmQfeLa/nsyOUNx2Lv6rugClItBb554mYn8IjDPYENyLo0vaNDMMQ/KvZ/WR2vnYViwaMn/e4ROipbfiGdlBWn5DEewtsEUoDkkCQQDSB6DixFODU6N0cWxW4aTpX/DcC+84ImM1BEdS9+s6Hfv5Enb+cEWeiVlkKjmP/Ovn3iH8+WqAsmtQh7mn8ZbLAkEAwEIkxgDDM3yYqwTWMtKXMxjiT4Ric1rBJrHze6loG2havAk/BAj6nHBvN3tCmVZx1XIxHwU1JLZwKWjL5tBevQJBALWmBV67H+OALeliw6msxD1XTfBynfX1v8m1pp46b4Y3MpsrfiD3Jy9DaT25S0meHMXQF6M8cAFYznm6uTZoOtsCQCyXKrBBvQRUAZSoqoVfEnJncxW+PpdClUnEPBSSVfMFYQX7nwHwky91ZFYZ4Hhv9DbtJTdsncbGCX2RMLl32oECQQCwadpjKTsLAKhjrsbmtB1J+jjxI+OHGahxfd/1zl7y+BN68zx7ny13T5J845d0zpdW2bUApWq63gCKuFlxv/m5"
//	publicKeyStr  string = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCdu/zNEtu5E3fyS+sGDmXCvFZpkqZKlRaXAxftP9/NVbDBa3W9zfSZmnO1bNr7jha2rEOruQESfV/vQKk5bEIKhebe353OplOlBEMMBhhn6VkAArbXy5kH9+s65YJkL8WwCBkTDEJALAEPSNSIXHomVe96A3B1du4oXex5mard3wIDAQAB"
//)
//
///**
//生成RSA密钥对
//*/
//func GenRSAKey() ([]byte,[]byte,error) {
//	//生成密钥
//	var publicKeyBytes, _ = base64.StdEncoding.DecodeString(publicKeyStr)
//	var privateKeyBytes, _ = base64.StdEncoding.DecodeString(privateKeyStr)
//
//	return privateKeyBytes, publicKeyBytes,nil
//}
//
///**
//公钥加密-分段
//*/
//func RsaEncryptBlock(src, publicKeyByte []byte) (bytesEncrypt []byte, err error) {
//	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyByte)
//	if err != nil {
//		return
//	}
//	keySize, srcSize := publicKey.Size(), len(src)
//	log.Println("密钥长度：", keySize, "\t明文长度：\t", srcSize)
//	//单次加密的长度需要减掉padding的长度，PKCS1为11
//	offSet, once := 0, keySize-11
//	buffer := bytes.Buffer{}
//	for offSet < srcSize {
//		endIndex := offSet + once
//		if endIndex > srcSize {
//			endIndex = srcSize
//		}
//		// 加密一部分
//		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src[offSet:endIndex])
//		if err != nil {
//			return nil, err
//		}
//		buffer.Write(bytesOnce)
//		offSet = endIndex
//	}
//	bytesEncrypt = buffer.Bytes()
//	return
//}
//
//
//
//
//func RsaDecryptBlock(src, privateKeyByte []byte) (bytesDecrypt  []byte, err error) {
//	block, _ := pem.Decode(privateKeyByte)
//	if block == nil {
//		return nil, errors.New("获取私钥失败")
//	}
//	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
//	if err != nil {
//		return
//	}
//	private := privateKey.(*rsa.PrivateKey)
//	keySize, srcSize := private.Size(), len(src)
//	//logs.Debug("密钥长度：", keySize, "\t密文长度：\t", srcSize)
//	var offSet = 0
//	var buffer = bytes.Buffer{}
//	for offSet < srcSize {
//		endIndex := offSet + keySize
//		if endIndex > srcSize {
//			endIndex = srcSize
//		}
//		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, private, src[offSet:endIndex])
//		if err != nil {
//			return nil, err
//		}
//		buffer.Write(bytesOnce)
//		offSet = endIndex
//	}
//	bytesDecrypt = buffer.Bytes()
//	return
//}
//
//func main() {
//	var mingwen = "2w3e$R%T"
//
//	log.Println("明文:--" + mingwen)
//
//	//RSA的内容使用base64打印
//	privateKey, publicKey, _ := GenRSAKey()
//	//log.Println("rsa私钥:\t", base64.StdEncoding.EncodeToString(privateKey))
//	//log.Println("rsa公钥:\t", base64.StdEncoding.EncodeToString(publicKey))
//
//
//	miwen, err := RsaEncryptBlock([]byte(mingwen), publicKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("加密后：\t", base64.StdEncoding.EncodeToString(miwen))
//
//	jiemi, err := RsaDecryptBlock(miwen, privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("解密后：\t", string(jiemi))
//}
