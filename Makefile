build-image:
	dagger call build-image --dir .

tests:
	dagger call tests --dir .

build-all:
	dagger call build-all --dir . --platforms "linux/amd64,linux/arm64,darwin/arm64" -o ./bin

build-binary:
	dagger call build-binary --dir . export --path ./bin/app
