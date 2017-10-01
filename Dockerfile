FROM alpine:3.6

MAINTAINER Egbert Pot

add target/linux_amd64/azure-demo-app /usr/local/bin/azure-demo-app

CMD azure-demo-app