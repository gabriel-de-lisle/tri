protobuf:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
    protocol/tri.proto

server/build:
	docker build --tag tri-server .

server/start:
	docker run -p 5001:5001 tri-server 

client/install:
	cd client; \
	go build -o $$(go env GOPATH)/bin/tri
