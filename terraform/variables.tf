variable "aws_region" {
  default = "us-east-1"
}

variable "cluster_name" {
  default = "chat-app-cluster"
}

variable "vpc_id" {
  description = "VPC ID for EKS"
}

variable "subnet_ids" {
  description = "List of subnet IDs for EKS worker nodes"
  type        = list(string)
}
