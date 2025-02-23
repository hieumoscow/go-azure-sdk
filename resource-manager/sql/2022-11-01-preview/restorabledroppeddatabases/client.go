package restorabledroppeddatabases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedDatabasesClient struct {
	Client *resourcemanager.Client
}

func NewRestorableDroppedDatabasesClientWithBaseURI(api environments.Api) (*RestorableDroppedDatabasesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "restorabledroppeddatabases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RestorableDroppedDatabasesClient: %+v", err)
	}

	return &RestorableDroppedDatabasesClient{
		Client: client,
	}, nil
}
