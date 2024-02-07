# Pulumi Projects Repository

This repository contains two separate projects demonstrating the use of Pulumi to manage infrastructure across different technologies: Kubernetes and AWS. Each project is contained within its own directory and can be used independently.

## Structure

- `/demo-pulumi-k8s`: Contains a Pulumi project that defines infrastructure resources for Kubernetes.
- `/demo-pulumi-aws`: Contains a Pulumi project that defines infrastructure resources for AWS (Amazon Web Services).

## Kubernetes Project

The Kubernetes project showcases how to use Pulumi to orchestrate Kubernetes resources such as Deployments, Services, and ConfigMaps. For more details, see the [README.md](./demo-pulumi-k8s/README.md) in the `/demo-pulumi-k8s` directory.

## AWS Project

The AWS project demonstrates using Pulumi with the AWS provider to create and manage cloud resources like S3 Buckets, EC2 Instances, and more. For more details, see the [README.md](./demo-pulumi-aws/README.md) in the `/demo-pulumi-aws` directory.

## Getting Started

To get started with either project, navigate to the respective directory and follow the instructions in the `README.md` file. Ensure you have Pulumi CLI installed and configured appropriately for the corresponding provider.

## Prerequisites

- Pulumi CLI
- Either one of the following, depending on the stack of choice:
  - Access to a Kubernetes cluster for the Kubernetes project (e.g., minikube, EKS)
  - AWS account and credentials configured for the AWS project

## Contributing

Contributions to either project are welcome! Please follow the guidelines outlined in each project's `README.md` for submitting changes or enhancements.

## License

This repository is licensed under the [MIT License](LICENSE).
