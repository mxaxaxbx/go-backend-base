ARG GO_VERSION=1.18.1

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /backend-base

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY .env .
COPY --from=builder /backend-base /backend-base

EXPOSE 5050

ENTRYPOINT [ "/backend-base" ]

# Run the executable
CMD ["./main"]
