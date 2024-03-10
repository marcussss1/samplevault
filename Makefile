start: |
	cd samplevault
	eval $(minikube docker-env)
	docker build -t "samplevault:latest" -f "Dockerfile" .
	kubectl delete deploy samplevault
	kubectl delete service samplevault
	kubectl delete ingress samplevault
	kubectl apply -f deployment.yml
	kubectl apply -f service.yml
	kubectl apply -f ingress.yml
	eval $(minikube docker-env -u)
