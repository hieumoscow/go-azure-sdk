package saprecommendations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SAPRecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewSAPRecommendationsClientWithBaseURI(api environments.Api) (*SAPRecommendationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "saprecommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SAPRecommendationsClient: %+v", err)
	}

	return &SAPRecommendationsClient{
		Client: client,
	}, nil
}
