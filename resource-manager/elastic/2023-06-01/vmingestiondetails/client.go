package vmingestiondetails

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMIngestionDetailsClient struct {
	Client *resourcemanager.Client
}

func NewVMIngestionDetailsClientWithBaseURI(api environments.Api) (*VMIngestionDetailsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "vmingestiondetails", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VMIngestionDetailsClient: %+v", err)
	}

	return &VMIngestionDetailsClient{
		Client: client,
	}, nil
}
