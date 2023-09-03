package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/victorvcruz/password_warehouse/protobuf/vault_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"vault.com/cmd/api/model"
	"vault.com/internal/auth"
	"vault.com/internal/folder"
	"vault.com/internal/utils"
	"vault.com/internal/utils/errors"
)

type FolderHandler struct {
	vault_pb.UnimplementedFolderServer
	folderService folder.ServiceClient
	authService   auth.AuthServiceClient
}

func NewFolderHandler(folderService folder.ServiceClient,
	authService auth.AuthServiceClient) *FolderHandler {
	return &FolderHandler{
		folderService: folderService,
		authService:   authService,
	}
}

func (v *FolderHandler) Find(ctx context.Context, _ *empty.Empty) (*vault_pb.FolderResponse, error) {
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

func (v *FolderHandler) FindAll(ctx context.Context, _ *empty.Empty) (*vault_pb.AllFolderResponse, error) {
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

func (v *FolderHandler) Create(ctx context.Context, req *vault_pb.FolderRequest) (*vault_pb.FolderResponse, error) {
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

func (v *FolderHandler) Update(ctx context.Context, req *vault_pb.FolderRequest) (*vault_pb.FolderResponse, error) {
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

func (v *FolderHandler) Delete(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
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
