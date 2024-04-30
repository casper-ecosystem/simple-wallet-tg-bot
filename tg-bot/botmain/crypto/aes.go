package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
)

const BlockSize = 32

func deriveKey(passphrase string, salt []byte) []byte {
	// http://www.ietf.org/rfc/rfc2898.txt
	if salt == nil {
		salt = make([]byte, 8)
		// rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, BlockSize, sha256.New)
}

// func addBase64Padding(value string) string {
// 	m := len(value) % 4
// 	if m != 0 {
// 		value += strings.Repeat("=", 4-m)
// 	}

// 	return value
// }

// func removeBase64Padding(value string) string {
// 	return strings.Replace(value, "=", "", -1)
// }

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encrypt(key []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	msg := data
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], msg)

	return ciphertext, nil
}

func decrypt(key []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := data[:aes.BlockSize]
	msg := data[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	return msg, nil
}
