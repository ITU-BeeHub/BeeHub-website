package auth

import (
	"errors"
	"os"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// AdminService, admin işlemleriyle ilgili iş mantığını içerir
type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

// Claims, JWT token'ı için gerekli olan bilgileri tutar.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// AuthenticateAdmin, kullanıcı adı ve şifreyi doğrular ve JWT token döner
func (s *AdminService) AuthenticateAdmin(username, password string) (string, error) {
	var admin models.Admin

	// Veritabanından admin kullanıcısını bul
	if err := s.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return "", errors.New("invalid username or password")
	}

	// Şifreyi kontrol et
	if !CheckPasswordHash(password, admin.Password) {
		return "", errors.New("invalid username or password")
	}

	// JWT token oluştur
	token, err := generateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GenerateJWT, admin kullanıcısı için bir JWT token oluşturur
func generateJWT(username string) (string, error) {
	expirataionTime := time.Now().Add(6 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirataionTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken, verilen token'ı doğrular ve JWT claim'lerini döner
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// HashPassword, şifreyi bcrypt ile hashler
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash, verilen şifre ile veritabanındaki hash'i karşılaştırır
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
