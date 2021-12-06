FROM golang:1.16.5 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN export CGO_ENABLED=0 && \
export GOOS=linux && \
go build -o main .

FROM alpine:latest AS production

COPY --from=builder /app .
CMD ["./main"]