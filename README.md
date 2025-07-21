# Example Workload

An example workload for Talos Linux and Omni.

## Features

- 📟 **Terminal-friendly output**
- 🔒 **Conforms to default Talos security guidelines**
- 🌐 **Multi-platform** (AMD64/ARM64)

## Quick Start

Deploy using one of the provided example configurations:

### Option 1: Omni Workload Proxy

Use this deployment if you are using Omni and want to expose the service using the [Omni workload proxy](https://omni.siderolabs.com/how-to-guides/expose-an-http-service-from-a-cluster).

```bash
kubectl apply -f https://raw.githubusercontent.com/siderolabs/example-workload/refs/heads/main/example-omni-workload-proxy.yaml
```

### Option 2: LoadBalancer Service

Use this deployment if you have a load balancer controller in your cluster and want to test if it works.

```bash
kubectl apply -f https://raw.githubusercontent.com/siderolabs/example-workload/refs/heads/main/example-svc-lb.yaml
```

### Option 3: NodePort Service

Use this deployment to verify that node port services are working in your cluster.
This is the simplest deployment option with the least amount of requirements.

```bash
kubectl apply -f https://raw.githubusercontent.com/siderolabs/example-workload/refs/heads/main/example-svc-nodeport.yaml
```

## Container Image

The container image is automatically built and published to `ghcr.io/siderolabs/example-workload`.

Built with:

- Non-root user (UID 1000)
- No kernel capabilities
- Read-only root filesystem compatible
- Multi-architecture support
