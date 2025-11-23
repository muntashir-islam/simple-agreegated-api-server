IMAGE_REGISTRY=local
IMAGE_NAME=gadget-api
IMAGE_TAG=latest
IMAGE=$(IMAGE_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# -------- Build Image --------
build:
	docker build -t $(IMAGE) .

# -------- Load into kind --------
kind-load: build
	kind load docker-image $(IMAGE)

# -------- Apply manifests --------
deploy:
	kubectl apply -f config/crd.yaml
	kubectl apply -f config/secret-tls.yaml
	kubectl apply -f config/deployment.yaml
	kubectl apply -f config/service.yaml
	kubectl apply -f config/apiservice.yaml

# -------- Delete everything --------
undeploy:
	kubectl delete -f config/apiservice.yaml --ignore-not-found
	kubectl delete -f config/service.yaml --ignore-not-found
	kubectl delete -f config/deployment.yaml --ignore-not-found
	kubectl delete -f config/secret-tls.yaml --ignore-not-found
	kubectl delete -f config/crd.yaml --ignore-not-found

.PHONY: build kind-load deploy undeploy