package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolPerformanceLevelCapability struct {
	IncludedMaxSize                          *MaxSizeCapability                                     `json:"includedMaxSize,omitempty"`
	MaxDatabaseCount                         *int64                                                 `json:"maxDatabaseCount,omitempty"`
	PerformanceLevel                         *PerformanceLevelCapability                            `json:"performanceLevel,omitempty"`
	Reason                                   *string                                                `json:"reason,omitempty"`
	Sku                                      *Sku                                                   `json:"sku,omitempty"`
	Status                                   *CapabilityStatus                                      `json:"status,omitempty"`
	SupportedLicenseTypes                    *[]LicenseTypeCapability                               `json:"supportedLicenseTypes,omitempty"`
	SupportedMaxSizes                        *[]MaxSizeRangeCapability                              `json:"supportedMaxSizes,omitempty"`
	SupportedPerDatabaseMaxPerformanceLevels *[]ElasticPoolPerDatabaseMaxPerformanceLevelCapability `json:"supportedPerDatabaseMaxPerformanceLevels,omitempty"`
	SupportedPerDatabaseMaxSizes             *[]MaxSizeRangeCapability                              `json:"supportedPerDatabaseMaxSizes,omitempty"`
	ZoneRedundant                            *bool                                                  `json:"zoneRedundant,omitempty"`
}
