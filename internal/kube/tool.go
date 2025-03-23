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
	}
)
