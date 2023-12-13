#!/bin/bash

build_type=$1
basepath=`pwd`

if [ $build_type == "optimism" ]; then
  echo "building optimism......"
  cd $basepath/optimism
  pnpm install
  make op-node op-batcher op-proposer
  pnpm build
elif [ $build_type == "op-geth" ]; then
  echo "building op-geth......"
  cd $basepath/op-geth
  make geth
else
  exit 1
fi
