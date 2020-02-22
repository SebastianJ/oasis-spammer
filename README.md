# oasis-spammer
This is a tool to spam Oasis chain with invalid transactions.

## Prerequisites
You need to import an existing key with funds to the keystore.

hmy is automatically downloaded as a part of the installation script.

Import a key using the following command:
```
./hmy keys import-ks ABSOLUTE_PATH_TO_YOUR_KEY NAME_OF_YOUR_KEY --passphrase ""
```

Find the address of your newly imported key:
```
./hmy keys list
```

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
