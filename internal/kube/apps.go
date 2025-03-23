package kube

import (
	"context"
	"errors"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Deployments(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, ok := request.Params.Arguments["namespace"].(string)
	if !ok {
		return nil, errors.New("namespace must be a string")
	}

	deployments, err := clientset.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var resp string
	if len(deployments.Items) == 0 {
		resp = "No deployments found in namespace: " + ns
	} else {
		resp = "Deployments in namespace " + ns + ":\n"
		for _, d := range deployments.Items {
			resp += "- " + d.Name + " (" + strconv.Itoa(int(d.Status.ReadyReplicas)) + "/" +
				strconv.Itoa(int(d.Status.Replicas)) + " ready)\n"
		}
	}

	return mcp.NewToolResultText(resp), nil
}

func StatefulSets(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, ok := request.Params.Arguments["namespace"].(string)
	if !ok {
		return nil, errors.New("namespace must be a string")
	}

	statefulsets, err := clientset.AppsV1().StatefulSets(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var resp string
	if len(statefulsets.Items) == 0 {
		resp = "No StatefulSets found in namespace: " + ns
	} else {
		resp = "StatefulSets in namespace " + ns + ":\n"
		for _, sts := range statefulsets.Items {
			resp += "- " + sts.Name + " (" + strconv.Itoa(int(sts.Status.ReadyReplicas)) + "/" +
				strconv.Itoa(int(sts.Status.Replicas)) + " ready)\n"
		}
	}

	return mcp.NewToolResultText(resp), nil
}
