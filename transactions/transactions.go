package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/SebastianJ/oasis-spammer/rpc"
	"github.com/SebastianJ/oasis-spammer/utils"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	fileSigner "github.com/oasislabs/oasis-core/go/common/crypto/signature/signers/file"
	"github.com/oasislabs/oasis-core/go/common/entity"
	"github.com/oasislabs/oasis-core/go/consensus/api/transaction"
	staking "github.com/oasislabs/oasis-core/go/staking/api"
)

// Send - send transactions
func Send(signer signature.Signer, to string, amount string, nonce uint64, gasFee string, gasLimit uint64, socket string) error {
	//defer signer.Reset()
	bigAmount, _ := utils.ConvertNumeralStringToBigInt(amount)

	var xfer staking.Transfer
	if err := xfer.To.UnmarshalText([]byte(to)); err != nil {
		fmt.Printf("failed to parse transfer destination ID, err: %s\n", err.Error())
		return err
	}
	if err := xfer.Tokens.FromBigInt(bigAmount); err != nil {
		fmt.Printf("failed to parse transfer amount, err: %s\n", err.Error())
		return err
	}

	var fee transaction.Fee
	if err := fee.Amount.UnmarshalText([]byte(gasFee)); err != nil {
		fmt.Printf("failed to parse fee amount, err: %s\n", err.Error())
		return err
	}
	fee.Gas = transaction.Gas(gasLimit)

	tx := staking.NewTransferTx(nonce, &fee, &xfer)

	fmt.Println("")
	fmt.Printf("tx: %+v\n", tx)
	tx.PrettyPrint("", os.Stdout)

	signedTx, rawSignedTx, err := sign(signer, tx)
	if err != nil {
		return err
	}

	fmt.Printf("Signed tx: %s\n", rawSignedTx)

	_, client, err := rpc.ConsensusClient(socket)
	if err != nil {
		return err
	}

	if err := client.SubmitTx(context.Background(), signedTx); err != nil {
		fmt.Printf("failed to submit transaction, err: %s\n", err.Error())
		return err
	}

	return nil
}

func sign(signer signature.Signer, tx *transaction.Transaction) (*transaction.SignedTransaction, []byte, error) {
	sigTx, err := transaction.Sign(signer, tx)
	if err != nil {
		fmt.Printf("failed to sign transaction, err: %s\n", err.Error())
		return nil, nil, err
	}

	fmt.Println("")
	fmt.Printf("sigTx: %+v\n", sigTx)
	sigTx.PrettyPrint("", os.Stdout)

	rawTx, err := json.Marshal(sigTx)
	if err != nil {
		fmt.Printf("failed to marshal transaction, err: %s\n", err.Error())
		return nil, nil, err
	}

	return sigTx, rawTx, nil
}

// LoadSigner - loads the signer from the PEM file
func LoadSigner(path string) (signature.Signer, error) {
	fmt.Printf("Path is: %s\n", path)
	_, signer, err := loadEntity(path)
	if err != nil {
		fmt.Printf("failed to load account entity, err: %s\n", err.Error())
		return nil, err
	}

	return signer, nil
}

func loadEntity(entityDir string) (*entity.Entity, signature.Signer, error) {
	factory := fileSigner.NewFactory(entityDir, signature.SignerEntity)
	return entity.Load(entityDir, factory)
}
