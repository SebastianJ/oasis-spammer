package transactions

import (
	"fmt"
	"math"
	"sync"

	"github.com/SebastianJ/oasis-spammer/config"
	"github.com/SebastianJ/oasis-spammer/rpc"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
)

// AsyncBulkSendTransactions - sends transactions in bulk asynchronously
func AsyncBulkSendTransactions(signer signature.Signer, amount string, nonce uint64, gasFee string, gasLimit uint64, count int, poolSize int) {
	pools := 1

	fmt.Println(fmt.Sprintf("Current nonce is: %d", nonce))

	if count > poolSize {
		pools = int(math.RoundToEven(float64(count) / float64(poolSize)))
		fmt.Println(fmt.Sprintf("Number of goroutine pools: %d", pools))
	}

	for poolIndex := 0; poolIndex < pools; poolIndex++ {
		var waitGroup sync.WaitGroup

		for i := 0; i < poolSize; i++ {
			newNonce, err := rpc.CurrentNonce(signer, config.Configuration.Socket)
			if err == nil {
				nonce = newNonce
				fmt.Println(fmt.Sprintf("Nonce refreshed! Nonce is now: %d", nonce))
			}

			waitGroup.Add(1)
			go AsyncSend(signer, amount, nonce, gasFee, gasLimit, &waitGroup)
			nonce++
		}

		waitGroup.Wait()
	}
}
