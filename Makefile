NAME=coflies

ifeq ($(OS),Windows_NT)
	SDIR=$(shell cygpath -u ${SRC_DIR})
	DDIR=$(shell cygpath -u ${DST_DIR})
else
  SDIR=${SRC_DIR}
	DDIR=${DST_DIR}
endif

all: clean get vet build

print-%:
	@echo $*

check:
	@scripts/check.sh

vet:
	go vet .

get:
	go get

build:
	CGO_ENABLED=0 go build -o $(NAME)

clean:
	go clean

build-in-dock:
	@docker run --mount 'type=bind,src=${SDIR},dst=/go/src/github.com/coflies/coflies' --rm -it coflies/runner make all

format:
	gofmt -l -w .
.PHONY:
	all-dock \
	build-in-dock \
	clean \
	vet \
	get \
	check \
	build
