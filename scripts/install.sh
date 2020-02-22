#!/usr/bin/env bash

echo "Installing Oasis Spammer"
curl -LOs http://tools.oasis-protocol.org.s3.amazonaws.com/release/linux-x86_64/oasis-spammer && chmod u+x oasis-spammer
curl -LOs https://raw.githubusercontent.com/SebastianJ/oasis-spammer/master/config.yml
echo "Oasis Spammer is now ready to use!"
echo "Invoke it using ./oasis-spammer - see ./oasis-spammer --help for all available options"
