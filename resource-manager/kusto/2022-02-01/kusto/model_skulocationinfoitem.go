package kusto

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuLocationInfoItem struct {
	Location string        `json:"location"`
	Zones    *zones.Schema `json:"zones,omitempty"`
}
