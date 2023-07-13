package auth

type LoginRequest struct {
	Email    string
	Password string
}

type LoginServiceRequest struct {
	Service     string
	ServiceConn string
	ApiToken    string
}

type AuthTokenRequest struct {
	AcessToken string
}

type AuthTokenService struct {
	Service     string
	ServiceConn string
	AcessToken  string
}

type Register struct {
	Service string
}
