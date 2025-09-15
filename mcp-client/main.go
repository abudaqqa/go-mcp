package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func callTool(ctx context.Context, session *mcp.ClientSession, toolName, name string) {
	params := &mcp.CallToolParams{
		Name:      toolName,
		Arguments: map[string]any{"name": name},
	}
	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatalf("CallTool failed: %v", err)
	}
	if res.IsError {
		log.Fatal("tool failed")
	}
	for _, c := range res.Content {
		log.Print(c.(*mcp.TextContent).Text)
	}
}

func main() {
	ctx := context.Background()

	// Create a new client, with no features.
	client := mcp.NewClient(&mcp.Implementation{Name: "mcp-client", Version: "v1.0.0"}, nil)

	// Connect to a server over stdin/stdout
	transport := &mcp.CommandTransport{Command: exec.Command("myserver")}
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Call a tool on the server.
	callTool(ctx, session, "greet", "Ahmed")
	callTool(ctx, session, "farewell", "Ahmed")
}
