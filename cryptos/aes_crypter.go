package cryptos

/// AES encryption and decryption
/// according to aes' rules, the length of the key is in this range [16, 24, 32]

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"sync"
)

type blockInfo struct {
	errMsg    string
	success   bool
	blockSize int
	decryptor cipher.BlockMode
	encryptor cipher.BlockMode
}

var (
	keyMap = make(map[string]blockInfo, 5)
	lock   = sync.Mutex{}
)

// AesEncrypt using aes to encrypt the given data string
// param: origDataStr the original data to be encrypted
// param: keyStr the key used to encrypted
func AesEncrypt(origDataStr, keyStr string) ([]byte, error) {
	key, origData := []byte(keyStr), []byte(origDataStr)
	block := getBlock(key)
	if !block.success {
		return nil, errors.New(block.errMsg)
	}
	origData = _PKCS5Padding(origData, block.blockSize)
	encrypted := make([]byte, len(origData))
	block.encryptor.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

// AesDecrypt decrypt the given str using given key
// param: encryptedStr  the encrypted result
// param: keyStr 		the key used to encrypt
func AesDecrypt(encryptedStr, keyStr string) ([]byte, error) {
	key, encrypted := []byte(keyStr), []byte(encryptedStr)
	block := getBlock(key)
	if !block.success {
		return nil, errors.New(block.errMsg)
	}
	origData := make([]byte, len(encrypted))
	block.decryptor.CryptBlocks(origData, encrypted)
	origData = _PKCS5UnPadding(origData)
	return origData, nil
}

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func _PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// using get block speed up the encryption and decryption
func getBlock(key []byte) blockInfo {
	if v, ok := keyMap[string(key)]; ok {
		return v
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return blockInfo{
			errMsg:  err.Error(),
			success: false,
		}
	}
	blockSize := block.BlockSize()
	decryptor := cipher.NewCBCDecrypter(block, key[:blockSize])
	encryptor := cipher.NewCBCEncrypter(block, key[:blockSize])
	result := blockInfo{
		errMsg:    "",
		success:   true,
		blockSize: blockSize,
		decryptor: decryptor,
		encryptor: encryptor,
	}
	lock.Lock()
	keyMap[string(key)] = result
	lock.Unlock()
	return result
}
