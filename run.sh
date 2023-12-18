#!/bin/bash

## 安装依赖包
./cobra_cli depend

## 编译程序包
./cobra_cli build -t optimism
./cobra_cli build -t op-geth

## 设置环境变量与配置网络
./cobra_cli env_vars
