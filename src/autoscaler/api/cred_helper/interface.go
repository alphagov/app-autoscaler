package cred_helper

import (
	"autoscaler/models"
)

const (
	MaxRetry = 5
)

type Credentials interface {
	Create(appId string, userProvidedCredential *models.Credential) (*models.Credential, error)
	Delete(appId string) error
	Get(appId string) (*models.Credential, error)
	// FIXME
	// InitializeConfig(config config.) error

}
