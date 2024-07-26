FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -v ./cmd/apiserver
CMD ["./apiserver"]