# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/astrix-network/astrixd

WORKDIR /go/src/github.com/astrix-network/astrixd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache astrixd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o astrixwallet ./cmd/astrixwallet

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

COPY --from=build /go/src/github.com/astrix-network/astrixd/astrixwallet /app/
#COPY --from=build /go/src/github.com/astrix-network/astrixd/infrastructure/config/sample-astrixd.conf /app/

WORKDIR /
ENTRYPOINT ["/app/astrixwallet start-daemon"]
