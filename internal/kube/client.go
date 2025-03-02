package kube

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	clientset *kubernetes.Clientset
)

func InitClient() {
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		InitInClusterClient()
	} else {
		InitLocalClient()
	}
}

func InitInClusterClient() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func InitLocalClient() {
	var kubeconfig string

	// Try to use kubeconfig from environment variable first
	if kubeconfigEnv := os.Getenv("KUBECONFIG"); kubeconfigEnv != "" {
		kubeconfig = kubeconfigEnv
	} else {
		// Fallback to default location in home directory
		home := homedir.HomeDir()
		if home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			panic("Failed to locate kubeconfig: home directory not found and KUBECONFIG environment variable not set")
		}
	}

	// Build configuration from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
