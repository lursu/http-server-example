FROM golang:1.19.1 AS builder

ARG workspace=/go/src/web-server

RUN mkdir -p ${workspace}
WORKDIR ${workspace}


COPY go.mod go.sum ${workspace} 
RUN go mod download 

COPY . ${workspace} 

RUN CGO_ENABLED=0 go build -a ./cmd/example-server

FROM alpine:latest

COPY --from=builder /go/src/web-server/example-server ./ 
CMD ["./example-server"]