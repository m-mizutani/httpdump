FROM golang:1.21.5 AS build-go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ARG BUILD_VERSION
COPY . /app
WORKDIR /app
RUN go build -o httpdump

FROM gcr.io/distroless/base
COPY --from=build-go /app/httpdump /httpdump

EXPOSE 8080
ENTRYPOINT ["/httpdump", "serve"]
