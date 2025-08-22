package main

import (
	"github.com/autosuggest-client/mcp-server/config"
	"github.com/autosuggest-client/mcp-server/models"
	tools_autosuggest "github.com/autosuggest-client/mcp-server/tools/autosuggest"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_autosuggest.CreateAutosuggestTool(cfg),
	}
}
