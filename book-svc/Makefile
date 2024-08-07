proto: 
	protoc -I./protos -I./protos/google/api \
  --go_out=./protos --go_opt=paths=source_relative \
  --go-grpc_out=./protos --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./protos --grpc-gateway_opt=paths=source_relative \
  ./protos/book.proto

dockerup: 
	docker-compose up -d

run: 
	go run cmd/main.go

.PHONY:
	proto dockerup run