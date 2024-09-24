help:
	@echo
	@echo "  binary - build binary"
	@echo "  build-consumer - build docker images for centos"
	@echo "  push images to docker hub"

binary:
	go build -o bin/consumer cmd/consumer/main.go

test:
	go clean -testcache
	gotestsum --format pkgname
