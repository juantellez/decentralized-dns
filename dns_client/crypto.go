package dnsclient

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

// GenerateKeyPair genera un par (pública, privada) base64 codificado
func GenerateKeyPair() (pubKeyBase64, privKeyBase64 string, err error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(priv), nil
}

// SignMessage firma un mensaje con la clave privada base64
func SignMessage(privKeyBase64 string, message []byte) (string, error) {
	privKeyBytes, err := base64.StdEncoding.DecodeString(privKeyBase64)
	if err != nil {
		return "", err
	}
	sig := ed25519.Sign(ed25519.PrivateKey(privKeyBytes), message)
	return base64.StdEncoding.EncodeToString(sig), nil
}

// VerifySignature verifica la firma base64 sobre el mensaje, con la clave pública base64
func VerifySignature(pubKeyBase64 string, message []byte, sigBase64 string) (bool, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		return false, err
	}
	sigBytes, err := base64.StdEncoding.DecodeString(sigBase64)
	if err != nil {
		return false, err
	}
	return ed25519.Verify(ed25519.PublicKey(pubKeyBytes), message, sigBytes), nil
}
