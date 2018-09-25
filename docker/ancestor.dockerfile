########################################
# Code sandbok for compilers
# Based on Alpine
########################################

FROM alpine:3.8
LABEL maintainer="gaconkzk <gaconkzk@gmail.com>"

RUN apk add --no-cache \
    ca-certificates

# make sure the package repository is up to date
RUN apk update && apk upgrade && apk add --no-cache git bash make && rm -rf /var/cache/*/*
RUN sed -i -e "s/bin\/ash/bin\/bash/" /etc/passwd

# create coflies user
# set default home folder
# use coflies user
# copy coflies runner into docker

