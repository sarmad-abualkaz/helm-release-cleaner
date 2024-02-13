# Build the manager binary
FROM golang:1.18.4 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY helm/ helm/
COPY cmd/ cmd/
COPY util/ util/

# Build
RUN CGO_ENABLED=0 go build -o /go/bin/helm-release-cleaner

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder /go/bin/tophat-cleaner /
ENTRYPOINT ["/helm-release-cleaner"]
