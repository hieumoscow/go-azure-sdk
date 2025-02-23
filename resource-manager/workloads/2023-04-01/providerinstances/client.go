package providerinstances

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderInstancesClient struct {
	Client *resourcemanager.Client
}

func NewProviderInstancesClientWithBaseURI(api environments.Api) (*ProviderInstancesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "providerinstances", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProviderInstancesClient: %+v", err)
	}

	return &ProviderInstancesClient{
		Client: client,
	}, nil
}
