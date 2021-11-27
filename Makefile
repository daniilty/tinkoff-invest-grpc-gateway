build:
	go build -o server github.com/daniilty/tinkoff-invest-grpc-gateway/cmd/server
build_docker:
	docker build -t tinkoff-invest-grpc-gateway:latest -f docker/Dockerfile .

