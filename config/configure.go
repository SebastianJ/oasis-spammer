package config

import (
	"fmt"
	"path/filepath"

	"github.com/SebastianJ/oasis-spammer/crypto"
	"github.com/SebastianJ/oasis-spammer/genesis"
	"github.com/SebastianJ/oasis-spammer/rpc"
	"github.com/SebastianJ/oasis-spammer/utils"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Configuration - the central configuration for the test suite tool
var Configuration Config

// Configure - configures the test suite tool using a combination of the YAML config file as well as command arguments
func Configure(basePath string, context *cli.Context) (err error) {
	configPath := filepath.Join(basePath, "config.yml")
	if err = loadYamlConfig(configPath); err != nil {
		return err
	}

	if Configuration.BasePath == "" {
		Configuration.BasePath = basePath
	}

	Configuration.Verbose = context.GlobalBool("verbose")

	genesisFile, entityPath, err := setupPaths(context)
	if err != nil {
		return err
	}

	Configuration.GenesisFile = genesisFile
	Configuration.EntityPath = entityPath

	if err := genesis.LoadDocument(Configuration.GenesisFile); err != nil {
		return err
	}

	if ctxAmount := context.GlobalString("amount"); ctxAmount != "" && ctxAmount != Configuration.Transactions.Amount {
		Configuration.Transactions.Amount = ctxAmount
	}

	if ctxSocket := context.GlobalString("socket"); ctxSocket != "" && ctxSocket != Configuration.Socket {
		Configuration.Socket = ctxSocket
	}

	signer, err := crypto.LoadSigner(Configuration.EntityPath)
	if err != nil {
		return err
	}
	Configuration.Signer = signer

	if ctxNonce := context.GlobalInt("nonce"); ctxNonce != Configuration.Transactions.NonceValue {
		Configuration.Transactions.NonceValue = ctxNonce
	}

	if Configuration.Transactions.NonceValue < 0 {
		if nonce, err := rpc.CurrentNonce(Configuration.Signer, Configuration.Socket); err != nil {
			Configuration.Transactions.Nonce = nonce
		}
	} else {
		Configuration.Transactions.Nonce = uint64(Configuration.Transactions.NonceValue)
	}

	if ctxGasFee := context.GlobalString("gas.fee"); ctxGasFee != "" && ctxGasFee != Configuration.Transactions.Gas.Fee {
		Configuration.Transactions.Gas.Fee = ctxGasFee
	}

	if ctxGasLimit := context.GlobalUint64("gas.limit"); ctxGasLimit != 0 && ctxGasLimit != Configuration.Transactions.Gas.Limit {
		Configuration.Transactions.Gas.Limit = ctxGasLimit
	}

	if ctxCount := context.GlobalInt("count"); ctxCount >= 0 && ctxCount != Configuration.Transactions.Count {
		Configuration.Transactions.Count = ctxCount
	}

	if ctxPoolSize := context.GlobalInt("pool-size"); ctxPoolSize >= 0 && ctxPoolSize != Configuration.Transactions.PoolSize {
		Configuration.Transactions.PoolSize = ctxPoolSize
	}

	receiversPath := filepath.Join(Configuration.BasePath, "data/receivers.txt")
	receivers, _ := utils.FetchReceivers(receiversPath)

	if len(receivers) == 0 {
		return fmt.Errorf("you need to create the file %s and add at least one receiver address to it", receiversPath)
	}

	Configuration.Transactions.Receivers = receivers

	return nil
}

func setupPaths(context *cli.Context) (string, string, error) {
	genesisFile, err := filepath.Abs(context.GlobalString("genesis-file"))
	if err != nil {
		return "", "", err
	}

	entityPath, err := filepath.Abs(context.GlobalString("entity-path"))
	if err != nil {
		return "", "", err
	}

	return genesisFile, entityPath, nil
}

func loadYamlConfig(path string) error {
	Configuration = Config{}

	yamlData, err := utils.ReadFileToString(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(yamlData), &Configuration)

	if err != nil {
		return err
	}

	return nil
}
