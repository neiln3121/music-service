FROM golang:1.13-buster as builder

RUN apt-get update && apt-get install -y --no-install-recommends \
	&& rm -rf /var/lib/apt/lists/*

COPY . /build
WORKDIR /build
RUN go build -mod vendor -o bin/music-service main.go

FROM debian:buster-slim

COPY --from=builder /build/bin/music-service /usr/local/bin

ENV PATH="/usr/local/bin:${PATH}"

EXPOSE 8080
CMD ["music-service"]

