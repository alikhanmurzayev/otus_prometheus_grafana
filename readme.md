# HW Prometheus. Grafana

### Clone the repo:

```bash
git clone https://github.com/alikhanmurzayev/otus_kuber_part_3.git && cd otus_kuber_part_3
```

### Prepare workspace:

```bash
minikube addons enable ingress
kubectl create namespace user
kubectl config set-context --current --namespace user
```

### Start services:

```bash
kubectl apply -f .
```

Wait a little.

### Health check

```bash
curl -H 'Host: arch.homework' "http://arch.homework/health"
# Output: {"status": "ok"}
```

### Run test:

```bash
newman run newman run postman_test/postman_test.json 
```
