package cred_helper

import (
	"autoscaler/db"
	"autoscaler/helpers"
	"autoscaler/models"
	"net/rpc"
)

var _ Credentials = &CredentialsClient{}

type CredentialsClient struct {
	client *rpc.Client
}

func (g *CredentialsClient) Create(appId string, userProvidedCredential *models.Credential) (*models.Credential, error) {
	args := CreateArgs{AppId: appId, UserProvidedCredential: userProvidedCredential}
	var reply CreateResponse
	err := g.client.Call("Plugin.Create", args, &reply)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return reply.creds, reply.err
}

func (g *CredentialsClient) Delete(appId string) error {
	var reply *interface{}
	err := g.client.Call("Plugin.Delete", appId, &reply)
	if err != nil {
		return err
	}
	return nil
}

func (g *CredentialsClient) Get(appId string) (*models.Credential, error) {
	var reply *models.Credential
	err := g.client.Call("Plugin.Get", appId, &reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (g *CredentialsClient) InitializeConfig(dbConfig map[string]db.DatabaseConfig, loggingConfig helpers.LoggingConfig) error {
	args := InitializeConfigArgs{DbConfig: dbConfig, LoggingConfig: loggingConfig}
	var reply *interface{}
	err := g.client.Call("Plugin.InitializeConfig", args, &reply)
	if err != nil {
		return err
	}
	return nil
}

type CredentialsRPCServer struct {
	Impl Credentials
}

func (s *CredentialsRPCServer) Create(args CreateArgs, reply *CreateResponse) error {
	creds, err := s.Impl.Create(args.AppId, args.UserProvidedCredential)
	if err != nil {
		return err
	}
	*reply = CreateResponse{creds: creds, err: err}
	return nil
}

func (s *CredentialsRPCServer) Delete(appId string, reply *error) error {
	*reply = s.Impl.Delete(appId)
	return nil
}

func (s *CredentialsRPCServer) Get(appId string, reply *GetResponse) error {
	creds, err := s.Impl.Get(appId)
	if err != nil {
		return err
	}
	*reply = GetResponse{creds: creds, err: err}
	return nil
}

func (s *CredentialsRPCServer) InitializeConfig(args InitializeConfigArgs, reply *error) error {
	*reply = s.Impl.InitializeConfig(args.DbConfig, args.LoggingConfig)
	return nil
}

var _ CredentialsServer = &CredentialsRPCServer{}
