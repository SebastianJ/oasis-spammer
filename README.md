# oasis-spammer
This is a tool to spam Oasis chain with invalid transactions.

## Prerequisites
You need to have access to the genesis.json file for the network you want to interact with.

You also need the entity.pem file for an account with sufficient funds to send transactions.

## Installation

```
bash <(curl -s -S -L https://raw.githubusercontent.com/SebastianJ/oasis-spammer/master/scripts/install.sh)
```

The installer script will also create the data/ folder where you'll find the files receivers.txt and data.txt

`data/receivers.txt` is the file where you enter the receiver accounts you want the tx sender to send tokens to
`data/data.txt` is the tx data that you want the tx sender to use for every transaction it sends.

## Usage
```
./oasis-spammer --genesis-file PATH/TO/genesis.json --entity-path PATH/TO/entity/ --socket unix:internal.sock --count 100000 --pool-size 100
```

### All options:

```
NAME:
   Oasis spammer - spam and stress test transactions on Oasis' network - Use --help to see all available arguments

USAGE:
   oasis-spammer [global options] command [command options] [arguments...]

VERSION:
   go1.13.7/linux-amd64

AUTHOR:
   Sebastian Johnsson

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --amount value        How many tokens to send per transaction (default: "0.000000000000000001")
   --nonce value         The nonce to use when sending txs (default: 0)
   --gas.fee value       The gas fee to pay when sending txs (default: "1")
   --gas.limit value     The gas limit to use when sending txs (default: 1000)
   --path value          The path relative to the binary where config.yml and other files can be found (default: "./")
   --genesis-file value  The path to the genesis file (default: "./etc/genesis.json")
   --entity-path value   The path to the genesis file (default: "./node/entity")
   --socket value        The path to the socket to use for sending API requests (default: "unix:node/internal.sock")
   --count value         How many transactions to send in total (default: 1000)
   --pool-size value     How many transactions to send simultaneously (default: 100)
   --data value          The tx data to send with each request
   --verbose             Enable more verbose output
   --help, -h            show help
   --version, -v         print the version
```
