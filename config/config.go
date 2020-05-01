package config

import (
	"net/http"

	"github.com/neelchoudhary/budgetmanagergrpc/utils"
	"github.com/plaid/plaid-go/plaid"
)

// PlaidConfig plaid configuration struct
type PlaidConfig struct {
	ClientID  string
	Secret    string
	PublicKey string
}

// NewPlaidConfig creates new PlaidConfig
func NewPlaidConfig(clientID string, secret string, publicKey string) *PlaidConfig {
	return &PlaidConfig{ClientID: clientID, Secret: secret, PublicKey: publicKey}
}

// ConnectToPlaid establishes a connection with Plaid
func ConnectToPlaid(plaidConfig *PlaidConfig) *plaid.Client {
	plaidClient, err := plaid.NewClient(plaid.ClientOptions{
		ClientID:    plaidConfig.ClientID,
		Secret:      plaidConfig.Secret,
		PublicKey:   plaidConfig.PublicKey,
		Environment: plaid.Development,
		HTTPClient:  &http.Client{},
	})
	utils.LogIfFatalAndExit(err, "Failed to start plaid client: ")

	return plaidClient
}

// DBConfig postgresql database configuration struct
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// NewDBConfig creates new DBConfig
func NewDBConfig(host string, port int, user string, password string, dbname string) *DBConfig {
	return &DBConfig{Host: host, Port: port, User: user, Password: password, Dbname: dbname}
}

// ServerConfig server configuration struct
type ServerConfig struct {
	Environment       string
	Host              string
	APIServerPort     string
	WebhookServerPort string
	TLSKeyPath        string
	TLSCertPath       string
}

// NewServerConfig creates new ServerConfig
func NewServerConfig(environment string, host string, apiServerPort string, webhookServerPort string,
	tlsKeyPath string, tlsCertPath string) *ServerConfig {
	return &ServerConfig{Environment: environment, Host: host, APIServerPort: apiServerPort,
		WebhookServerPort: webhookServerPort, TLSKeyPath: tlsKeyPath, TLSCertPath: tlsCertPath}
}
