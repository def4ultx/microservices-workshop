# Deployment in kubernetes

- Create kind cluster

```bash
kind create cluster --config=kind-config.yml
```

- Add kind cluster info to your local kubectl config

```bash
kubectl cluster-info --context kind-kind
```

- Load docker image to kind nodes. You might need to load multiple images if you're deploying custom images or push to public registry

```bash
kind load docker-image microservice-workshop/inventory-api
```

- Apply kubernetes manifest

```bash
kubectl apply -f crdb-statefulsets.yml
kubectl apply -f inventory-api.yml
```

- Access Inventory API with host port 30000, crdb WebUI with port 30001
