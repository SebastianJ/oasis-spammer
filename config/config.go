package config

import "github.com/oasislabs/oasis-core/go/common/crypto/signature"

// Config - represents the config
type Config struct {
	BasePath     string           `yaml:"-"`
	Verbose      bool             `yaml:"-"`
	Signer       signature.Signer `yaml:"-"`
	GenesisFile  string           `yaml:"genesis_file"`
	EntityPath   string           `yaml:"entity_path"`
	Socket       string           `yaml:"socket"`
	Transactions Transactions     `yaml:"transactions"`
}

// Transactions - represents the transactions settings group
type Transactions struct {
	Amount     string `yaml:"amount"`
	Data       string `yaml:"data"`
	Count      int    `yaml:"count"`
	PoolSize   int    `yaml:"pool_size"`
	NonceValue int    `yaml:"nonce"`
	Nonce      uint64 `yaml:"-"`
	Gas        Gas    `yaml:"gas"`
	Receivers  []string
}

// Gas - represents the gas settings
type Gas struct {
	Fee   string `yaml:"fee"`
	Limit uint64 `yaml:"limit"`
}
