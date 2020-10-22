FROM golang:alpine AS build-env
WORKDIR /lile_test
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /lile_test/go.mod
COPY go.sum /lile_test/go.sum
RUN go mod download
COPY . /lile_test
RUN CGO_ENABLED=0 GOOS=linux go build -o build/lile_test ./lile_test


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /lile_test/build/lile_test /
ENTRYPOINT ["/lile_test"]
CMD ["up", "--grpc-port=80"]
