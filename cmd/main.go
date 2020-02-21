package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/SebastianJ/oasis-spammer/genesis"
	"github.com/SebastianJ/oasis-spammer/transactions"

	"github.com/urfave/cli"
)

func main() {
	// Force usage of Go's own DNS implementation
	os.Setenv("GODEBUG", "netdns=go")

	app := cli.NewApp()
	app.Name = "Harmony Tx Sender - stress test and bulk transaction sending tool"
	app.Version = fmt.Sprintf("%s/%s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	app.Usage = "Use --help to see all available arguments"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "to",
			Usage: "Where to send the tokens",
			Value: "8uNiDud/L0d0muEGb2t5BFnjupStWasyjdErHFnjQXk=",
		},

		cli.StringFlag{
			Name:  "amount",
			Usage: "How many tokens to send per transaction",
			Value: "0.000000000000000001",
		},

		cli.Uint64Flag{
			Name:  "nonce",
			Usage: "The nonce to use when sending txs",
			Value: 0,
		},

		cli.StringFlag{
			Name:  "gas.fee",
			Usage: "The gas fee to pay when sending txs",
			Value: "1",
		},

		cli.Uint64Flag{
			Name:  "gas.limit",
			Usage: "The gas limit to use when sending txs",
			Value: 1000,
		},

		cli.StringFlag{
			Name:  "path",
			Usage: "The base path for all files and certificates",
			Value: "./data",
		},

		cli.StringFlag{
			Name:  "socket",
			Usage: "The path to the socket to use for sending API requests",
			Value: "unix:node/internal.sock",
		},

		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable more verbose output",
		},
	}

	app.Authors = []cli.Author{
		{
			Name:  "Sebastian Johnsson",
			Email: "",
		},
	}

	app.Action = func(context *cli.Context) error {
		return run(context)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run(context *cli.Context) error {
	basePath, err := filepath.Abs(context.GlobalString("path"))
	if err != nil {
		return err
	}

	if err := genesis.LoadDocument(basePath); err != nil {
		return err
	}

	to := context.GlobalString("to")
	amount := context.GlobalString("amount")
	nonce := context.GlobalUint64("nonce")
	gasFee := context.GlobalString("gas.fee")
	gasLimit := context.GlobalUint64("gas.limit")
	socket := context.GlobalString("socket")
	fmt.Println(fmt.Sprintf("The socket address is now %s", socket))

	signer, err := transactions.LoadSigner(basePath)
	if err != nil {
		return err
	}

	err = transactions.Send(signer, to, amount, nonce, gasFee, gasLimit, socket)
	if err != nil {
		return err
	}

	return nil
}
