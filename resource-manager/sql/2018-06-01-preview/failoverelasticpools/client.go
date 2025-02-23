package failoverelasticpools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailoverElasticPoolsClient struct {
	Client *resourcemanager.Client
}

func NewFailoverElasticPoolsClientWithBaseURI(api environments.Api) (*FailoverElasticPoolsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "failoverelasticpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FailoverElasticPoolsClient: %+v", err)
	}

	return &FailoverElasticPoolsClient{
		Client: client,
	}, nil
}
