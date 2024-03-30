
all: go-build go-test docker-build-all

go-build:
	dagger call build --source . --platforms "linux/amd64,linux/arm64" -o bin

go-test:
	dagger call test --source . --verbose=true

docker-build-all:
	dagger call docker-build --bin bin --platform "linux/amd64" export --path image-linux-amd64.tar
	dagger call docker-build --bin bin --platform "linux/arm64" export --path image-linux-arm64.tar
