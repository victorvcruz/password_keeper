package model

import (
	"auth.com/internal/auth"
	"encoding/json"
	"github.com/victorvcruz/password_warehouse/protobuf/auth_pb"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func LoginFromProto(proto *auth_pb.LoginRequest) *auth.Request {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var login auth.Request
	err = json.Unmarshal(bytes, &login)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &login
}

func AuthFromProto(proto *auth_pb.AuthTokenRequest) *auth.TokenRequest {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var auth auth.TokenRequest
	err = json.Unmarshal(bytes, &auth)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &auth
}
