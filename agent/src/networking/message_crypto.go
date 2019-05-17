package networking

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"io"
)

func RandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	random_string := fmt.Sprintf("%X", b)
	return string(random_string)
}

func EncryptKey(pubkey *rsa.PublicKey, message string) []byte {
	message_bytes := []byte(message)
	hash := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pubkey, message_bytes, nil)
	if err != nil {
		fmt.Println(err)
	}
	return ciphertext
}

// The following AES CFB Encrypt/Decrypt functions are interoperable between
// Python and Golang. Thanks to Blixt @ https://stackoverflow.com/a/42770165
func DecryptMessage(key string, ciphertext []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic(err)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}

func EncryptMessage(aes_key string, message string) []byte {
	block, err := aes.NewCipher([]byte(aes_key))
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len([]byte(message)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(message))
	return ciphertext
}
