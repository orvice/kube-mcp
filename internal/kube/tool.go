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
	}
)
