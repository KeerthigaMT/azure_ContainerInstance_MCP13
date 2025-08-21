package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/containerinstancemanagementclient/mcp-server/config"
	"github.com/containerinstancemanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Container_listlogsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		containerGroupNameVal, ok := args["containerGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: containerGroupName"), nil
		}
		containerGroupName, ok := containerGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: containerGroupName"), nil
		}
		containerNameVal, ok := args["containerName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: containerName"), nil
		}
		containerName, ok := containerNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: containerName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		if val, ok := args["tail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tail=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerInstance/containerGroups/%s/containers/%s/logs%s", cfg.BaseURL, subscriptionId, resourceGroupName, containerGroupName, containerName, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateContainer_listlogsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft.ContainerInstance_containerGroups_containerGroupName_containers_containerName_logs",
		mcp.WithDescription("Get the logs for a specified container instance."),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Client API version")),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("containerGroupName", mcp.Required(), mcp.Description("The name of the container group.")),
		mcp.WithString("containerName", mcp.Required(), mcp.Description("The name of the container instance.")),
		mcp.WithString("tail", mcp.Description("The number of lines to show from the tail of the container instance log. If not provided, all available logs are shown up to 4mb.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Container_listlogsHandler(cfg),
	}
}
