package service

type TokenRequest struct {
	Service  string `json:"service"`
	Password string `json:"password"`
}