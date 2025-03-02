package kube

import (
	"context"
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListDeploymentsArg struct {
	Namespace string
}

func ListDeployments(arg ListDeploymentsArg) (*mcp_golang.ToolResponse, error) {
	deployments, err := clientset.AppsV1().Deployments(arg.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var text = ""
	for _, deployment := range deployments.Items {
		text += fmt.Sprintf("namespace: %s deployment: %s replicas: %d/%d \n",
			deployment.Namespace, deployment.Name, 
			deployment.Status.ReadyReplicas, *deployment.Spec.Replicas)
	}
	content := mcp_golang.NewTextContent(text)
	return &mcp_golang.ToolResponse{
		Content: []*mcp_golang.Content{content},
	}, nil
}
