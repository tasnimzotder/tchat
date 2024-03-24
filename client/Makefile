build_dev:
	go build -o bin/tchat main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/tchat.linux.amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/tchat.linux.arm64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/tchat.darwin.arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/tchat.windows.amd64.exe main.go

bin:
	sudo cp bin/tchat /usr/local/bin/tchat

test:
	go test -v -count=1 ./... -cover