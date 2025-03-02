package kube

type Tool struct {
	Name        string
	Description string
	Handler     any
}

var (
	ToolList = []*Tool{
		{
			Name:        "list_pods",
			Description: "List pods in the current namespace",
			Handler:     ListPods,
		},
		{
			Name:        "list_deployments",
			Description: "List deployments in the current namespace",
			Handler:     ListDeployments,
		},
		{
			Name:        "list_ingresses",
			Description: "List ingresses in the current namespace",
			Handler:     ListIngresses,
		},
		{
			Name:        "list_namespaces",
			Description: "List all namespaces in the cluster",
			Handler:     ListNamespaces,
		},
	}
)
