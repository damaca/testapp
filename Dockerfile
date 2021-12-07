FROM golang:1.16 as builder

WORKDIR /app
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o testapp main.go

FROM gcr.io/distroless/static:nonroot
ENV APP_VERSION=1
WORKDIR /
COPY --from=builder /app/testapp .
USER 65532:65532

ENTRYPOINT ["/testapp"]
