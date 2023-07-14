package authorization

import (
	"auth.com/internal/auth"
	"log"
)

type AuthorizationClient interface {
	Login(service string) (string, error)
}

type Authorization struct {
	serviceName string
	client      auth.AuthServiceClient
	apiToken    string
}

func NewAuthorization(service string, auth auth.AuthServiceClient) AuthorizationClient {
	a := &Authorization{
		serviceName: service,
		client:      auth,
	}
	a.apiToken = a.registerService()
	return a
}

func (a *Authorization) Login(service string) (string, error) {
	auth, err := a.client.LoginService(&auth.LoginServiceRequest{Service: a.serviceName, ServiceConn: service, ApiToken: a.apiToken})
	if err != nil {
		return "", err
	}
	return auth, err
}

func (a *Authorization) registerService() string {
	auth, err := a.client.RegisterService(&auth.Register{Service: a.serviceName})
	if err != nil {
		log.Fatalf("failed to register authorization service:%s", err.Error())
	}
	return auth
}
