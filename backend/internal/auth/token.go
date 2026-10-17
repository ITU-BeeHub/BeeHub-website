package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const issuer = "beehub-website"

// TokenManager issues and verifies admin session tokens.
type TokenManager struct {
	secret []byte
	ttl    time.Duration
	now    func() time.Time
}

func NewTokenManager(secret []byte, ttl time.Duration) *TokenManager {
	return &TokenManager{secret: secret, ttl: ttl, now: time.Now}
}

// Issue returns a signed token for username and its expiry time.
func (m *TokenManager) Issue(username string) (string, time.Time, error) {
	now := m.now()
	expiresAt := now.Add(m.ttl)
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   username,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign token: %w", err)
	}
	return signed, expiresAt, nil
}

// Verify checks the token's signature, issuer and expiry and returns the
// username it was issued for.
func (m *TokenManager) Verify(token string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims,
		func(*jwt.Token) (any, error) { return m.secret, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(issuer),
		jwt.WithExpirationRequired(),
		jwt.WithTimeFunc(m.now),
	)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}
