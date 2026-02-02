.PHONY: all clean linux windows

BINARY_NAME=hassin8or
LDFLAGS=-ldflags="-s -w"

all: linux windows

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $(BINARY_NAME)-linux-amd64

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $(BINARY_NAME)-windows-amd64.exe

clean:
	rm -f $(BINARY_NAME)-linux-amd64 $(BINARY_NAME)-windows-amd64.exe
