package model

import (
	"github.com/victorvcruz/password_warehouse/protobuf/vault_pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"vault.com/internal/folder"
	"vault.com/internal/vault"
)

func VaultResponseToProto(resp vault.Response) *vault_pb.VaultResponse {
	return &vault_pb.VaultResponse{
		Id:        resp.ID,
		UserId:    resp.UserID,
		FolderId:  resp.FolderID,
		Username:  resp.Username,
		Name:      resp.Name,
		Password:  resp.Password,
		Url:       resp.URL,
		Notes:     resp.Notes,
		Favorite:  resp.Favorite,
		CreatedAt: timestamppb.New(*resp.CreatedAt),
		UpdatedAt: timestamppb.New(*resp.UpdatedAt),
	}
}

func ProtoToVaultRequest(req *vault_pb.VaultRequest, userId uint64) vault.Request {
	return vault.Request{
		Name:     req.Name,
		UserID:   userId,
		FolderID: req.FolderId,
		Username: req.Username,
		Password: req.Password,
		URL:      req.Url,
		Notes:    req.Notes,
		Favorite: req.Favorite,
	}
}

func FolderResponseToProto(resp folder.Response) *vault_pb.FolderResponse {
	return &vault_pb.FolderResponse{
		Id:        resp.ID,
		UserId:    resp.UserID,
		Name:      resp.Name,
		CreatedAt: timestamppb.New(resp.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UpdatedAt),
	}
}

func ProtoToFolderRequest(req *vault_pb.FolderRequest, userId uint64) folder.Request {
	return folder.Request{
		Name:   req.Name,
		UserID: userId,
	}
}
