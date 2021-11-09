package cred_helper

import (
	"autoscaler/db"
	"autoscaler/helpers"
	"autoscaler/models"
)

const (
	MaxRetry = 5
)

type Credentials interface {
	Create(args CreateArgs, reply *models.Credential) error
	Delete(appId string, reply *interface{}) error
	Get(appId string, reply *models.Credential) error
	InitializeConfig(args InitializeConfigArgs, reply *interface{}) error
}

type CreateArgs struct {
	AppId                  string
	UserProvidedCredential *models.Credential
}

type InitializeConfigArgs struct {
	DbConfig      map[string]db.DatabaseConfig
	LoggingConfig helpers.LoggingConfig
}
