build:
	LOG_LEVEL=error go build -o bin/tchat main.go

bin:
	sudo cp bin/tchat /usr/local/bin/

test:
	go test -v -count=1 ./... -cover