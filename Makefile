start: |
	docker build -t "samplevault:$$IMAGE_VERSION" -f "Dockerfile" .
	kubectl set image deployments/samplevault samplevault=samplevault:$$IMAGE_VERSION
	kubectl rollout status deployments/samplevault
	docker rmi -f samplevault:$$IMAGE_VERSION
	eval $(minikube docker-env -u)
