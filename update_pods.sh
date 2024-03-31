#!/bin/bash

image_version=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)
export IMAGE_VERSION="$image_version"
echo "Версия образа: $IMAGE_VERSION"
eval $(minikube docker-env)
docker build -t "samplevault:$$IMAGE_VERSION" -f "Dockerfile" .
kubectl set image deployments/samplevault samplevault=samplevault:$$IMAGE_VERSION
kubectl rollout status deployments/samplevault
eval $(minikube docker-env -u)
