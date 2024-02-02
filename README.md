## Getting Started with the Project

This guide will walk you through setting up a local Kubernetes environment using Minikube and deploying Kubernetes resources with Pulumi.

### Prerequisites
- Minikube
- Pulumi
- kubectl

### **Step 1: Install Minikube**
To work with Kubernetes locally, you'll need Minikube. If it's not already installed, download and install Minikube by following the instructions on the [Minikube GitHub page](https://github.com/kubernetes/minikube).

### **Step 2: Start Minikube**
Initialize your local Kubernetes cluster:

```bash
minikube start
```
### **Step 3: Verify Cluster Access**
Check if kubectl is properly configured to communicate with your Minikube cluster:

```bash
kubectl cluster-info
```

You should see information about your cluster, confirming that kubectl is set up correctly.

### **Step 4: Install Pulumi & Connect to Minikube**
Install Pulumi on a Mac using Homebrew:

```bash
brew install pulumi
```

Confirm the installation:

```bash
pulumi version
```

Pulumi uses the same kubeconfig file as kubectl for cluster interaction. This file typically resides at ~/.kube/config.

### **Step 5: Initialize a Pulumi Project for Kubernetes**
Set up a new Pulumi project for Kubernetes management:

```bash
mkdir pulumi-minikube-demo && cd pulumi-minikube-demo
pulumi new kubernetes-go
```

This initializes a new Pulumi project with a Go template for Kubernetes resource management.

### **Step 6: Define Your Kubernetes Resources**
Customize the generated main.go file to specify the Kubernetes resources (pods, services, deployments, etc.) you want to manage.

### **Step 7: Deploy with Pulumi**
Deploy your Kubernetes resources to the Minikube cluster:

```bash
pulumi up
```

You'll be prompted to review and confirm the deployment. This command applies your defined resource configurations to the cluster.

### **Step 8: Verify the Deployment**
Ensure your resources are correctly deployed:

```bash
kubectl get all
```

This lists all resources in the default namespace, including those defined in your Pulumi project.

---

### Deleting resources
If you wish to delete any of the resources previouly deployed via `pulumi up`, simply run the following:

```bash
pulumi destroy
```
Note: the operation needs to be confirmed in the CLI