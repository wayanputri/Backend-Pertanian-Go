PROTOC = protoc
PROTOC_GEN_GO = $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GRPC_GO = $(GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY = $(GOPATH)/bin/protoc-gen-grpc-gateway
PROTOC_GEN_SWAGGER = $(GOPATH)/bin/protoc-gen-swagger
PROTO_FILES = proto/*.proto
EXTERNAL_DIR = lib/external/googleapis
OUT_DIR_GO = .
OUT_DIR_SWAGGER = ./swagger

# Create output directories
$(shell mkdir -p $(OUT_DIR_GO) $(OUT_DIR_SWAGGER))

# Command to generate Go and Swagger code from Protobuf
generate:
	@echo "Generating Go code from Protobuf..."
	$(PROTOC) --proto_path=proto --proto_path=$(EXTERNAL_DIR) \
		--go_out=$(OUT_DIR_GO) --go-grpc_out=$(OUT_DIR_GO) \
		--grpc-gateway_out=$(OUT_DIR_GO) --swagger_out=logtostderr=true:$(OUT_DIR_SWAGGER) $(PROTO_FILES)

# Default command
.PHONY: generate

run:
	go run main.go

build:
	docker build -t backendApp .

