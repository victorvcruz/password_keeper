package handlers

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/victorvcruz/password_warehouse/src/protobuf/vault_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"vault.com/cmd/api/model"
	"vault.com/internal/auth"
	"vault.com/internal/folder"
	"vault.com/internal/utils"
	"vault.com/internal/utils/errors"
	"vault.com/internal/vault"
)

type VaultHandler struct {
	vault_pb.UnimplementedVaultServer
	vaultService  vault.ServiceClient
	folderService folder.ServiceClient
	authService   auth.AuthServiceClient
}

func NewVaultHandler(vaultService vault.ServiceClient, folderService folder.ServiceClient,
	authService auth.AuthServiceClient) *VaultHandler {
	return &VaultHandler{
		vaultService:  vaultService,
		folderService: folderService,
		authService:   authService,
	}
}

func (a *VaultHandler) Find(ctx context.Context, _ *empty.Empty) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	_, err = a.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vaultId, err := strconv.ParseUint(utils.GetMetadata(ctx, "vault_id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	vault, err := a.vaultService.Find(uint(vaultId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}
	return model.VaultResponseToProto(vault), nil
}

func (a *VaultHandler) FindAll(ctx context.Context, _ *empty.Empty) (*vault_pb.AllVaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := a.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vaults, err := a.vaultService.FindAll(uint(id))
	if err != nil {
		switch err.(type) {
		case *errors.NotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	respb := make([]*vault_pb.VaultResponse, len(vaults))
	for i := range vaults {
		respb[i] = model.VaultResponseToProto(vaults[i])
	}
	return &vault_pb.AllVaultResponse{VaultResponse: respb}, nil
}

func (a *VaultHandler) Create(ctx context.Context, req *vault_pb.VaultRequest) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := a.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vault, err := a.vaultService.Create(model.ProtoToVaultRequest(req, uint(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return model.VaultResponseToProto(vault), nil
}

func (a *VaultHandler) Update(ctx context.Context, req *vault_pb.VaultRequest) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := a.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vault, err := a.vaultService.Update(model.ProtoToVaultRequest(req, uint(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return model.VaultResponseToProto(vault), nil
}

func (a *VaultHandler) Delete(ctx context.Context, _ *empty.Empty) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	_, err = a.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vaultId, err := strconv.ParseUint(utils.GetMetadata(ctx, "vault_id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	err = a.vaultService.Delete(uint(vaultId))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return nil, nil
}
