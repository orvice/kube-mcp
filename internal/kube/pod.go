package kube

import (
	"context"
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ListPodsArg struct {
	Namespace string
}

func ListPods(arg ListPodsArg) (*mcp_golang.ToolResponse, error) {
	pods, err := clientset.CoreV1().Pods(arg.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var text = ""
	for _, pod := range pods.Items {
		text += fmt.Sprintf("namespace: %s pod: %s status:%s \n",
			pod.Namespace, pod.Name, pod.Status.Phase)
	}
	content := mcp_golang.NewTextContent(text)
	return &mcp_golang.ToolResponse{
		Content: []*mcp_golang.Content{content},
	}, nil
}

type DeletePodArg struct {
	Namespace string
	PodName   string
}

func DeletePod(arg DeletePodArg) (*mcp_golang.ToolResponse, error) {
	err := clientset.CoreV1().Pods(arg.Namespace).Delete(context.TODO(), arg.PodName, metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}
	
	text := fmt.Sprintf("Pod %s in namespace %s has been deleted\n", arg.PodName, arg.Namespace)
	content := mcp_golang.NewTextContent(text)
	return &mcp_golang.ToolResponse{
		Content: []*mcp_golang.Content{content},
	}, nil
}
