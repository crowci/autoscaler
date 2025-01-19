package hetznercloud

import (
	"context"
	"testing"
	"text/template"

	"github.com/crowci/autoscaler/config"
	"github.com/crowci/autoscaler/providers/hetznercloud/hcapi/mocks"
	crow "github.com/crowci/crow/v3/crow-go/crow"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeployAgent(t *testing.T) {
	tests := []struct {
		name          string
		setupMocks    func(*mocks.MockClient)
		userdata      string
		sshkeys       []string
		expectedError string
	}{
		{
			name:          "InvalidUserData",
			setupMocks:    func(_ *mocks.MockClient) {},
			userdata:      "{{.InvalidField}}",
			expectedError: "RenderUserDataTemplate",
		},
		{
			name: "ServerTypeNotFound",
			setupMocks: func(mockClient *mocks.MockClient) {
				mockServerTypeClient := mocks.NewMockServerTypeClient(t)
				mockServerTypeClient.On("GetByName", mock.Anything, mock.Anything).Return(nil, nil, nil)
				mockClient.On("ServerType").Return(mockServerTypeClient)
			},
			expectedError: ErrServerTypeNotFound.Error(),
		},
		{
			name: "ImageNotFound",
			setupMocks: func(mockClient *mocks.MockClient) {
				mockServerType := &hcloud.ServerType{Architecture: "amd64"}
				mockServerTypeClient := mocks.NewMockServerTypeClient(t)
				mockServerTypeClient.On("GetByName", mock.Anything, mock.Anything).Return(mockServerType, nil, nil)
				mockClient.On("ServerType").Return(mockServerTypeClient)

				mockImageClient := mocks.NewMockImageClient(t)
				mockImageClient.On("GetByNameAndArchitecture", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil, nil)
				mockClient.On("Image").Return(mockImageClient)
			},
			expectedError: ErrImageNotFound.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.NewMockClient(t)
			tt.setupMocks(mockClient)

			provider := &Provider{
				client:   mockClient,
				config:   &config.Config{},
				userData: template.Must(template.New("").Parse(tt.userdata)),
				sshKeys:  tt.sshkeys,
			}

			agent := &crow.Agent{}
			err := provider.DeployAgent(context.Background(), agent)
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
