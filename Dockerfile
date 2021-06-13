FROM golang:1.16-alpine AS build

LABEL MAINTAINER = 'Andrés Díaz (@mrandycodes) and Kevin Torres (@KDTV)'

RUN apk add --update git
RUN apk add ca-certificates
WORKDIR /go/src/github.com/mrandycodes/kyber
COPY . .
RUN go mod tidy && TAG=$(git describe --tags --abbrev=0) \
    && LDFLAGS=$(echo "-s -w -X main.version="$TAG) \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/kyber -ldflags "$LDFLAGS" cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/kyber /go/bin/kyber
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/kyber"]