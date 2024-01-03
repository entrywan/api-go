# \VPCAPI

All URIs are relative to *https://api.entrywan.com/v1beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VpcGet**](VPCAPI.md#VpcGet) | **Get** /vpc | List VPC
[**VpcIdDelete**](VPCAPI.md#VpcIdDelete) | **Delete** /vpc/{id} | Delete VPC
[**VpcIdPatch**](VPCAPI.md#VpcIdPatch) | **Patch** /vpc/{id} | Remove member from VPC
[**VpcIdPost**](VPCAPI.md#VpcIdPost) | **Post** /vpc/{id} | Create a VPC
[**VpcIdPut**](VPCAPI.md#VpcIdPut) | **Put** /vpc/{id} | Add member to VPC



## VpcGet

> []VPC VpcGet(ctx).Execute()

List VPC

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.VpcGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.VpcGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VpcGet`: []VPC
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.VpcGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVpcGetRequest struct via the builder pattern


### Return type

[**[]VPC**](VPC.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VpcIdDelete

> VpcIdDelete(ctx, id).Execute()

Delete VPC

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of VPC

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCAPI.VpcIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.VpcIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VPC | 

### Other Parameters

Other parameters are passed through a pointer to a apiVpcIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VpcIdPatch

> VpcIdPatch(ctx, id).VpcIdPatchRequest(vpcIdPatchRequest).Execute()

Remove member from VPC

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | VPC ID
	vpcIdPatchRequest := *openapiclient.NewVpcIdPatchRequest("Ip4private_example") // VpcIdPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCAPI.VpcIdPatch(context.Background(), id).VpcIdPatchRequest(vpcIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.VpcIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | VPC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVpcIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcIdPatchRequest** | [**VpcIdPatchRequest**](VpcIdPatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VpcIdPost

> VpcIdPost(ctx, id).VpcIdPostRequest(vpcIdPostRequest).Execute()

Create a VPC

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | VPC ID
	vpcIdPostRequest := *openapiclient.NewVpcIdPostRequest() // VpcIdPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCAPI.VpcIdPost(context.Background(), id).VpcIdPostRequest(vpcIdPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.VpcIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | VPC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVpcIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcIdPostRequest** | [**VpcIdPostRequest**](VpcIdPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VpcIdPut

> VpcIdPut(ctx, id).VpcIdPutRequest(vpcIdPutRequest).Execute()

Add member to VPC

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | VPC ID
	vpcIdPutRequest := *openapiclient.NewVpcIdPutRequest("Ip4public_example") // VpcIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCAPI.VpcIdPut(context.Background(), id).VpcIdPutRequest(vpcIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.VpcIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | VPC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVpcIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcIdPutRequest** | [**VpcIdPutRequest**](VpcIdPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

