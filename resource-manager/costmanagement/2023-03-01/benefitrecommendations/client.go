package benefitrecommendations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitRecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewBenefitRecommendationsClientWithBaseURI(api environments.Api) (*BenefitRecommendationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "benefitrecommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BenefitRecommendationsClient: %+v", err)
	}

	return &BenefitRecommendationsClient{
		Client: client,
	}, nil
}
