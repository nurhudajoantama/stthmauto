FROM golang:1.25-alpine AS builder
WORKDIR /src

# cache deps
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /out/stthmauto ./

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/stthmauto /usr/local/bin/stthmauto
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/stthmauto"]
