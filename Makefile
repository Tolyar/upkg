build: bin
	go build -o bin/upkg cmd/upkg/main.go
linux: bin
	GOOS=linux GOARCH=amd64 go build -o bin/upkg.linux cmd/upkg/main.go
bin:
	mkdir bin
