package databases

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseProperties struct {
	AutoPauseDelay                    *int64                      `json:"autoPauseDelay,omitempty"`
	AvailabilityZone                  *AvailabilityZoneType       `json:"availabilityZone,omitempty"`
	CatalogCollation                  *CatalogCollationType       `json:"catalogCollation,omitempty"`
	Collation                         *string                     `json:"collation,omitempty"`
	CreateMode                        *CreateMode                 `json:"createMode,omitempty"`
	CreationDate                      *string                     `json:"creationDate,omitempty"`
	CurrentBackupStorageRedundancy    *BackupStorageRedundancy    `json:"currentBackupStorageRedundancy,omitempty"`
	CurrentServiceObjectiveName       *string                     `json:"currentServiceObjectiveName,omitempty"`
	CurrentSku                        *Sku                        `json:"currentSku,omitempty"`
	DatabaseId                        *string                     `json:"databaseId,omitempty"`
	DefaultSecondaryLocation          *string                     `json:"defaultSecondaryLocation,omitempty"`
	EarliestRestoreDate               *string                     `json:"earliestRestoreDate,omitempty"`
	ElasticPoolId                     *string                     `json:"elasticPoolId,omitempty"`
	EncryptionProtector               *string                     `json:"encryptionProtector,omitempty"`
	FailoverGroupId                   *string                     `json:"failoverGroupId,omitempty"`
	FederatedClientId                 *string                     `json:"federatedClientId,omitempty"`
	HighAvailabilityReplicaCount      *int64                      `json:"highAvailabilityReplicaCount,omitempty"`
	IsInfraEncryptionEnabled          *bool                       `json:"isInfraEncryptionEnabled,omitempty"`
	IsLedgerOn                        *bool                       `json:"isLedgerOn,omitempty"`
	Keys                              *map[string]DatabaseKey     `json:"keys,omitempty"`
	LicenseType                       *DatabaseLicenseType        `json:"licenseType,omitempty"`
	LongTermRetentionBackupResourceId *string                     `json:"longTermRetentionBackupResourceId,omitempty"`
	MaintenanceConfigurationId        *string                     `json:"maintenanceConfigurationId,omitempty"`
	ManualCutover                     *bool                       `json:"manualCutover,omitempty"`
	MaxLogSizeBytes                   *int64                      `json:"maxLogSizeBytes,omitempty"`
	MaxSizeBytes                      *int64                      `json:"maxSizeBytes,omitempty"`
	MinCapacity                       *float64                    `json:"minCapacity,omitempty"`
	PausedDate                        *string                     `json:"pausedDate,omitempty"`
	PerformCutover                    *bool                       `json:"performCutover,omitempty"`
	PreferredEnclaveType              *AlwaysEncryptedEnclaveType `json:"preferredEnclaveType,omitempty"`
	ReadScale                         *DatabaseReadScale          `json:"readScale,omitempty"`
	RecoverableDatabaseId             *string                     `json:"recoverableDatabaseId,omitempty"`
	RecoveryServicesRecoveryPointId   *string                     `json:"recoveryServicesRecoveryPointId,omitempty"`
	RequestedBackupStorageRedundancy  *BackupStorageRedundancy    `json:"requestedBackupStorageRedundancy,omitempty"`
	RequestedServiceObjectiveName     *string                     `json:"requestedServiceObjectiveName,omitempty"`
	RestorableDroppedDatabaseId       *string                     `json:"restorableDroppedDatabaseId,omitempty"`
	RestorePointInTime                *string                     `json:"restorePointInTime,omitempty"`
	ResumedDate                       *string                     `json:"resumedDate,omitempty"`
	SampleName                        *SampleName                 `json:"sampleName,omitempty"`
	SecondaryType                     *SecondaryType              `json:"secondaryType,omitempty"`
	SourceDatabaseDeletionDate        *string                     `json:"sourceDatabaseDeletionDate,omitempty"`
	SourceDatabaseId                  *string                     `json:"sourceDatabaseId,omitempty"`
	SourceResourceId                  *string                     `json:"sourceResourceId,omitempty"`
	Status                            *DatabaseStatus             `json:"status,omitempty"`
	ZoneRedundant                     *bool                       `json:"zoneRedundant,omitempty"`
}

func (o *DatabaseProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}

func (o *DatabaseProperties) GetEarliestRestoreDateAsTime() (*time.Time, error) {
	if o.EarliestRestoreDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EarliestRestoreDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetEarliestRestoreDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EarliestRestoreDate = &formatted
}

func (o *DatabaseProperties) GetPausedDateAsTime() (*time.Time, error) {
	if o.PausedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PausedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetPausedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PausedDate = &formatted
}

func (o *DatabaseProperties) GetRestorePointInTimeAsTime() (*time.Time, error) {
	if o.RestorePointInTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RestorePointInTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetRestorePointInTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RestorePointInTime = &formatted
}

func (o *DatabaseProperties) GetResumedDateAsTime() (*time.Time, error) {
	if o.ResumedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ResumedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetResumedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ResumedDate = &formatted
}

func (o *DatabaseProperties) GetSourceDatabaseDeletionDateAsTime() (*time.Time, error) {
	if o.SourceDatabaseDeletionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SourceDatabaseDeletionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetSourceDatabaseDeletionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SourceDatabaseDeletionDate = &formatted
}
