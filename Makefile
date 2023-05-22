BINARY_NAME=selection
.PHONY: all dep build
# 默认目标
all: build
# 更新依赖
dep:
	go mod tidy
# 编译
build:
	go mod tidy
	go env -w CGO_ENABLED=0
	go env -w GOARCH=amd64
	go env -w GOOS=windows
	go build -o $(BINARY_NAME).exe .