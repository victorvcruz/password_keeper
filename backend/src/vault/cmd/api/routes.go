package api

import (
	"fmt"
	"github.com/victorvcruz/password_warehouse/protobuf/vault_pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"vault.com/cmd/api/handlers"
)

func New(vault *handlers.VaultHandler, folder *handlers.FolderHandler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	if err != nil {
		return err
	}

	app := grpc.NewServer()
	vault_pb.RegisterVaultServer(app, vault)
	vault_pb.RegisterFolderServer(app, folder)

	err = app.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
