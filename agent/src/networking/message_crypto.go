package networking

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func EncodeMessage(message []byte, public_key_string string) ([]byte, string) {
	decoded_public_key, _ := base64.StdEncoding.DecodeString(public_key_string)
	public_key_block, _ := pem.Decode(decoded_public_key)
	public_key_object, err := x509.ParsePKIXPublicKey(public_key_block.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	public_key_rsa := public_key_object.(*rsa.PublicKey)

	// Generate random AES key
	aes_key := RandomString(16)

	// Encrypt AES key and nonce with RSA pub key
	encrypted_key := EncryptKey(public_key_rsa, aes_key)

	// Encrypt message with AES key
	encrypted_message := EncryptMessage(aes_key, message)
	full_message := append(encrypted_key[:], encrypted_message[:]...)
	encoded_message := base64.StdEncoding.EncodeToString(full_message)
	return []byte(encoded_message), aes_key

}

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

func EncryptMessage(aes_key string, message []byte) []byte {
	block, err := aes.NewCipher([]byte(aes_key))
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(message))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], message)
	return ciphertext
}
