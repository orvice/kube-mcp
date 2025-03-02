package kube

import (
	"context"
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListIngressesArg struct {
	Namespace string
}

func ListIngresses(arg ListIngressesArg) (*mcp_golang.ToolResponse, error) {
	ingresses, err := clientset.NetworkingV1().Ingresses(arg.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var text = ""
	for _, ingress := range ingresses.Items {
		text += fmt.Sprintf("namespace: %s ingress: %s \n", ingress.Namespace, ingress.Name)
		
		// Display hosts and paths
		for _, rule := range ingress.Spec.Rules {
			host := rule.Host
			if host == "" {
				host = "*"
			}
			
			text += fmt.Sprintf("  host: %s\n", host)
			
			if rule.HTTP != nil {
				for _, path := range rule.HTTP.Paths {
					serviceName := path.Backend.Service.Name
					servicePort := path.Backend.Service.Port.Number
					pathType := string(*path.PathType)
					pathValue := path.Path
					if pathValue == "" {
						pathValue = "/"
					}
					
					text += fmt.Sprintf("    path: %s (type: %s) -> %s:%d\n", 
						pathValue, pathType, serviceName, servicePort)
				}
			}
		}
	}
	content := mcp_golang.NewTextContent(text)
	return &mcp_golang.ToolResponse{
		Content: []*mcp_golang.Content{content},
	}, nil
}
