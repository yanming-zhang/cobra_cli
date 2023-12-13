#!/bin/bash

## Load Environment variables
cd /data/optimism
direnv allow

## configure network
source .envrc
cd packages/contracts-bedrock
./scripts/getting-started/config.sh
