#!/bin/bash
export RESOURCE_GROUP="YOUR-RESOURCE-GROUP"
export SERVICE_PLAN="YourAppServicePlan"

export LOCATION=westeurope
export LOCATION_SHORT=WE

# export LOCATION=northeurope
# export LOCATION_SHORT=NE

printf "\nThe resource group is set to: %s\nService plan is set to: %s.\n\nWe are going to deploy in region: %s\n\n" "${RESOURCE_GROUP}" "${SERVICE_PLAN}" "${LOCATION}"