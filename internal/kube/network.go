package kube

import (
	"context"
	"errors"

	"github.com/mark3labs/mcp-go/mcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Ingresses(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, ok := request.Params.Arguments["namespace"].(string)
	if !ok {
		return nil, errors.New("namespace must be a string")
	}

	ingresses, err := clientset.NetworkingV1().Ingresses(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var resp string
	if len(ingresses.Items) == 0 {
		resp = "No ingresses found in namespace: " + ns
	} else {
		resp = "Ingresses in namespace " + ns + ":\n"
		for _, ing := range ingresses.Items {
			resp += "- " + ing.Name + "\n"
			// Add hosts and paths
			for _, rule := range ing.Spec.Rules {
				resp += "  Host: " + rule.Host + "\n"
				if rule.HTTP != nil {
					for _, path := range rule.HTTP.Paths {
						resp += "    Path: " + path.Path
						if path.PathType != nil {
							resp += " (Type: " + string(*path.PathType) + ")"
						}
						resp += " -> " + path.Backend.Service.Name + ":" + path.Backend.Service.Port.String() + "\n"
					}
				}
			}
		}
	}

	return mcp.NewToolResultText(resp), nil
}
