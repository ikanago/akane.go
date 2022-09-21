FROM golang:1.18.6@sha256:ebd59a91af0c61cf28a22a9651ffd45612848bdcc1a43c2c3063825ba2e613fc AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir bin && CGO_ENABLED=0 go build -o ./bin ./main

FROM gcr.io/distroless/base@sha256:c06bf48fa67dab06db6027109d3d802aa5b7d213c86a9eabc4d83f806d18ce1c

WORKDIR /app

COPY --from=builder /build/bin/main ./

CMD ["./main"]
