package ethereum

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func GenerateKeys() (string, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return "", "", err
	}

	publicKey := privateKey.PublicKey

	privateKeyBytes := privateKey.D.Bytes()
	publicKeyBytes := elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y)

	return hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes), nil
}

func GetAddress(publicKeyHex string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)

	if err != nil {
		return "", err
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	address := hash.Sum(nil)[12:]

	return hex.EncodeToString(address), nil
}
