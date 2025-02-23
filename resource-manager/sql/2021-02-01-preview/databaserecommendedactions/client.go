package databaserecommendedactions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseRecommendedActionsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseRecommendedActionsClientWithBaseURI(api environments.Api) (*DatabaseRecommendedActionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "databaserecommendedactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseRecommendedActionsClient: %+v", err)
	}

	return &DatabaseRecommendedActionsClient{
		Client: client,
	}, nil
}
