
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/get` Documentation

The `get` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/get"
```


### Client Initialization

```go
client := get.NewGETClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GETClient.FilesGet`

```go
ctx := context.TODO()
id := get.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

read, err := client.FilesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.FilesList`

```go
ctx := context.TODO()
id := get.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

// alternatively `client.FilesList(ctx, id)` can be used to do batched pagination
items, err := client.FilesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.ProjectsGet`

```go
ctx := context.TODO()
id := get.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

read, err := client.ProjectsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.ProjectsList`

```go
ctx := context.TODO()
id := get.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ProjectsList(ctx, id)` can be used to do batched pagination
items, err := client.ProjectsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.ResourceSkusListSkus`

```go
ctx := context.TODO()
id := get.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ResourceSkusListSkus(ctx, id)` can be used to do batched pagination
items, err := client.ResourceSkusListSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.ServiceTasksGet`

```go
ctx := context.TODO()
id := get.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksGet(ctx, id, get.DefaultServiceTasksGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.ServiceTasksList`

```go
ctx := context.TODO()
id := get.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ServiceTasksList(ctx, id, get.DefaultServiceTasksListOperationOptions())` can be used to do batched pagination
items, err := client.ServiceTasksListComplete(ctx, id, get.DefaultServiceTasksListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.ServicesGet`

```go
ctx := context.TODO()
id := get.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

read, err := client.ServicesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.ServicesList`

```go
ctx := context.TODO()
id := get.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ServicesList(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.ServicesListByResourceGroup`

```go
ctx := context.TODO()
id := get.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "resourceGroupValue")

// alternatively `client.ServicesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.ServicesListSkus`

```go
ctx := context.TODO()
id := get.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ServicesListSkus(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.TasksGet`

```go
ctx := context.TODO()
id := get.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksGet(ctx, id, get.DefaultTasksGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.TasksList`

```go
ctx := context.TODO()
id := get.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

// alternatively `client.TasksList(ctx, id, get.DefaultTasksListOperationOptions())` can be used to do batched pagination
items, err := client.TasksListComplete(ctx, id, get.DefaultTasksListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.UsagesList`

```go
ctx := context.TODO()
id := get.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.UsagesList(ctx, id)` can be used to do batched pagination
items, err := client.UsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
