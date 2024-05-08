build-all:
	GOOS=linux GOARCH=amd64 go build -o ./build/firesui-linux-amd64 ./cmd/firesui/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/firesui-linux-arm64 ./cmd/firesui/main.go
	GOOS=darwin GOARCH=amd64 go build -o ./build/firesui-darwin-amd64 ./cmd/firesui/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/firesui-darwin-arm64 ./cmd/firesui/main.go

do-checksum:
	cd build && sha256sum \
		firesui-linux-amd64 firesui-linux-arm64 \
		firesui-darwin-amd64 firesui-darwin-arm64 \
		> selfchain_checksum

build-with-checksum: build-all do-checksum
