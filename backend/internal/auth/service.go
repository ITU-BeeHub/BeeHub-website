package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

// dummyHash is compared against when the username does not exist, so a
// login attempt takes the same time whether or not the user exists.
var dummyHash, _ = bcrypt.GenerateFromPassword([]byte("dummy-password"), bcrypt.DefaultCost)

// Session is an authenticated admin session.
type Session struct {
	Token     string
	ExpiresAt time.Time
}

// Service authenticates admins against the database.
type Service struct {
	db     *gorm.DB
	tokens *TokenManager
}

func NewService(db *gorm.DB, tokens *TokenManager) *Service {
	return &Service{db: db, tokens: tokens}
}

// Login verifies the credentials and opens a session.
func (s *Service) Login(ctx context.Context, username, password string) (Session, error) {
	var admin Admin
	err := s.db.WithContext(ctx).Where("username = ?", username).First(&admin).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		_ = bcrypt.CompareHashAndPassword(dummyHash, []byte(password))
		return Session{}, ErrInvalidCredentials
	case err != nil:
		return Session{}, fmt.Errorf("find admin: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)) != nil {
		return Session{}, ErrInvalidCredentials
	}

	token, expiresAt, err := s.tokens.Issue(admin.Username)
	if err != nil {
		return Session{}, err
	}
	return Session{Token: token, ExpiresAt: expiresAt}, nil
}

// SetPassword creates the admin or replaces its password.
func (s *Service) SetPassword(ctx context.Context, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	admin := Admin{Username: username}
	return s.db.WithContext(ctx).
		Where(Admin{Username: username}).
		Assign(Admin{PasswordHash: string(hash)}).
		FirstOrCreate(&admin).Error
}
