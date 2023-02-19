#!/bin/bash

go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go env -w GOPROXY=https://goproxy.io,direct

echo "export GOPROXY=https://mirrors.aliyun.com/goproxy/" >> ~/.profile && source ~/.profile

