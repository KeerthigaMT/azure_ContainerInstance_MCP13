package main

import (
	"github.com/containerinstancemanagementclient/mcp-server/config"
	"github.com/containerinstancemanagementclient/mcp-server/models"
	tools_subscriptions "github.com/containerinstancemanagementclient/mcp-server/tools/subscriptions"
	tools_operations "github.com/containerinstancemanagementclient/mcp-server/tools/operations"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_subscriptions.CreateContainer_executecommandTool(cfg),
		tools_operations.CreateOperations_listTool(cfg),
		tools_subscriptions.CreateContainer_listlogsTool(cfg),
		tools_subscriptions.CreateContainergroups_restartTool(cfg),
		tools_subscriptions.CreateContainergroups_stopTool(cfg),
		tools_subscriptions.CreateListcachedimagesTool(cfg),
		tools_subscriptions.CreateContainergroups_listbyresourcegroupTool(cfg),
		tools_subscriptions.CreateContainergroups_deleteTool(cfg),
		tools_subscriptions.CreateContainergroups_getTool(cfg),
		tools_subscriptions.CreateContainergroups_updateTool(cfg),
		tools_subscriptions.CreateContainergroups_createorupdateTool(cfg),
		tools_subscriptions.CreateContainergroups_startTool(cfg),
		tools_subscriptions.CreateServiceassociationlink_deleteTool(cfg),
		tools_subscriptions.CreateContainergroups_listTool(cfg),
		tools_subscriptions.CreateListcapabilitiesTool(cfg),
		tools_subscriptions.CreateContainergroupusage_listTool(cfg),
	}
}
