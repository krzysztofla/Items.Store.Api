# Items.Store.Api
Use make to compose images and start the application.

Kubernetes commands:

```bash 
minikube start // starts docker cluster with one node
minikube stop 

minikube tunnel <service name> // starts tunneling/proxy to node via configured ip and protocol
minikube service <service name>

kubectl apply -f  deployment/aks/pods/api/api-deployment.yml // deployments

kubectl get deployments
kubectl get nodes
kubectl get pods

kubectl describe pod/node items-store-api-854887494d-g22xm // showing details about pod/node like ip, names, docker ...

kubectl logs items-store-api-776fc8b99c-24kw7 // showing pod/container logs
``` 