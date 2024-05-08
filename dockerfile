FROM golang:1.21-bullseye as builder

LABEL org.opencontainers.image.authors="hong"

ARG SOURCE_FILES

WORKDIR /app

COPY . /app
RUN go build -o $SOURCE_FILES

FROM busybox:1.34.0-glibc

COPY --from=builder /app /app

WORKDIR /app

RUN chmod +x main

CMD ["./main"]