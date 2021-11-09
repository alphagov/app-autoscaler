package cred_helper

import (
	"autoscaler/models"
	"net/rpc"
)

// Here is an implementation that talks over RPC
type CredentialsRPC struct {
	client *rpc.Client
}

func (g *CredentialsRPC) Create(args CreateArgs, reply *models.Credential) error {
	err := g.client.Call("Plugin.Create", args, &reply)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return nil
}

func (g *CredentialsRPC) Delete(appId string, reply *interface{}) error {
	err := g.client.Call("Plugin.Delete", appId, &reply)
	if err != nil {
		return err
	}

	return nil
}

func (g *CredentialsRPC) Get(appId string, reply *models.Credential) error {
	err := g.client.Call("Plugin.Get", appId, &reply)
	if err != nil {
		return err
	}

	return nil
}

func (g *CredentialsRPC) InitializeConfig(args InitializeConfigArgs, reply *interface{}) error {
	err := g.client.Call("Plugin.InitializeConfig", args, &reply)
	if err != nil {
		return err
	}

	return nil
}

var _ Credentials = &CredentialsRPC{}

type CredentialsRPCServer struct {
	// This is the real implementation
	Impl Credentials
}

func (s *CredentialsRPCServer) Create(args CreateArgs, reply *models.Credential) error {
	return s.Impl.Create(args, reply)
}

func (s *CredentialsRPCServer) Delete(appId string, reply *interface{}) error {
	return s.Impl.Delete(appId, reply)
}

func (s *CredentialsRPCServer) Get(appId string, reply *models.Credential) error {
	return s.Impl.Get(appId, reply)
}

func (s *CredentialsRPCServer) InitializeConfig(args InitializeConfigArgs, reply *interface{}) error {
	return s.Impl.InitializeConfig(args, reply)
}

var _ Credentials = &CredentialsRPCServer{}
