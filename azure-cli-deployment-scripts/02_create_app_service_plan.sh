#!/bin/bash

az appservice plan create \
    --name ${SERVICE_PLAN} \
    --resource-group ${RESOURCE_GROUP} \
    --sku S1 \
    --is-linux
    