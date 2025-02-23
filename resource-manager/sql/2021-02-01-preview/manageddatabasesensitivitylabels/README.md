
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesensitivitylabels` Documentation

The `manageddatabasesensitivitylabels` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesensitivitylabels"
```


### Client Initialization

```go
client := manageddatabasesensitivitylabels.NewManagedDatabaseSensitivityLabelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

payload := manageddatabasesensitivitylabels.SensitivityLabel{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.Delete`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.DisableRecommendation`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.DisableRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.EnableRecommendation`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.EnableRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.Get`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ListCurrentByDatabase`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

// alternatively `client.ListCurrentByDatabase(ctx, id, manageddatabasesensitivitylabels.DefaultListCurrentByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListCurrentByDatabaseComplete(ctx, id, manageddatabasesensitivitylabels.DefaultListCurrentByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ListRecommendedByDatabase`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

// alternatively `client.ListRecommendedByDatabase(ctx, id, manageddatabasesensitivitylabels.DefaultListRecommendedByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendedByDatabaseComplete(ctx, id, manageddatabasesensitivitylabels.DefaultListRecommendedByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.ManagedDatabaseRecommendedSensitivityLabelsUpdate`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

payload := manageddatabasesensitivitylabels.RecommendedSensitivityLabelUpdateList{
	// ...
}


read, err := client.ManagedDatabaseRecommendedSensitivityLabelsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSensitivityLabelsClient.Update`

```go
ctx := context.TODO()
id := manageddatabasesensitivitylabels.NewManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

payload := manageddatabasesensitivitylabels.SensitivityLabelUpdateList{
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
