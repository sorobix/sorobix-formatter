FROM golang:1.19
RUN make build

EXPOSE 3000
CMD ["./bin/sorobix-formatter"]