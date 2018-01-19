package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Pipeline     PipelineOperations
	Activity     ActivityOperations
	GitAccount   GitAccountOperations
	GitRepoCache GitRepoCacheOperations
	SCMSetting   SCMSettingOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Pipeline = newPipelineClient(client)
	client.Activity = newActivityClient(client)
	client.GitAccount = newGitAccountClient(client)
	client.GitRepoCache = newGitRepoCacheClient(client)
	client.SCMSetting = newSCMSettingClient(client)

	return client, nil
}
