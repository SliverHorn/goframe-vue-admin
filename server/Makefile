BINARY="github.com/flipped-aurora/gf-vue-admin/server"
GFVA = "gfva"

all: check gfva run

gfva:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.io,direct
	go build -o ${GFVA} cmd/main.go
	@if [ -f ${GFVA} ] ; then ./${GFVA} initdb && rm ${GFVA} ; fi

linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}

run:
	@go run main.go

swagger:
	@gf swagger

check:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -f ${GFVA} ] ; then rm ${GFVA} ; fi

help:
	@echo "make - 构建gfva终端工具并初始化数据,初始化数据后删除gfva终端工具,启动server项目"
	@echo "make gfva - 构建gfva终端工具 并初始化数据"
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make run - 直接运行 main.go"
	@echo "make clean - 移除二进制文件"
	@echo "make check - 运行 Go 工具 'fmt' and 'vet'"