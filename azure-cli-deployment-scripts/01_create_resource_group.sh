#!/bin/bash

printf "INFO: You are about to create the resource group %s\n" "${RESOURCE_GROUP}"
az group create \
	--location ${LOCATION} \
	--name "${RESOURCE_GROUP}"
