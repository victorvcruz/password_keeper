package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/victorvcruz/password_warehouse/src/protobuf/vault_pb"
	"vault.com/internal/vault"
)

type VaultHandler struct {
	vault_pb.UnimplementedVaultServer
	service vault.ServiceClient
}

func NewVaultHandler(_authService vault.ServiceClient) *VaultHandler {
	return &VaultHandler{
		service: _authService,
	}
}

func (a *VaultHandler) Find(context.Context, *empty.Empty) (*vault_pb.VaultResponse, error) {
	return nil, nil
}

func (a *VaultHandler) FindAll(context.Context, *empty.Empty) (*vault_pb.AllVaultResponse, error) {
	return nil, nil
}

func (a *VaultHandler) Create(context.Context, *vault_pb.VaultRequest) (*vault_pb.VaultResponse, error) {
	return nil, nil
}

func (a *VaultHandler) Update(context.Context, *vault_pb.VaultRequest) (*vault_pb.VaultResponse, error) {
	return nil, nil
}

func (a *VaultHandler) Delete(context.Context, *empty.Empty) (*vault_pb.VaultResponse, error) {
	return nil, nil
}
