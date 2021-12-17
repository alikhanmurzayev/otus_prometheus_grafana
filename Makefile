prepare-workspace:
	minikube addons disable ingress
	kubectl create namespace monitoring
	kubectl create namespace user
	kubectl config set-context --current --namespace=monitoring
	helm repo add prometheus-community
	helm repo add ingress-nginx
	helm repo add stable https://charts.helm.sh/stable
	helm repo update
	helm install prom prometheus-community/kube-prometheus-stack -f prometheus.yaml --atomic
	helm install nginx ingress-nginx/ingress-nginx -f nginx-ingress.yaml --atomic

start-user-service:
	kubectl config set-context --current --namespace=user
	helm install userrelease ./user-chart/ -n user

load-test:
	go run ./load_testing/
