package handlers

import (
	"context"
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
	vault_pb.UnimplementedFolderServer
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

func (v *VaultHandler) FindVault(ctx context.Context, _ *empty.Empty) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	_, err = v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vaultId, err := strconv.ParseUint(utils.GetMetadata(ctx, "vault_id"), 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "vault_id must be a number")
	}

	vault, err := v.vaultService.Find(vaultId)
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

func (v *VaultHandler) FindAllVaults(ctx context.Context, _ *empty.Empty) (*vault_pb.AllVaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vaults, err := v.vaultService.FindAll(uint64(id))
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

func (v *VaultHandler) CreateVault(ctx context.Context, req *vault_pb.VaultRequest) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vault, err := v.vaultService.Create(model.ProtoToVaultRequest(req, uint64(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return model.VaultResponseToProto(vault), nil
}

func (v *VaultHandler) UpdateVault(ctx context.Context, req *vault_pb.VaultRequest) (*vault_pb.VaultResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vault, err := v.vaultService.Update(model.ProtoToVaultRequest(req, uint64(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return model.VaultResponseToProto(vault), nil
}

func (v *VaultHandler) DeleteVault(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	_, err = v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	vaultId, err := strconv.ParseUint(utils.GetMetadata(ctx, "vault_id"), 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "vault_id must be a number")
	}

	err = v.vaultService.Delete(vaultId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return nil, nil
}

func (v *VaultHandler) FindFolder(ctx context.Context, _ *empty.Empty) (*vault_pb.FolderResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	_, err = v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	folderId, err := strconv.ParseUint(utils.GetMetadata(ctx, "folder_id"), 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "vault_id must be a number")
	}

	folder, err := v.folderService.Find(folderId)
	if err != nil {
		switch err.(type) {
		case *errors.NotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}
	return model.FolderResponseToProto(folder), nil
}

func (v *VaultHandler) FindAllFolders(ctx context.Context, _ *empty.Empty) (*vault_pb.AllFolderResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	folders, err := v.folderService.FindAllByUserId(uint64(id))
	if err != nil {
		switch err.(type) {
		case *errors.NotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	respb := make([]*vault_pb.FolderResponse, len(folders))
	for i := range folders {
		respb[i] = model.FolderResponseToProto(folders[i])
	}
	return &vault_pb.AllFolderResponse{FolderResponse: respb}, nil
}

func (v *VaultHandler) CreateFolder(ctx context.Context, req *vault_pb.FolderRequest) (*vault_pb.FolderResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	folder, err := v.folderService.Create(model.ProtoToFolderRequest(req, uint64(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return model.FolderResponseToProto(folder), nil
}

func (v *VaultHandler) UpdateFolder(ctx context.Context, req *vault_pb.FolderRequest) (*vault_pb.FolderResponse, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	id, err := v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	folder, err := v.folderService.Update(model.ProtoToFolderRequest(req, uint64(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return model.FolderResponseToProto(folder), nil
}

func (v *VaultHandler) DeleteFolder(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	token, err := utils.BearerToken(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Add token")
	}

	_, err = v.authService.AuthUserToken(token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	folderId, err := strconv.ParseUint(utils.GetMetadata(ctx, "folder_id"), 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "folder_id must be a number")
	}

	err = v.folderService.Delete(folderId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return nil, nil
}
