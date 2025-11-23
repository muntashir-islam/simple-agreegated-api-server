# 📦 Simple Aggregated API Server

A minimal Kubernetes **Aggregated API (AA) extension** with a custom API group.

This project is a straightforward demonstration of how to build a **custom Kubernetes API** using the **Aggregated API Server** model. This is the standard way to extend the Kubernetes API without modifying the core `$kube-apiserver$` itself.

Instead of embedding logic directly inside the Kubernetes API server, we run our own server behind an $APIService$ resource. Kubernetes automatically routes specific requests—like the example below—to our custom Go server, completely bypassing the built-in API server for that resource.

> `kubectl get gadgets.example.com` will be routed to our custom Go API server, not the built-in API server.

This mechanism is the same one used by key components in the Kubernetes ecosystem:
* **Metrics Server**
* **Kube Aggregator**
* **Service Catalog**
* **Cluster API**

---

## 🚀 Features

* **Custom API Group:** `gadgets.example.com`
* **API Version:** `v1`
* **Resource Kind:** `Gadget`
* **Secure Serving:** Uses **HTTPS with TLS** to secure communication between the $kube-apiserver$ and our custom API server.
* **Valid Integration:** Fully compliant Kubernetes $APIService$ integration.
* **Local Setup:** Uses **kind** for an easy, fully local cluster deployment.
* **Fast Builds:** Multi-stage Docker build for creating small, production-ready images quickly.
* **Minimal Server:** Example Go server implementation using `net/http` for simple JSON responses.

---

## 🛠️ Build & Run

These instructions assume you have **Docker**, **kind**, and **kubectl** installed.

### 1. Build Docker image
This step compiles the Go server and creates the Docker image.

```bash
make build
```
### 2. Load the image into your local kind cluster
This makes the custom API server image available to the Kubernetes deployment.
```bash
make kind-load
```
### 3. Deploy all Kubernetes components
This step deploys the required resources in the correct order.
```bash
make deploy
```
This applies the following resources from the config/ directory:

* **CRD** (CustomResourceDefinition)
* **TLS secret** (containing the certificate)
* **Deployment** (the custom Go API server)
* **Service** (internal cluster access)
* **APIService** (the resource that links the API server to Kubernetes)

## 🔍 Verify the API is Registered

The following commands confirm that your Aggregated API server has been successfully registered and integrated with the main Kubernetes API server.

---

### Check the APIService Status

Verify that the `$kube-apiserver$` has successfully registered and connected to your custom server. The presence of **`True`** in the output confirms a successful connection.

**Command:**

```bash
kubectl get apiservices | grep gadget
```
**Expected Output:**

```bash
v1.gadgets.example.com   Local   True
```
Check for the Custom Resource
Ensure the new resource kind (`Gadget`) is available and recognized by the `$kubectl$` client.

```bash
kubectl api-resources | grep gadget
```
Expected Output:
```bash
gadgets                             gd           muntashirislam.com/v1             true         Gadget

```

### 🧰 Development

**Generate TLS Certificates (If needed)**
The config/secret-tls.yaml should be pre-configured, but if you need to regenerate the TLS secret, use this process. The Common Name (CN) is critical and must match the service name: gadget-api.default.svc.

```bash
openssl req -x509 -newkey rsa:2048 \
  -keyout tls.key -out tls.crt -days 365 -nodes \
  -subj "/CN=gadget-api.default.svc"

kubectl create secret tls gadget-server-tls \
  --cert=tls.crt --key=tls.key
```

###🏗️ Docker Build Notes
The included `Dockerfile` uses a multi-stage build to ensure a lean final image:

Stage 1: Builds the static Go binary (targeting `linux/amd64` by default).

Stage 2: Copies the binary into a minimal Alpine image for a small runtime footprint.

To build for **Apple Silicon / ARM64** architectures, specify the platform flag:
```bash
docker build --platform=linux/arm64 -t gadget-api:latest .
```

### 🧹 Cleanup
To remove all deployed resources from your cluster:

```bash
make undeploy
```
