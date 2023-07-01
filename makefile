.PHONY: all run clean help

APP = XDataFlowProxy
RACE = -gcflags='-l -l -l' -ldflags='-s -w' -pgo=auto
GLOBAL_CONFIG = CGO_ENABLED=0

## linux: 编译打包linux
.PHONY: linux-amd64
linux-amd64:
	${GLOBAL_CONFIG} GOOS=linux GOARCH=amd64 ${GO_PATH}go build $(RACE) -o ./bin/${APP}_linux_amd64 ./main.go
.PHONY: linux-arm64
linux-arm64:
	${GLOBAL_CONFIG} GOOS=linux GOARCH=arm64 ${GO_PATH}go build $(RACE) -o ./bin/${APP}_linux_arm64 ./main.go

## win: 编译打包win
.PHONY: win-amd64
win-amd64:
	${GLOBAL_CONFIG} GOOS=windows GOARCH=amd64 ${GO_PATH}go build $(RACE) -o ./bin/${APP}_windows_amd64.exe ./main.go
.PHONY: win-arm64
win-arm64:
	${GLOBAL_CONFIG} GOOS=windows GOARCH=arm64 ${GO_PATH}go build $(RACE) -o ./bin/${APP}_windows_arm64.exe ./main.go

## mac: 编译打包mac
.PHONY: mac-amd64
mac-amd64:
	${GLOBAL_CONFIG} GOOS=darwin GOARCH=amd64 ${GO_PATH}go build $(RACE) -o ./bin/${APP}_darwin_amd64 ./main.go
.PHONY: mac-arm64
mac-arm64:
	${GLOBAL_CONFIG} GOOS=darwin GOARCH=arm64 ${GO_PATH}go build $(RACE) -o ./bin/${APP}_darwin_arm64 ./main.go

.PHONY: darwin-arm64-lib
darwin-arm64-lib:
	go build $(RACE) -buildmode=c-shared -o ./bin/${APP}darwin_arm64_lib.so ./main.go

## 编译win，linux，mac平台
.PHONY: all
all:win-amd64 win-arm64 linux-amd64 linux-arm64 mac-amd64 mac-arm64

run:
	@go run ./

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: update
update:
	@go get -u

## 清理二进制文件
clean:
	@if [ -f ./bin/${APP}_darwin_amd64 ] ; then rm ./bin/${APP}_darwin_amd64; fi
	@if [ -f ./bin/${APP}_windows_amd64.exe ] ; then rm ./bin/${APP}_windows_amd64.exe; fi
	@if [ -f ./bin/${APP}_linux_amd64 ] ; then rm ./bin/${APP}_linux_amd64; fi
	@if [ -f ./bin/${APP}_darwin_arm64 ] ; then rm ./bin/${APP}_darwin_arm64; fi
	@if [ -f ./bin/${APP}_windows_arm64.exe ] ; then rm ./bin/${APP}_windows_arm64.exe; fi
	@if [ -f ./bin/${APP}_linux_arm64 ] ; then rm ./bin/${APP}_linux_arm64; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make mac-amd64 - 编译 Go 代码, 生成mac-amd64的二进制文件"
	@echo "make linux-amd64 - 编译 Go 代码, 生成linux-amd64二进制文件"
	@echo "make win-amd64 - 编译 Go 代码, 生成windows-amd64二进制文件"
	@echo "make mac-arm64 - 编译 Go 代码, 生成mac-arm64的二进制文件"
	@echo "make linux-arm64 - 编译 Go 代码, 生成linux-arm64二进制文件"
	@echo "make win-arm64 - 编译 Go 代码, 生成windows-arm64二进制文件"
	@echo "make tidy - 执行go mod tidy"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除编译的二进制文件"
	@echo "make all - 编译多平台的二进制文件"
	@echo "make update - 更新 mod 扩展库"