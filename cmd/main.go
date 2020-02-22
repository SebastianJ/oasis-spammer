package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/SebastianJ/oasis-spammer/config"
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
			Usage: "The path relative to the binary where config.yml and other files can be found",
			Value: "./",
		},

		cli.StringFlag{
			Name:  "genesis-file",
			Usage: "The path to the genesis file",
			Value: "./etc/genesis.json",
		},

		cli.StringFlag{
			Name:  "entity-path",
			Usage: "The path to the genesis file",
			Value: "./node/entity",
		},

		cli.StringFlag{
			Name:  "socket",
			Usage: "The path to the socket to use for sending API requests",
			Value: "unix:node/internal.sock",
		},

		cli.IntFlag{
			Name:  "count",
			Usage: "How many transactions to send in total",
			Value: 1000,
		},

		cli.IntFlag{
			Name:  "pool-size",
			Usage: "How many transactions to send simultaneously",
			Value: 100,
		},

		cli.StringFlag{
			Name:  "data",
			Usage: "The tx data to send with each request",
			Value: "",
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

	if err := config.Configure(basePath, context); err != nil {
		return err
	}

	transactions.AsyncBulkSendTransactions(config.Configuration.Transactions.Signer, config.Configuration.Transactions.Amount, config.Configuration.Transactions.Nonce, config.Configuration.Transactions.Gas.Fee, config.Configuration.Transactions.Gas.Limit, config.Configuration.Transactions.Count, config.Configuration.Transactions.PoolSize)

	return nil
}
