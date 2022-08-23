export RESOURCE_GROUP=akstestrg
export LOCATION=westeurope
export ACR_NAME=acrziqqtest

az acr create -n $ACR_NAME -g $RESOURCE_GROUP --sku Basic --admin-enabled true --location $LOCATION 