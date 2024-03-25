build_dev:
	go build -o bin/tchat main.go

build:
	CGO_ENABLED=1 go build -o bin/tchat main.go

bin:
	sudo cp bin/tchat /usr/local/bin/tchat

test:
	go test -v -count=1 ./... -cover
