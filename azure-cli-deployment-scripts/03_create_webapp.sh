#!/bin/bash
az webapp create \
    --resource-group ${RESOURCE_GROUP} \
    --plan ${SERVICE_PLAN} \
    --name azure-demo-app-yourname \
    --deployment-container-image-name poteb/azure-demo-app:v1.0.0
