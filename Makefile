up_all:
	minikube start --memory=2048 --cpus=2 --disk-size=2g
	minikube addons enable ingress
	kubectl apply -f deployment.yml
	kubectl apply -f service.yml
	kubectl apply -f ingress.yml
	kubectl create cm --from-file init.lua tarantool-secrets
	kubectl apply -f statefulset.yml
	kubectl apply -f service.yml

delete_all:
	minikube delete --all=true --purge=true

lint:
	smartimports .
	golangci-lint run -v --color=always --timeout 4m

clean_cache:
	docker builder prune
