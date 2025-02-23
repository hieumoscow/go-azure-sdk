package containers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainersClient struct {
	Client *resourcemanager.Client
}

func NewContainersClientWithBaseURI(api environments.Api) (*ContainersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "containers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContainersClient: %+v", err)
	}

	return &ContainersClient{
		Client: client,
	}, nil
}
