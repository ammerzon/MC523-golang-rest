FROM golang:1.16 as builder
WORKDIR /app
ENV CGO_ENABLED=0

COPY go.* ./
RUN go mod download

COPY . /app
RUN go build -o ./server ./cmd

FROM alpine:3
WORKDIR /root/
COPY --from=builder /app/server /bin/server
EXPOSE 8010
CMD ["/bin/server"]