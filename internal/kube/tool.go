package kube

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Tool struct {
	Tool    mcp.Tool
	Handler server.ToolHandlerFunc
}

var (
	ToolList = []*Tool{
		{Tool: mcp.NewTool("k8s_list_deployments",
			mcp.WithDescription("List K8s Deployment"),
			mcp.WithString("namespace",
				mcp.Required(),
				mcp.Description("namespace"),
			),
		),
			Handler: Deployments,
		},
		{Tool: mcp.NewTool("k8s_list_pods",
			mcp.WithDescription("List K8s Pods"),
			mcp.WithString("namespace",
				mcp.Required(),
				mcp.Description("namespace"),
			),
		),
			Handler: Pods,
		},
		{Tool: mcp.NewTool("k8s_list_namespaces",
			mcp.WithDescription("List K8s Namespaces"),
		),
			Handler: Namespaces,
		},
		{Tool: mcp.NewTool("k8s_list_ingresses",
			mcp.WithDescription("List K8s Ingresses"),
			mcp.WithString("namespace",
				mcp.Required(),
				mcp.Description("namespace"),
			),
		),
			Handler: Ingresses,
		},
	}
)
