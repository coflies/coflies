FROM coflies/golang

# Create runner directory
ENV RUNNER_DIR=/go/src/github.com/coflies/coflies

RUN mkdir -p ${RUNNER_DIR}
WORKDIR ${RUNNER_DIR}
