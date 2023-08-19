package model

import (
	"github.com/victorvcruz/password_warehouse/src/protobuf/vault_pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"vault.com/internal/vault"
)

func VaultResponseToProto(resp vault.Response) *vault_pb.VaultResponse {
	return &vault_pb.VaultResponse{
		Id:        int64(resp.ID),
		UserId:    int64(resp.UserID),
		FolderId:  int64(resp.FolderID),
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

func ProtoToVaultRequest(req *vault_pb.VaultRequest, userId uint) vault.Request {
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
