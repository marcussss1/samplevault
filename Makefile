generate_image_version:
    @image_version=$$(openssl rand -base64 12); \
    export IMAGE_VERSION=$$image_version;

.PHONY: generate_image_version

start: |
	docker build -t "samplevault:$$IMAGE_VERSION" -f "Dockerfile" .
	kubectl set image deployments/samplevault samplevault=samplevault:$$IMAGE_VERSION
	kubectl rollout status deployments/samplevault
	docker rmi -f samplevault:$$IMAGE_VERSION
	eval $(minikube docker-env -u)
