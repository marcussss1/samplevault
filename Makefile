start: |
	docker build -t "samplevault:$$IMAGE_VERSION" -f "Dockerfile" .
	kubectl set image deployments/samplevault samplevault=samplevault:$$IMAGE_VERSION
	kubectl rollout status deployments/samplevault
	docker rmi -f samplevault:$$IMAGE_VERSION
	eval $(minikube docker-env -u)

all:
	minikube start --memory=2048 --cpus=2 --disk-size=2g
	minikube addons enable ingress
	kubectl apply -f deployment.yml
	kubectl apply -f service.yml
	kubectl apply -f ingress.yml
	kubectl create cm --from-file init.lua app
	kubectl apply -f tarantool-statefulset.yml
	kubectl apply -f tarantool-service.yml
