help:
	@echo
	@echo "  binary - build binary"
	@echo "  build-producer - build docker images for centos"
	@echo "  swag - regenerate swag"
	@echo "  push images to docker hub"

swag:
	swag init -g cmd/producer/main.go

binary:
	go build -o bin/producer cmd/producer/main.go

test:
	go clean -testcache
	gotestsum --format pkgname
