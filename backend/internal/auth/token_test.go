package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var testSecret = []byte("0123456789abcdef0123456789abcdef")

func TestTokenRoundTrip(t *testing.T) {
	m := NewTokenManager(testSecret, time.Hour)
	token, expiresAt, err := m.Issue("admin")
	if err != nil {
		t.Fatalf("Issue: %v", err)
	}
	if time.Until(expiresAt) <= 0 {
		t.Errorf("expiresAt %v is not in the future", expiresAt)
	}

	username, err := m.Verify(token)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if username != "admin" {
		t.Errorf("username = %q, want admin", username)
	}
}

func TestVerifyRejects(t *testing.T) {
	m := NewTokenManager(testSecret, time.Hour)
	valid, _, _ := m.Issue("admin")

	expired := NewTokenManager(testSecret, time.Hour)
	expired.now = func() time.Time { return time.Now().Add(-2 * time.Hour) }
	expiredToken, _, _ := expired.Issue("admin")

	otherKey, _, _ := NewTokenManager([]byte("another-secret-another-secret-xx"), time.Hour).Issue("admin")

	noneToken, _ := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   "admin",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}).SignedString(jwt.UnsafeAllowNoneSignatureType)

	tests := map[string]string{
		"tampered":    valid + "x",
		"expired":     expiredToken,
		"other key":   otherKey,
		"alg none":    noneToken,
		"empty":       "",
		"not a token": "abc.def.ghi",
	}
	for name, token := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := m.Verify(token); err == nil {
				t.Error("Verify accepted an invalid token")
			}
		})
	}
}
