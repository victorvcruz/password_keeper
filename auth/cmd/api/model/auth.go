package model

import (
	"auth.com/internal/auth"
	"encoding/json"
	"github.com/victorvcruz/password_warehouse/protobuf/auth_pb"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func LoginFromProto(proto *auth_pb.LoginRequest) *auth.LoginRequest {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var login auth.LoginRequest
	err = json.Unmarshal(bytes, &login)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &login
}

func LoginServiceFromProto(proto *auth_pb.LoginService) *auth.LoginServiceRequest {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var login auth.LoginServiceRequest
	err = json.Unmarshal(bytes, &login)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &login
}

func AuthFromProto(proto *auth_pb.AuthTokenRequest) *auth.AuthTokenRequest {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var auth auth.AuthTokenRequest
	err = json.Unmarshal(bytes, &auth)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &auth
}

func AuthServiceFromProto(proto *auth_pb.AuthTokenService) *auth.AuthTokenService {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var auth auth.AuthTokenService
	err = json.Unmarshal(bytes, &auth)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &auth
}

func RegisterFromProto(proto *auth_pb.Register) *auth.Register {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, UseEnumNumbers: false}.Marshal(proto)
	if err != nil {
		log.Printf("error when marshal to json %s", err.Error())
		return nil
	}

	var auth auth.Register
	err = json.Unmarshal(bytes, &auth)
	if err != nil {
		log.Printf("error when unmarshal to session %s", err.Error())
		return nil
	}

	return &auth
}
