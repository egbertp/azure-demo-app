FROM alpine:3.6

MAINTAINER Egbert Pot

add target/linux_amd64/azure-demo-app /usr/local/bin/azure-demo-app

# Configure ports
EXPOSE 8000

ENV AZURE_DEMO_APP_ADDRESS "0.0.0.0:8000"

CMD azure-demo-app