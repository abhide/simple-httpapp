CLUSTER=dev01
IMAGE=simple-httpapp
IMAGE_TAG=latest

init:
	kind create cluster --name=${CLUSTER}

fmt:
	go fmt ./

build:
	docker build -t ${IMAGE}:${IMAGE_TAG} ./

kindly-push:
	kind load docker-image ${IMAGE}:${IMAGE_TAG} --name=${CLUSTER}

kindly-deploy:
	kubectl create namespace v1
	kubectl apply -f k8s/simpleapp-v1.yaml -n v1
	kubectl create namespace v2
	kubectl apply -f k8s/simpleapp-v2.yaml -n v2

clean:
	kind delete cluster --name=${CLUSTER}

all: init fmt build kindly-push kindly-deploy