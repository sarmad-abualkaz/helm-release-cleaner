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
RUN CGO_ENABLExD=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o helm-release-cleaner main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM golang:1.18.4
WORKDIR /
COPY --from=builder /workspace/helm-release-cleaner .

ENTRYPOINT ["/helm-release-cleaner"]
