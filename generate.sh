#!/usr/bin/env bash

cd src
protoc -I . --go_out=grpc_plugins=grpc:../golang ./*.proto
cd ..
