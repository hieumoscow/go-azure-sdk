
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-10-01/workspaceprivatelinkresources` Documentation

The `workspaceprivatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-10-01/workspaceprivatelinkresources"
```


### Client Initialization

```go
client := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspacePrivateLinkResourcesClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := workspaceprivatelinkresources.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.PrivateLinkResourcesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
