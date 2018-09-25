# build the runner inside docker
SRC_DIR=$PWD DST_DIR=/go/src/github.com/coflies/coflies make build-in-dock

docker build -f docker/ancestor.dockerfile -t coflies/ancestor:latest .