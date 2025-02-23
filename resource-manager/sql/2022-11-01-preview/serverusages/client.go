package serverusages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerUsagesClient struct {
	Client *resourcemanager.Client
}

func NewServerUsagesClientWithBaseURI(api environments.Api) (*ServerUsagesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serverusages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerUsagesClient: %+v", err)
	}

	return &ServerUsagesClient{
		Client: client,
	}, nil
}
