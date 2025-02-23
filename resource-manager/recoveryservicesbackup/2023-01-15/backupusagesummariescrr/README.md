
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupusagesummariescrr` Documentation

The `backupusagesummariescrr` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-01-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupusagesummariescrr"
```


### Client Initialization

```go
client := backupusagesummariescrr.NewBackupUsageSummariesCRRClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupUsageSummariesCRRClient.BackupUsageSummariesCRRList`

```go
ctx := context.TODO()
id := backupusagesummariescrr.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

read, err := client.BackupUsageSummariesCRRList(ctx, id, backupusagesummariescrr.DefaultBackupUsageSummariesCRRListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
