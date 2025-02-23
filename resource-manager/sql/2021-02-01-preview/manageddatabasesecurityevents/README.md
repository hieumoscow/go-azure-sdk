
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesecurityevents` Documentation

The `manageddatabasesecurityevents` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesecurityevents"
```


### Client Initialization

```go
client := manageddatabasesecurityevents.NewManagedDatabaseSecurityEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseSecurityEventsClient.ListByDatabase`

```go
ctx := context.TODO()
id := manageddatabasesecurityevents.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id, manageddatabasesecurityevents.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, manageddatabasesecurityevents.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
