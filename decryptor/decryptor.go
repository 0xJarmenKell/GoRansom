package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	key := []byte("thisisaverysecure32bytekeyphrase") // 32 bytes for AES-256
	directory := "C:\\Users\\user1"

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			decryptFile(path, key)
		}
		return nil
	})
}

func decryptFile(filename string, key []byte) {
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(ciphertext) < aes.BlockSize {
		fmt.Println("Ciphertext too short")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	err = os.WriteFile(filename, ciphertext, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
