FROM golang:1.16 as builder
WORKDIR /app
ENV CGO_ENABLED=0

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /dist/app .

FROM alpine:3
RUN echo $(ls -l)
COPY --from=builder /dist/app .
EXPOSE 8010
CMD ["./app"]