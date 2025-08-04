package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log/slog"
	"time"
)

type AuthConfig struct {
	JWTPrivateKey string `env:"JWT_PRIVATE_KEY,required"`
	JWTPublicKey  string `env:"JWT_PUBLIC_KEY,required"`
	TokenExpiry   string `env:"JWT_EXPIRATION" envDefault:"24h"`
	Issuer        string `env:"JWT_ISSUER" envDefault:"agent"`

	JWTPrivateKeyPEM *rsa.PrivateKey
	JWTPublicKeyPEM  *rsa.PublicKey
}

func (c AuthConfig) TokenExpiryDuration() time.Duration {
	dur, err := time.ParseDuration(c.TokenExpiry)
	if err != nil {
		slog.Error("Invalid JWT_EXPIRATION format, using default 24h", "error", err)
		dur = 24 * time.Hour
	}
	return dur
}

func (c *AuthConfig) LoadJWT() error {

	// Parse private key
	block, _ := pem.Decode([]byte(c.JWTPrivateKey))
	if block == nil {
		return fmt.Errorf("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}
	c.JWTPrivateKeyPEM = privateKey

	// Parse public key
	block, _ = pem.Decode([]byte(c.JWTPublicKey))
	if block == nil {
		return fmt.Errorf("failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %v", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("not an RSA public key")
	}
	c.JWTPublicKeyPEM = publicKey
	return nil
}
