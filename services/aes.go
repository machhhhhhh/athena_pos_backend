package services

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type ObjectAES struct {
	UserID int `json:"user_id"`
}

func AESDecrypted(encrypted string) (*ObjectAES, error) {

	var object ObjectAES

	cipher_text, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return &object, err
	}

	block, err := aes.NewCipher([]byte(os.Getenv("ATHENA_AES_KEY")))

	if err != nil {
		return &object, err
	}

	if len(cipher_text)%aes.BlockSize != 0 {
		return &object, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(os.Getenv("ATHENA_AES_IV")))
	mode.CryptBlocks(cipher_text, cipher_text)
	cipher_text = PKCS5UnPadding(cipher_text)

	// Deserialize the JSON
	err = json.Unmarshal(cipher_text, &object)
	if err != nil {
		return &object, err
	}

	return &object, nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

// GetAESEncrypted encrypts gAES_SECRET_IVen text in AES 256 CBC
func AESEncrypted(object *ObjectAES) (string, error) {
	// Serialize the object to JSON
	plain_text, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	// Pad the plain_text to a multiple of 16 bytes
	length := len(plain_text)
	var plain_text_block []byte

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plain_text_block = make([]byte, length+extendBlock)
		copy(plain_text_block[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plain_text_block = make([]byte, length)
	}

	copy(plain_text_block, plain_text)

	// Create a new AES cipher with the provided key
	key := os.Getenv("ATHENA_AES_KEY")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new CBC encrypter with the provided IV
	iv := os.Getenv("ATHENA_AES_IV")
	ciphertext := make([]byte, len(plain_text_block))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plain_text_block)

	// Base64 encode the ciphertext
	str := base64.StdEncoding.EncodeToString(ciphertext)
	return str, nil
}
