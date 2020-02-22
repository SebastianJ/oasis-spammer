package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/SebastianJ/oasis-spammer/config"
	"github.com/SebastianJ/oasis-spammer/rpc"
	"github.com/SebastianJ/oasis-spammer/utils"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/common/quantity"
	"github.com/oasislabs/oasis-core/go/consensus/api/transaction"
	"github.com/oasislabs/oasis-core/go/staking/api"
)

type Transfer struct {
	To     signature.PublicKey `json:"xfer_to"`
	Tokens quantity.Quantity   `json:"xfer_tokens"`
	Data   []byte              `json:"xfer_data"`
}

// AsyncSend - send transactions using goroutines/waitgroups
func AsyncSend(signer signature.Signer, amount string, nonce uint64, gasFee string, gasLimit uint64, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	Send(signer, amount, nonce, gasFee, gasLimit)
}

// Send - send transactions
func Send(signer signature.Signer, amount string, nonce uint64, gasFee string, gasLimit uint64) error {
	//defer signer.Reset()
	bigAmount, _ := utils.ConvertNumeralStringToBigInt(amount)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	toAddress := utils.RandomStringSliceItem(r, config.Configuration.Transactions.Receivers)

	var xfer Transfer
	if err := xfer.To.UnmarshalText([]byte(toAddress)); err != nil {
		fmt.Printf("failed to parse transfer destination ID, err: %s\n", err.Error())
		return err
	}

	if err := xfer.Tokens.FromBigInt(bigAmount); err != nil {
		fmt.Printf("failed to parse transfer amount, err: %s\n", err.Error())
		return err
	}

	xfer.Data = []byte(config.Configuration.Transactions.Data)

	var fee transaction.Fee
	if err := fee.Amount.UnmarshalText([]byte(gasFee)); err != nil {
		fmt.Printf("failed to parse fee amount, err: %s\n", err.Error())
		return err
	}
	fee.Gas = transaction.Gas(gasLimit)

	tx := transaction.NewTransaction(nonce, &fee, api.MethodTransfer, xfer)

	//tx := staking.NewTransferTx(nonce, &fee, &xfer)

	/*if config.Configuration.Verbose {
		tx.PrettyPrint("", os.Stdout)
		fmt.Println("")
	}*/

	if config.Configuration.Verbose {
		fmt.Printf("Sending Transaction:\n\tTo: %s\n\tAmount: %s\n\tNonce: %d\n\tData (bytes): %d\n\n", toAddress, amount, nonce, len(xfer.Data))
	}

	signedTx, _, err := sign(signer, tx)
	if err != nil {
		return err
	}

	/*if config.Configuration.Verbose {
		fmt.Printf("Signed tx: %s\n", rawSignedTx)
	}*/

	_, client, err := rpc.ConsensusClient(config.Configuration.Socket)
	if err != nil {
		return err
	}

	if err := client.SubmitTx(context.Background(), signedTx); err != nil {
		if !strings.Contains(err.Error(), "staking: forbidden by policy") {
			fmt.Printf("Failed to submit transaction, err: %s\n", err.Error())
			return err
		}
	}

	return nil
}

func sign(signer signature.Signer, tx *transaction.Transaction) (*transaction.SignedTransaction, []byte, error) {
	sigTx, err := transaction.Sign(signer, tx)
	if err != nil {
		fmt.Printf("failed to sign transaction, err: %s\n", err.Error())
		return nil, nil, err
	}

	/*if config.Configuration.Verbose {
		sigTx.PrettyPrint("", os.Stdout)
		fmt.Println("")
	}*/

	rawTx, err := json.Marshal(sigTx)
	if err != nil {
		fmt.Printf("failed to marshal transaction, err: %s\n", err.Error())
		return nil, nil, err
	}

	return sigTx, rawTx, nil
}
