# Terraform Configuration for EKS Cluster

**Note: AI generated content**

This Terraform configuration deploys an EKS cluster in AWS. The infrastructure
includes a VPC with 2 public and 2 private subnets, and an EKS cluster deployed
to that VPC. The EKS cluster has 2 nodes of type `m6a.large` in the private
subnets.

## Infrastructure

- **VPC**: The VPC is configured with a CIDR block of `10.0.0.0/16`. It has 2
  public and 2 private subnets.
- **EKS Cluster**: The EKS cluster is deployed within the VPC. It has 2 nodes of
  type `m6a.large` in the private subnets.

## AWS Authentication

Terraform uses the AWS provider to interact with AWS. The provider needs to be
configured with the proper credentials to authenticate with AWS.

The easiest way to authenticate is to set the `AWS_PROFILE` environment variable
to the name of the profile you want to use. This will instruct the provider to
use the credentials associated with that profile in your AWS credentials file.

For example, if you have a profile named `myprofile` in your AWS credentials
file, you can select it by setting the `AWS_PROFILE` environment variable:

```sh
export AWS_PROFILE=myprofile
```

Then, when you run `terraform apply`, the AWS provider will use the credentials
from the myprofile profile.

Please note that you need to have the AWS CLI installed and configured on your
machine to use this method of authentication.

## Terraform Commands

Before you can apply this configuration, you need to initialize your Terraform
workspace, which will download the provider plugins for AWS:

```sh
terraform init
```

To see what changes Terraform will apply:

```sh
terraform plan
```

To apply the changes:

```sh
terraform apply
```

To destroy the resources when you're done:

```sh
terraform destroy
```
