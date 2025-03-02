package kube

import (
	"context"
	"fmt"
	"time"

	mcp_golang "github.com/metoro-io/mcp-golang"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListNamespacesArg struct{}

func ListNamespaces(arg ListNamespacesArg) (*mcp_golang.ToolResponse, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var text = ""
	for _, namespace := range namespaces.Items {
		age := time.Since(namespace.CreationTimestamp.Time).Round(time.Second)
		status := string(namespace.Status.Phase)
		
		text += fmt.Sprintf("namespace: %s status: %s age: %s\n",
			namespace.Name, status, age)
	}
	content := mcp_golang.NewTextContent(text)
	return &mcp_golang.ToolResponse{
		Content: []*mcp_golang.Content{content},
	}, nil
}
