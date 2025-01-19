### variables
CGO_ENABLED := "1"
TARGETOS := `go env GOOS`
TARGETARCH := `go env GOARCH`
LDFLAGS := "-s -w -extldflags '-static'"
GO_PACKAGES := `go list ./... | grep -v /vendor/ | tr '\n' ' '`

### recipes

# Define the lint rule
lint:
    @echo "Running golangci-lint"
    golangci-lint run --timeout 10m

fmt:
    gci write --skip-vendor --skip-generated -s standard -s default --custom-order .

vendor:
  go mod vendor

build: vendor
	CGO_ENABLED=0 GOOS={{TARGETOS}} GOARCH={{TARGETARCH}} go build -ldflags "{{LDFLAGS}}" -o dist/crow-autoscaler github.com/crowci/autoscaler/cmd/crow-autoscaler

test: vendor
  go test -race -cover -coverprofile autoscaler-coverage.out -timeout 30s {{GO_PACKAGES}}

# env PLATFORMS='linux/amd64,linux/arm64' just image-autoscaler
image-autoscaler:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-autoscaler:dev -f Dockerfile --push .

generate:
	mockery
