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
	Create(appId string, userProvidedCredential *models.Credential) (*models.Credential, error)
	Delete(appId string) error
	Get(appId string) (*models.Credential, error)
	InitializeConfig(dbConfig map[string]db.DatabaseConfig, loggingConfig helpers.LoggingConfig) error
}
