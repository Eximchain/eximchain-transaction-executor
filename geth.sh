#!/usr/bin/env bash

geth --dev \
  --rpc \
  --rpcapi admin,shh,personal,net,eth,web3,txpool \
  --rpccorsdomain "*" \
  --keystore ./keystore
