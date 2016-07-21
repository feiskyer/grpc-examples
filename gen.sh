#!/bin/bash

for version in `ls api/`; do
	PROTO_ROOT=api/${version}/helloworld
	protoc -I${PROTO_ROOT} --gogo_out=plugins=grpc:${PROTO_ROOT} ${PROTO_ROOT}/helloworld.proto
done
