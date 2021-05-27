#!/bin/bash
protoc *.proto --go_out=./ --go-grpc_out=./