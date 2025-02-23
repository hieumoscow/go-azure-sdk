package managedinstances

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceProperties struct {
	AdministratorLogin         *string                       `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string                       `json:"administratorLoginPassword,omitempty"`
	Collation                  *string                       `json:"collation,omitempty"`
	DnsZone                    *string                       `json:"dnsZone,omitempty"`
	DnsZonePartner             *string                       `json:"dnsZonePartner,omitempty"`
	FullyQualifiedDomainName   *string                       `json:"fullyQualifiedDomainName,omitempty"`
	InstancePoolId             *string                       `json:"instancePoolId,omitempty"`
	LicenseType                *ManagedInstanceLicenseType   `json:"licenseType,omitempty"`
	MaintenanceConfigurationId *string                       `json:"maintenanceConfigurationId,omitempty"`
	ManagedInstanceCreateMode  *ManagedServerCreateMode      `json:"managedInstanceCreateMode,omitempty"`
	MinimalTlsVersion          *string                       `json:"minimalTlsVersion,omitempty"`
	ProxyOverride              *ManagedInstanceProxyOverride `json:"proxyOverride,omitempty"`
	PublicDataEndpointEnabled  *bool                         `json:"publicDataEndpointEnabled,omitempty"`
	RestorePointInTime         *string                       `json:"restorePointInTime,omitempty"`
	SourceManagedInstanceId    *string                       `json:"sourceManagedInstanceId,omitempty"`
	State                      *string                       `json:"state,omitempty"`
	StorageSizeInGB            *int64                        `json:"storageSizeInGB,omitempty"`
	SubnetId                   *string                       `json:"subnetId,omitempty"`
	TimezoneId                 *string                       `json:"timezoneId,omitempty"`
	VCores                     *int64                        `json:"vCores,omitempty"`
}

func (o *ManagedInstanceProperties) GetRestorePointInTimeAsTime() (*time.Time, error) {
	if o.RestorePointInTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RestorePointInTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceProperties) SetRestorePointInTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RestorePointInTime = &formatted
}
