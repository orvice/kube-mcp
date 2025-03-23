package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/orvice/kube-mcp/internal/kube"
)

const (
	version = "0.0.1"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"KubeMCP",
		version,
	)

	kube.InitClient()
	for _, tool := range kube.ToolList {
		s.AddTool(tool.Tool, tool.Handler)

	}

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
