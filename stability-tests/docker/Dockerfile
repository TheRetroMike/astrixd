ARG ASTRIXD_IMAGE
ARG ASTRIXMINER_IMAGE

FROM ${ASTRIXD_IMAGE} as astrixd
FROM ${ASTRIXMINER_IMAGE} as astrixminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/astrix-network/astrixd

WORKDIR /go/src/github.com/astrix-network/astrixd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=astrixd /app/ /app/
COPY --from=astrixminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/astrix-network/astrixd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]