
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabases` Documentation

The `manageddatabases` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2018-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/manageddatabases"
```


### Client Initialization

```go
client := manageddatabases.NewManagedDatabasesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabasesClient.CompleteRestore`

```go
ctx := context.TODO()
id := manageddatabases.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

payload := manageddatabases.CompleteDatabaseRestoreDefinition{
	// ...
}


if err := client.CompleteRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedDatabasesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := manageddatabases.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

payload := manageddatabases.ManagedDatabase{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedDatabasesClient.Delete`

```go
ctx := context.TODO()
id := manageddatabases.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedDatabasesClient.Get`

```go
ctx := context.TODO()
id := manageddatabases.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabasesClient.ListByInstance`

```go
ctx := context.TODO()
id := manageddatabases.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabasesClient.Update`

```go
ctx := context.TODO()
id := manageddatabases.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

payload := manageddatabases.ManagedDatabaseUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
