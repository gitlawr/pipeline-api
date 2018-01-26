package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Pipeline        PipelineOperations
	PipelineHistory PipelineHistoryOperations
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
	client.PipelineHistory = newPipelineHistoryClient(client)

	return client, nil
}
