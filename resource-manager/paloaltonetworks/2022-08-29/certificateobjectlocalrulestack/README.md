
## `github.com/hashicorp/go-azure-sdk/resource-manager/paloaltonetworks/2022-08-29/certificateobjectlocalrulestack` Documentation

The `certificateobjectlocalrulestack` SDK allows for interaction with the Azure Resource Manager Service `paloaltonetworks` (API Version `2022-08-29`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/paloaltonetworks/2022-08-29/certificateobjectlocalrulestack"
```


### Client Initialization

```go
client := certificateobjectlocalrulestack.NewCertificateObjectLocalRulestackClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificateObjectLocalRulestackClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := certificateobjectlocalrulestack.NewLocalRuleStackCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRuleStackValue", "certificateValue")

payload := certificateobjectlocalrulestack.CertificateObjectLocalRulestackResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CertificateObjectLocalRulestackClient.Delete`

```go
ctx := context.TODO()
id := certificateobjectlocalrulestack.NewLocalRuleStackCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRuleStackValue", "certificateValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CertificateObjectLocalRulestackClient.Get`

```go
ctx := context.TODO()
id := certificateobjectlocalrulestack.NewLocalRuleStackCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRuleStackValue", "certificateValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateObjectLocalRulestackClient.ListByLocalRulestacks`

```go
ctx := context.TODO()
id := certificateobjectlocalrulestack.NewLocalRuleStackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRuleStackValue")

// alternatively `client.ListByLocalRulestacks(ctx, id)` can be used to do batched pagination
items, err := client.ListByLocalRulestacksComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
