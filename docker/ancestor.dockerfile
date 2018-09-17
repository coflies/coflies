########################################
# Code sandbok for compilers
# Based on Alpine
########################################

FROM alpine:3.8
LABEL maintainer="gaconkzk <gaconkzk@gmail.com>"

RUN apk add --no-cache \
    ca-certificates