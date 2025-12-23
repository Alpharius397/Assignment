package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
)

const AES_IV_LENGTH int = 16

// Check for valid AES key
func checkAesKey(key []byte) ([]byte, error) {
	keyLength := len(key)

	switch keyLength {
	case 16, 24, 32:
		return key, nil
	default:
		return nil, &InvalidKeyLength{Length: keyLength}
	}
}

// Encrypt src into dst using provided key and iv
func aesEncrypt(dst []byte, src []byte, key string, iv []byte) ([]byte, error) {

	Key, err := checkAesKey([]byte(key))

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(Key))

	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	if (len(src) % aes.BlockSize) != 0 {
		return nil, &InvalidSrcBlock{Src: len(src), Block: aes.BlockSize}
	}

	if len(dst) < len(src) {
		return nil, &InvalidDstBlock{Src: len(src), Dst: len(dst)}
	}

	mode.CryptBlocks(dst, src)

	return dst, nil
}

// Decrypt src into dst using provided key and iv
func aesDecrypt(dst []byte, src []byte, key string, iv []byte) ([]byte, error) {

	Key, err := checkAesKey([]byte(key))

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(Key)

	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	if (len(src) % aes.BlockSize) != 0 {
		return nil, &InvalidSrcBlock{Src: len(src), Block: aes.BlockSize}
	}

	if len(dst) < len(src) {
		return nil, &InvalidDstBlock{Src: len(src), Dst: len(dst)}
	}

	mode.CryptBlocks(dst, src)

	return dst, nil
}

// Pad the data block using PKCS7 padding
func pad(data []byte, blockSize int) []byte {

	paddingLength := blockSize - (len(data) % (blockSize))

	lastSlice := make([]byte, paddingLength)

	for i := range lastSlice {
		lastSlice[i] = byte(paddingLength)
	}

	data = append(data, lastSlice...)
	return data
}

// unpad the data block using PKCS7 padding
func unpad(data []byte, blockSize int) ([]byte, error) {
	dataLength := len(data)

	padLength := int((data)[dataLength-1])

	if padLength < 1 || padLength > blockSize {
		return nil, &InvalidPadding{Length: padLength, BlockSize: blockSize}
	}

	actualPadding := make([]byte, padLength)

	for i := range actualPadding {
		actualPadding[i] = byte(padLength)
	}

	givenPadding := string(data[dataLength-padLength:])
	givenData := data[:(dataLength - padLength)]

	if givenPadding != string(actualPadding) {
		return nil, errors.New("incorrect padding found")
	}

	return givenData, nil
}

// a wrapper on aesEncrypt
func encrypt(data []byte, AES_KEY string) (string, error) {

	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)

	if err != nil {
		return "", err
	}

	src := pad(data, aes.BlockSize)
	dst := make([]byte, len(src))

	dst, err = aesEncrypt(dst, src, AES_KEY, iv)

	if err != nil {
		return "", err
	}

	iv = append(iv, dst...)
	iv = pad(iv, 3)

	return base64.URLEncoding.EncodeToString(iv), nil
}

// a wrapper on aesDecrypt
func decrypt(data string, AES_KEY string) (string, error) {

	encrypted, err := base64.URLEncoding.DecodeString(data)

	if err != nil {
		return "", err
	}

	encrypted, err = unpad(encrypted, 3)

	if err != nil {
		return "", err
	}

	iv := encrypted[:AES_IV_LENGTH]
	src := encrypted[AES_IV_LENGTH:]

	dst := make([]byte, len(src))

	dst, err = aesDecrypt(dst, src, AES_KEY, iv)

	if err != nil {
		return "", err
	}

	dst, err = unpad(dst, aes.BlockSize)

	if err != nil {
		return "", err
	}

	return string(dst), nil
}

// exported AES Encryption function, relies of AES_KEY environment variable
func AesEncrypt(data []byte) (string, error) {

	AES_KEY, ok := os.LookupEnv("AES_KEY")

	if !ok {
		return "", &KeyNotFound{}
	}

	return encrypt(data, AES_KEY)
}

// exported AES Decryption function, relies of AES_KEY environment variable
func AesDecrypt(data string) (string, error) {

	AES_KEY, ok := os.LookupEnv("AES_KEY")

	if !ok {
		return "", &KeyNotFound{}
	}

	return decrypt(data, AES_KEY)
}
