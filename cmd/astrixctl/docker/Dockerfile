# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/astrix-network/astrixd

WORKDIR /go/src/github.com/astrix-network/astrixd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/astrix-network/astrixd/cmd/astrixctl

RUN GOOS=linux go build -a -installsuffix cgo -o astrixctl .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/astrix-network/astrixd/cmd/astrixctl/astrixctl /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
