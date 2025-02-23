
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/syncgroups` Documentation

The `syncgroups` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/syncgroups"
```


### Client Initialization

```go
client := syncgroups.NewSyncGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncGroupsClient.CancelSync`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

read, err := client.CancelSync(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

payload := syncgroups.SyncGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SyncGroupsClient.Delete`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncGroupsClient.Get`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.ListByDatabase`

```go
ctx := context.TODO()
id := syncgroups.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.ListHubSchemas`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

// alternatively `client.ListHubSchemas(ctx, id)` can be used to do batched pagination
items, err := client.ListHubSchemasComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.ListLogs`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

// alternatively `client.ListLogs(ctx, id, syncgroups.DefaultListLogsOperationOptions())` can be used to do batched pagination
items, err := client.ListLogsComplete(ctx, id, syncgroups.DefaultListLogsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.ListSyncDatabaseIds`

```go
ctx := context.TODO()
id := syncgroups.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListSyncDatabaseIds(ctx, id)` can be used to do batched pagination
items, err := client.ListSyncDatabaseIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.RefreshHubSchema`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

if err := client.RefreshHubSchemaThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncGroupsClient.TriggerSync`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

read, err := client.TriggerSync(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.Update`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

payload := syncgroups.SyncGroup{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
