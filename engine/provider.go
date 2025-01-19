package engine

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/crowci/autoscaler/config"

	crow "github.com/crowci/crow/v3/crow-go/crow"
)

type Provider interface {
	DeployAgent(context.Context, *crow.Agent) error
	RemoveAgent(context.Context, *crow.Agent) error
	ListDeployedAgentNames(context.Context) ([]string, error)
}

// RenderUserDataTemplate renders the user data template for an Agent
// using the provided configuration.
func RenderUserDataTemplate(config *config.Config, agent *crow.Agent, tmpl *template.Template) (string, error) {
	params := struct {
		Image       string
		Environment map[string]string
	}{
		Image: config.Image,
		Environment: map[string]string{
			"CROW_SERVER":        config.GRPCAddress,
			"CROW_AGENT_SECRET":  agent.Token,
			"CROW_MAX_WORKFLOWS": fmt.Sprintf("%d", config.WorkflowsPerAgent),
		},
	}

	if config.GRPCSecure {
		params.Environment["CROW_GRPC_SECURE"] = "true"
	}

	for key, value := range config.Environment {
		params.Environment[key] = value
	}

	var userData bytes.Buffer
	if err := tmpl.Execute(&userData, params); err != nil {
		return "", err
	}

	return userData.String(), nil
}
