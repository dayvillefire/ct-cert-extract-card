NAME := ct-cert-extract-card
GO   := /opt/go

all: clean build

build:
	#go generate
	#GOROOT=$(GO) CGO_ENABLED=0 GOOS=darwin  GOARCH=386   go build -o $(NAME).x86.darwin
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o $(NAME).x64.darwin
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=linux   GOARCH=386   go build -o $(NAME).x86.linux
	#GOROOT=$(GO) CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o $(NAME).x64.linux
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=windows GOARCH=386   go build -o $(NAME).win32.exe
	#GOROOT=$(GO) CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(NAME).win64.exe

clean:
	rm -f *.darwin *.linux *.exe *.syso
