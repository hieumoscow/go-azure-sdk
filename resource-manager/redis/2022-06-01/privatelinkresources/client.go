package privatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkResourcesClientWithBaseURI(api environments.Api) (*PrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "privatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkResourcesClient: %+v", err)
	}

	return &PrivateLinkResourcesClient{
		Client: client,
	}, nil
}
