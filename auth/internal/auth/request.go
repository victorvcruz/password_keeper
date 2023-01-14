package auth

type LoginRequest struct {
	Email    string
	Password string
}

type AuthTokenRequest struct {
	AcessToken string
}
