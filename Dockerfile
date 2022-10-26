FROM golang:1.19.2@sha256:2fddf0539591f8e364c9adb3d495d1ba2ca8a8df420ad23b58e7bcee7986ea6c AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir bin && CGO_ENABLED=0 go build -o ./bin ./main

FROM gcr.io/distroless/base@sha256:c06bf48fa67dab06db6027109d3d802aa5b7d213c86a9eabc4d83f806d18ce1c

WORKDIR /app

COPY --from=builder /build/bin/main ./

CMD ["./main"]
