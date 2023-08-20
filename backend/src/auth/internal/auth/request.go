package auth

type Request struct {
	Email    string
	Password string
}

type ServiceRequest struct {
	Service     string
	ServiceConn string
	ApiToken    string
}

type TokenRequest struct {
	AcessToken string
}

type TokenService struct {
	Service     string
	ServiceConn string
	AcessToken  string
}

type Register struct {
	Service string
}
