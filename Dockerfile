FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.23-alpine AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /go/src/tagliatelle
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/tagliatelle ./cmd/tagliatelle/main.go

FROM alpine:3.18
ENV CGO_ENABLED=0 GOOS=linux
RUN apk add --no-cache libc6-compat
COPY --from=builder /go/bin/tagliatelle /go/bin/tagliatelle
ENTRYPOINT ["/go/bin/tagliatelle"]
CMD ["-h"]
