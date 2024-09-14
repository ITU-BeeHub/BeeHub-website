package models

// ErrorResponse, bir hata durumunda döndürülen struct'tır.
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse, başarılı bir işlemden sonra döndürülen struct'tır.
type SuccessResponse struct {
	Message string `json:"message"`
}

// TokenResponse, JWT token döndüren struct'tır.
type TokenResponse struct {
	Token string `json:"token"`
}
