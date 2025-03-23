package kube

import (
	"context"
	"errors"

	"github.com/mark3labs/mcp-go/mcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Pods(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, ok := request.Params.Arguments["namespace"].(string)
	if !ok {
		return nil, errors.New("namespace must be a string")
	}

	pods, err := clientset.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var resp string
	if len(pods.Items) == 0 {
		resp = "No pods found in namespace: " + ns
	} else {
		resp = "Pods in namespace " + ns + ":\n"
		for _, pod := range pods.Items {
			resp += "- " + pod.Name + " (" + string(pod.Status.Phase) + ")\n"
		}
	}

	return mcp.NewToolResultText(resp), nil
}

func Namespaces(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var resp string
	if len(namespaces.Items) == 0 {
		resp = "No namespaces found"
	} else {
		resp = "Namespaces in the cluster:\n"
		for _, ns := range namespaces.Items {
			resp += "- " + ns.Name + "\n"
		}
	}

	return mcp.NewToolResultText(resp), nil
}
