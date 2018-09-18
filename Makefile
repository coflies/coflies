NAME=coflies

all: vet build

vet:
	go vet .

build:
	CGO_ENABLED=0 go build -tags netgo -o $(NAME)

format:
	gofmt -l -w .
.PHONY:
	all \
	vet \
	build
