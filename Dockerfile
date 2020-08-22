From golang:1.15-buster AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir bin && CGO_ENABLED=0 go build -o ./bin ./main

FROM alpine

WORKDIR /app

COPY --from=builder /build/bin/main ./

COPY etc/ ./etc

CMD ["./main"]
