# HW Prometheus. Grafana

### Clone the repo:

```bash
git clone https://github.com/alikhanmurzayev/otus_prometheus_grafana.git && cd otus_prometheus_grafana
```

### Prepare workspace:

```bash
make prepare-workspace
```

### Start User service:

```bash
make start-user-service
```

Wait a little.

### Health check

```bash
curl -H 'Host: arch.homework' "http://$(minikube ip)/health"
# Output: {"status": "ok"}
```

### Run test:

```bash
go run ./load_testing/
```

### Grafana dashboard

![Screenshot](./screenshots/1.png?raw=true)
![Screenshot](./screenshots/2.png?raw=true)

### Alert

![Screenshot](./screenshots/3.png?raw=true)
