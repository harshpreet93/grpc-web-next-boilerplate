#!/bin/bash
set -euo pipefail
docker run --rm -v $PWD:/workspace:consistent harshpreet93/rpcz:latest
rm -rf ../backend/grpc_gen
mv grpc_gen ../backend/grpc_gen
mv service_client_out ../frontend