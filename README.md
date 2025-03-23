# kube-mcp

## Overview

kube-mcp is a Kubernetes MCP server that provides a set of tools for managing and interacting with Kubernetes clusters.

## Supported Commands

The following tool commands are supported by kube-mcp:

| Command | Description |
|---------|-------------|
| `k8s_list_pods` | List pods in the specified namespace |
| `k8s_list_deployments` | List deployments in the specified namespace |
| `k8s_list_ingresses` | List ingresses in the specified namespace |
| `k8s_list_namespaces` | List all namespaces in the cluster |
| `k8s_list_statefulsets` | List StatefulSets in the specified namespace |


## Install

```bash
go install github.com/orvice/kube-mcp/cmd/kubemcp@latest
```
