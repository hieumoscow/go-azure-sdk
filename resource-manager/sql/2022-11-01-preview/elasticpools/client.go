package elasticpools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolsClient struct {
	Client *resourcemanager.Client
}

func NewElasticPoolsClientWithBaseURI(api environments.Api) (*ElasticPoolsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "elasticpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ElasticPoolsClient: %+v", err)
	}

	return &ElasticPoolsClient{
		Client: client,
	}, nil
}
