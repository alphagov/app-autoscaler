package cred_helper

import (
	"autoscaler/db"
	"autoscaler/helpers"
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "CREDENTIALS_PLUGIN",
	MagicCookieValue: "somerandomstring",
}

func LoadCredentialPlugin(dbConfig map[string]db.DatabaseConfig, loggingConfig helpers.LoggingConfig) (Credentials, error) {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "Plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("/Users/kevincross/SAPDevelop/cf/app-autoscaler/src/autoscaler/build/cred_helper"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		logger.Error("failed to create rpcClient", err)
		return nil, fmt.Errorf("failed to create rpcClient %w", err)
	}
	// Request the plugin
	raw, err := rpcClient.Dispense("credHelper")
	if err != nil {
		return nil, fmt.Errorf("failed to dispense plugin %w", err)
	}
	// We should have a customMetricsCredHelper now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	credentials := raw.(Credentials)
	err = credentials.InitializeConfig(dbConfig, loggingConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plugin %w", err)
	}
	return credentials, nil
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"credHelper": &CredentialsPlugin{},
}
