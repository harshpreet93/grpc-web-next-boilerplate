FROM golangci/golangci-lint:v1.28-alpine AS lint
COPY . /app
WORKDIR /app
RUN golangci-lint run --timeout 3m

FROM node:12.18.3-alpine3.11 as frontend-build
COPY ./frontend /app
WORKDIR /app
RUN apk add yarn
RUN yarn build_static

FROM golang:1.15.4-alpine3.12 AS backend-build
COPY . /app
COPY --from=frontend-build /app/out ./frontend/out 
WORKDIR /app/backend
RUN ["go", "build", "-o", "app", "."]

FROM alpine:3.12.0 AS app
# COPY --from=frontend-build /app/out ./out 
COPY --from=backend-build /app/backend/app .
ENTRYPOINT ./app