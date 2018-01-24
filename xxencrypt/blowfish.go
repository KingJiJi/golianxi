package blowfish

import (
	"encoding/base64"

	"golang.org/x/crypto/blowfish"

	"dbms/lib/blowfish/ecb"
	"dbms/lib/blowfish/padding"
)

func Encrypt(text string, key, salt []byte) (string, error) {

	pt := []byte(text)
	//block, err := blowfish.NewCipher(key)
	block, err := blowfish.NewSaltedCipher(key, salt)
	if err != nil {
		return "", err
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt) // padd last block of plaintext if block size less than block cipher size
	if err != nil {
		return "", err
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return base64.StdEncoding.EncodeToString(ct), nil
}

func Decrypt(text string, key, salt []byte) (string, error) {

	ct, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	//block, err := blowfish.NewCipher(key)
	block, err := blowfish.NewSaltedCipher(key, salt)
	if err != nil {
		return "", err
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		return "", err
	}
	return string(pt), nil
}
