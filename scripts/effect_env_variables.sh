#!/bin/bash

basepath=`pwd`

## Load Environment variables
cd $basepath/optimism
direnv allow

## configure network
source .envrc
cd packages/contracts-bedrock
./scripts/getting-started/config.sh
