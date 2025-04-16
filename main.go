package main

import (
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	var (
		appID      = flag.Int64("app-id", 0, "GitHub App ID")
		privateKey = flag.String("private-key", "", "Path to private key file")
		expiration = flag.Duration("expiration", 10*time.Minute, "Token expiration duration")
	)
	flag.Parse()

	if *appID == 0 || *privateKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	token, err := generateToken(*appID, *privateKey, *expiration)
	if err != nil {
		log.Fatalf("Error generating token: %v", err)
	}

	fmt.Println(token)
}

func generateToken(appID int64, privateKeyPath string, expiration time.Duration) (string, error) {
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to read private key: %w", err)
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	now := time.Now()
	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expiration)),
		Issuer:    fmt.Sprintf("%d", appID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}
