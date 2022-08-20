export RESOURCE_GROUP=akstestrg
export CLUSTER_NAME=aks-testing-ziqq
export LOCATION=westeurope

# Run if you don't have resource group
# az group create --name=$RESOURCE_GROUP --location=$LOCATION

# It is not a good practice but for simplicity we will use system node to host app
az aks create \
    --resource-group $RESOURCE_GROUP \
    --name $CLUSTER_NAME \
    --node-count 1 \
    --enable-addons http_application_routing \
    --generate-ssh-keys \
    --node-vm-size Standard_B2ms \
    --network-plugin azure

# az aks nodepool add \
#     --resource-group $RESOURCE_GROUP \
#     --cluster-name $CLUSTER_NAME \
#     --name userpool \
#     --node-count 2 \
#     --node-vm-size Standard_B2ms

# The above command adds a new node pool (User mode) to an existing AKS cluster (created in previous command). 
# This new node pool can be used to host applications and workloads, 
# instead of using System node pool, which gets created during previous command az aks create.

