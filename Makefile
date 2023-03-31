GO := go
GOLINT := golangci-lint

lint:
	$(GOLINT) run

vet:
	go vet ./...

test:
	$(GO) test -v -coverprofile cover.out -cover .

cover: | test
	$(GO) tool cover -html cover.out

build: bin
	$(GO) build -o bin/upkg cmd/upkg/main.go

run: build
	./bin/upkg $(filter-out $@, $(MAKECMDGOALS))

linux: bin
	GOOS=linux GOARCH=amd64 $(GO) build -o bin/upkg.linux cmd/upkg/main.go

bin:
	mkdir bin

%:
	@true
