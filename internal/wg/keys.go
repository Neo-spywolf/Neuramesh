package wg

import (
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/curve25519"
    "os"
)

// GenerateKeys creates a WireGuard-style X25519 key pair
func GenerateKeys() (string, string, error) {
    var privateKey [32]byte
    _, err := rand.Read(privateKey[:])
    if err != nil {
        return "", "", err
    }

    publicKey, err := curve25519.X25519(privateKey[:], curve25519.Basepoint)
    if err != nil {
        return "", "", err
    }

    return base64.StdEncoding.EncodeToString(privateKey[:]),
           base64.StdEncoding.EncodeToString(publicKey), nil
}

// SaveToFile stores the keys in a simple config file
func SaveToFile(priv, pub string, filename string) error {
    data := []byte("[keys]\nprivate=" + priv + "\npublic=" + pub + "\n")
    return os.WriteFile(filename, data, 0600)
}
