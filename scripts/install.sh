#!/usr/bin/env bash

echo "Installing Oasis Spammer"
curl -LOs http://tools.oasis-protocol.org.s3.amazonaws.com/release/linux-x86_64/oasis-spammer && chmod u+x oasis-spammer
curl -LOs https://raw.githubusercontent.com/SebastianJ/oasis-spammer/master/config.yml
mkdir -p data && cd data && touch data.txt && curl -LOs https://gist.githubusercontent.com/SebastianJ/a855d7aae724d6a8fb6cd143cdf222eb/raw/f43e00b8e9f9ae910fc6d13d55ae085bbfbac18f/receivers.txt && cd ..
echo "Oasis Spammer is now ready to use!"
echo "Invoke it using ./oasis-spammer - see ./oasis-spammer --help for all available options"
