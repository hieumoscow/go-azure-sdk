
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/databaseadvisors` Documentation

The `databaseadvisors` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/databaseadvisors"
```


### Client Initialization

```go
client := databaseadvisors.NewDatabaseAdvisorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseAdvisorsClient.Get`

```go
ctx := context.TODO()
id := databaseadvisors.NewDatabaseAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "advisorValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseAdvisorsClient.ListByDatabase`

```go
ctx := context.TODO()
id := databaseadvisors.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.ListByDatabase(ctx, id, databaseadvisors.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseAdvisorsClient.Update`

```go
ctx := context.TODO()
id := databaseadvisors.NewDatabaseAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "advisorValue")

payload := databaseadvisors.Advisor{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
