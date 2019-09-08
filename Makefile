# Go parameters
SERVER_MAIN_PATH=cmd/server/main.go
CLIENT_MAIN_PATH=cmd/client/main.go
BINARY_SERVER_NAME=$(BINARY_PATH)/wishlist-server
BINARY_CLIENT_NAME=$(BINARY_PATH)/wishlist

BINARY_PATH=bin

server-run: proto
	go build -o $(BINARY_SERVER_NAME) -race $(SERVER_MAIN_PATH)
	./$(BINARY_SERVER_NAME)

client-run: proto
	go build -o $(BINARY_CLIENT_NAME) -race $(CLIENT_MAIN_PATH)
	./$(BINARY_CLIENT_NAME)

clean:
	go clean $(CLIENT_MAIN_PATH)
	go clean $(SERVER_MAIN_PATH)
	rm -f $(BINARY_PATH)/*


# Proto parameters
PROTO_FILES_PATH=api/proto
PROTO_OUT=internal/net/grpc

proto:
	protoc -I $(PROTO_FILES_PATH) --go_out=plugins=grpc:$(PROTO_OUT) $(PROTO_FILES_PATH)/*.proto

clean-proto:
	rm -f $(PROTO_OUT)/*.pb.go