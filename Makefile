# Image URL to use all building/pushing image targets
IMG ?= ghcr.io/naveensrinivasan/improved-invention
TAG ?= $(eval TAG := $(shell date +v%Y%m%d)-$(shell git describe --tags --always --dirty)-$(shell git diff | shasum -a256 | cut -c -6))$(TAG)

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Run tests
test: fmt vet
	go test ./...

# Build tests
build: fmt vet
	go test ./... -c -o ./static

clean:
	rm  -f ./static

docker-build: test
	docker build . -t $(IMG):$(TAG)

docker-push: docker-build
	docker push $(IMG):$(TAG)

test-network:
	ginkgo  --focus=NETWORK ./...

test-database:
	ginkgo --focus=DATABASE ./...
