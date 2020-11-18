#!/bin/sh
(cd frontend && yarn build_static)
(cd backend && go run main.go)