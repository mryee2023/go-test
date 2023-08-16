package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func LoadPKCS1PrivateKey(privateKey []byte) (privKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {

		return nil, errors.New("invalid private key ")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		oldErr := err
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("ParsePKCS1PrivateKey error: %s, ParsePKCS8PrivateKey error: %s", oldErr.Error(), err.Error())
		}

		switch t := key.(type) {
		case *rsa.PrivateKey:
			priv = key.(*rsa.PrivateKey)
		default:
			return nil, fmt.Errorf("ParsePKCS1PrivateKey error: %s, ParsePKCS8PrivateKey error: Not supported privatekey format, should be *rsa.PrivateKey, got %T", oldErr.Error(), t)
		}
	}

	return priv, nil
}

func SignPKCS1v15WithSha256(priv *rsa.PrivateKey, plaintext []byte) (sign string, err error) {

	h := sha256.New()
	h.Write(plaintext)
	hashed := h.Sum(nil)

	sig, e := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	if e != nil {
		return "", e
	}

	return base64.StdEncoding.EncodeToString(sig), nil
}
