package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"encoding/json"
)

var secretKey = []byte("32-character-secret-key-abcdef!!")

func EncryptToken(data any) (string, error) {
	plainText, _ := json.Marshal(data)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		fmt.Println("secret length:", len(secretKey))
		fmt.Println("Cipher", err.Error())
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("NewGCM", err.Error())
		return "", err
	}

	nonce := []byte("123456789012")
	cipherText := aesGCM.Seal(nil, nonce, plainText, nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptToken(encrypted string, out any) error {
	cipherData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := []byte("123456789012")
	plainText, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(plainText, &out)
}
