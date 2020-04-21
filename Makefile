all: client server

client:
	@echo "build client image"
	./build/build-image.sh client

server:
	@echo "build server image"
	./build/build-image.sh server