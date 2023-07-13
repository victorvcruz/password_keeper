package model

import (
	"auth.com/internal/auth"
	pb "auth.com/pkg/pb"
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func LoginFromProto(proto *pb.LoginRequest) *auth.LoginRequest {
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

func LoginServiceFromProto(proto *pb.LoginService) *auth.LoginServiceRequest {
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

func AuthFromProto(proto *pb.AuthTokenRequest) *auth.AuthTokenRequest {
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

func AuthServiceFromProto(proto *pb.AuthTokenService) *auth.AuthTokenService {
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

func RegisterFromProto(proto *pb.Register) *auth.Register {
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
