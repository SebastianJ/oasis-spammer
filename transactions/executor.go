package transactions

import (
	"fmt"
	"math"
	"sync"

	"github.com/SebastianJ/oasis-spammer/rpc"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
)

// AsyncBulkSendTransactions - sends transactions in bulk asynchronously
func AsyncBulkSendTransactions(signer signature.Signer, to string, amount string, nonce uint64, gasFee string, gasLimit uint64, socket string, count int, poolSize int) {
	pools := 1

	fmt.Println(fmt.Sprintf("Current nonce is: %d", nonce))

	newNonce, err := rpc.CurrentNonce(signer, socket)
	if err == nil {
		fmt.Printf("NewNonce: %+v\n", newNonce)
	}

	if count > poolSize {
		pools = int(math.RoundToEven(float64(count) / float64(poolSize)))
		fmt.Println(fmt.Sprintf("Number of goroutine pools: %d", pools))
	}

	for poolIndex := 0; poolIndex < pools; poolIndex++ {
		var waitGroup sync.WaitGroup

		if poolIndex > 1 {
			//refresh nonce
			//fmt.Println(fmt.Sprintf("Nonce refreshed! Nonce is now: %d", currentNonce))
		}

		for i := 0; i < poolSize; i++ {
			waitGroup.Add(1)
			go AsyncSend(signer, to, amount, nonce, gasFee, gasLimit, socket, &waitGroup)
			nonce++
		}

		waitGroup.Wait()
	}
}