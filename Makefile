prepare-workspace:
	minikube addons disable ingress
	kubectl create namespace monitoring || echo "monitoring namespace already exists"
	kubectl create namespace user || echo "user namespace already exists"
	kubectl config set-context --current --namespace=monitoring
	sleep 1
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	helm repo add stable https://charts.helm.sh/stable
	helm repo update
	helm install prom prometheus-community/kube-prometheus-stack -f prometheus.yaml --atomic
	helm install nginx ingress-nginx/ingress-nginx -f nginx-ingress.yaml --atomic

start-user-service:
	kubectl config set-context --current --namespace=user
	sleep 1
	helm install userrelease ./user-chart/ -n user
	kubectl apply -f grafana.yaml

load-test:
	go run ./load_testing/

open-grafana:
	kubectl config set-context --current --namespace=monitoring
	sleep 1
	echo navigate to http://localhost:9000
	kubectl port-forward service/prom-grafana 9000:80

open-prometheus:
	kubectl config set-context --current --namespace=monitoring
	sleep 1
	echo navigate to http://localhost:9090
	kubectl port-forward service/prom-kube-prometheus-stack-prometheus 9090
