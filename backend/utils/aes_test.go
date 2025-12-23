package utils

import (
	"crypto/rand"
	"errors"
	"testing"
)

// Panics as we need the keys
func generateRandomKey(size int) string {
	key := make([]byte, size)
	_, err := rand.Read(key)

	if err != nil {
		panic(err)
	}

	return string(key)
}

func encrypt_and_Decrypt(t *testing.T, VALID_AES_KEY string) {

	data := "Hello World"

	value, err := encrypt([]byte(data), VALID_AES_KEY)

	if err != nil {
		t.Error(err)
		return
	}

	finalValue, err := decrypt(value, VALID_AES_KEY)

	if err != nil {
		t.Error(err)
	} else if finalValue != data {
		t.Errorf("Decrypted value doesn't match actual value. Expected: %s, Got: %s", data, finalValue)
	}
}

func failed_Encrypt_and_Decrypt(t *testing.T, INVALID_AES_KEY string) {

	data := "Hello World"

	_, err := encrypt([]byte(data), INVALID_AES_KEY)

	var expectedError *InvalidKeyLength

	if err == nil {
		t.Error("Encryption passed without valid keys")
		return
	} else if !errors.As(err, &expectedError) {
		t.Errorf("Invalid Error returned. Expected %#v, Got: %#v", expectedError, err)
		return
	}

}

func Test_Encrypt_Decrypt(t *testing.T) {
	encrypt_and_Decrypt(t, generateRandomKey(16))
	encrypt_and_Decrypt(t, generateRandomKey(24))
	encrypt_and_Decrypt(t, generateRandomKey(32))
}

func Test_Failed_Encrypt_Decrypt(t *testing.T) {
	failed_Encrypt_and_Decrypt(t, generateRandomKey(11))
	failed_Encrypt_and_Decrypt(t, generateRandomKey(22))
	failed_Encrypt_and_Decrypt(t, generateRandomKey(31))
}
