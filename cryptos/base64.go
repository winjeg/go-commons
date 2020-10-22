package cryptos

import "encoding/base64"

/// standard base64 encoding
/// using aes as the basic algorithm for crypting
func Base64Encrypt(data, key string) (string, error) {
	d, err := AesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(d), nil
}

/// standard base64 decoding
/// using aes as basic algorithm
func Base64Decrypt(data, key string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	d, decErr := AesDecrypt(string(decodeBytes), key)
	if decErr != nil {
		return "", decErr
	}
	return string(d), nil
}
