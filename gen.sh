#!/bin/bash
ROOT_V1=api/v1/helloworld
protoc -I${ROOT_V1} --gogo_out=plugins=grpc:${ROOT_V1} ${ROOT_V1}/helloworld.proto

ROOT_V2=api/v2/helloworld
protoc -I${ROOT_V2} --gogo_out=plugins=grpc:${ROOT_V2} ${ROOT_V2}/helloworld.proto

ROOT_V3=api/v3/helloworld
protoc -I${ROOT_V3} --gogo_out=plugins=grpc:${ROOT_V3} ${ROOT_V3}/helloworld.proto
