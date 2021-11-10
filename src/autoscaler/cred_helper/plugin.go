package cred_helper

import (
	"autoscaler/db"
	"autoscaler/helpers"
	"autoscaler/models"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type CredentialsServer interface {
	Create(args CreateArgs, reply *CreateResponse) error
	Delete(appId string, reply *error) error
	Get(appId string, reply *GetResponse) error
	InitializeConfig(args InitializeConfigArgs, reply *error) error
}

type CreateArgs struct {
	AppId                  string
	UserProvidedCredential *models.Credential
}
type CreateResponse struct {
	creds *models.Credential
	err   error
}
type GetResponse struct {
	creds *models.Credential
	err   error
}

type InitializeConfigArgs struct {
	DbConfig      map[string]db.DatabaseConfig
	LoggingConfig helpers.LoggingConfig
}

type CredentialsPlugin struct {
	Impl Credentials
}

func (p *CredentialsPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &CredentialsRPCServer{Impl: p.Impl}, nil
}

func (CredentialsPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &CredentialsClient{client: c}, nil
}
