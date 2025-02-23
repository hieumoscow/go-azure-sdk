
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/failovergroups` Documentation

The `failovergroups` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/failovergroups"
```


### Client Initialization

```go
client := failovergroups.NewFailoverGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FailoverGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := failovergroups.NewFailoverGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "failoverGroupValue")

payload := failovergroups.FailoverGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FailoverGroupsClient.Delete`

```go
ctx := context.TODO()
id := failovergroups.NewFailoverGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "failoverGroupValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FailoverGroupsClient.Failover`

```go
ctx := context.TODO()
id := failovergroups.NewFailoverGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "failoverGroupValue")

if err := client.FailoverThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FailoverGroupsClient.ForceFailoverAllowDataLoss`

```go
ctx := context.TODO()
id := failovergroups.NewFailoverGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "failoverGroupValue")

if err := client.ForceFailoverAllowDataLossThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FailoverGroupsClient.Get`

```go
ctx := context.TODO()
id := failovergroups.NewFailoverGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "failoverGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FailoverGroupsClient.ListByServer`

```go
ctx := context.TODO()
id := failovergroups.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FailoverGroupsClient.Update`

```go
ctx := context.TODO()
id := failovergroups.NewFailoverGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "failoverGroupValue")

payload := failovergroups.FailoverGroupUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
