PROTO_DIR := proto
GO_OUT := go_out
GO_PROTO_PACKAGE := proto

# Protocol Buffers 정의 파일 경로
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Protocol Buffers를 사용하여 Go 코드 생성
generate:
	protoc -I=$(PROTO_DIR) --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)