#!/bin/bash

apt update
apt install -y git

if [ -z `which node` ]; then
  wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
  tar xzf go1.21.5.linux-amd64.tar.gz -C /usr/local
  ln -s /usr/local/go/bin/go /usr/local/bin/go
  ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
  rm -f go1.21.5.linux-amd64.tar.gz
fi

if [ -z `which node` ]; then
  wget https://nodejs.org/dist/v20.10.0/node-v20.10.0-linux-x64.tar.xz
  tar xf node-v20.10.0-linux-x64.tar.xz
  mv node-v20.10.0-linux-x64 /usr/local/node
  ln -s /usr/local/node/bin/node /usr/local/bin/node
  ln -s /usr/local/node/bin/npm /usr/local/bin/npm
  ln -s /usr/local/node/bin/npx /usr/local/bin/npx
  rm -f node-v20.10.0-linux-x64.tar.xz
fi

curl -fsSL https://get.pnpm.io/install.sh | sh -
source ~/.bashrc

if [ -z `which forge` ]; then
  wget https://github.com/foundry-rs/foundry/releases/download/nightly-cdbaf9dda688cab08b9f6945af287534d68b1e1f/foundry_nightly_linux_amd64.tar.gz
  tar xvzf foundry_nightly_linux_amd64.tar.gz -C /usr/local/bin/
  rm -f foundry_nightly_linux_amd64.tar.gz
fi

apt install -y make jq

if [ -z `which direnv` ]; then
  curl -sfL https://direnv.net/install.sh | bash
  echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
  source ~/.bashrc
fi

