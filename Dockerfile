FROM golang:1.19 as builder
WORKDIR /go/src/app
COPY . .
RUN make build

FROM rust:alpine
RUN rustup component add rustfmt
COPY --from=builder /go/src/app/bin/sorobix-formatter ./
EXPOSE 3000
CMD ["./sorobix-formatter"]