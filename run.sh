#!/bin/bash

./cobra_cli depend

./cobra_cli build -t optimism
./cobra_cli build -t op-geth

./cobra_cli env_vars

