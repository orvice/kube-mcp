package main

import (
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/orvice/kube-mcp/internal/kube"
)

func main() {
	done := make(chan struct{})
	kube.InitClient()
	var err error
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport(), mcp_golang.WithName("kubemcp"))

	// Register tools
	for _, v := range kube.ToolList {
		err = server.RegisterTool(v.Name, v.Description, v.Handler)
		if err != nil {
			panic(err)
		}
	}

	// Start the server after registering all resources and tools
	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done

}
