FROM golang:1.19.2@sha256:2fddf0539591f8e364c9adb3d495d1ba2ca8a8df420ad23b58e7bcee7986ea6c AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir bin && CGO_ENABLED=0 go build -o ./bin ./main

FROM gcr.io/distroless/base@sha256:b31a6e02605827e77b7ebb82a0ac9669ec51091edd62c2c076175e05556f4ab9

WORKDIR /app

COPY --from=builder /build/bin/main ./

CMD ["./main"]
