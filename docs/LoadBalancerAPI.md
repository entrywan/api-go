# \LoadBalancerAPI

All URIs are relative to *https://api.entrywan.com/v1beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadbalancerGet**](LoadBalancerAPI.md#LoadbalancerGet) | **Get** /loadbalancer | List load balancers
[**LoadbalancerIdDelete**](LoadBalancerAPI.md#LoadbalancerIdDelete) | **Delete** /loadbalancer/{id} | Delete load balancer
[**LoadbalancerIdPut**](LoadBalancerAPI.md#LoadbalancerIdPut) | **Put** /loadbalancer/{id} | Update load balancer targets
[**LoadbalancerPost**](LoadBalancerAPI.md#LoadbalancerPost) | **Post** /loadbalancer | Create a load balancer



## LoadbalancerGet

> []LoadBalancer LoadbalancerGet(ctx).Execute()

List load balancers

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
	resp, r, err := apiClient.LoadBalancerAPI.LoadbalancerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.LoadbalancerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadbalancerGet`: []LoadBalancer
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.LoadbalancerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoadbalancerGetRequest struct via the builder pattern


### Return type

[**[]LoadBalancer**](LoadBalancer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadbalancerIdDelete

> LoadbalancerIdDelete(ctx, id).Execute()

Delete load balancer

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerAPI.LoadbalancerIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.LoadbalancerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadbalancerIdDeleteRequest struct via the builder pattern


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


## LoadbalancerIdPut

> LoadbalancerIdPut(ctx, id).Execute()

Update load balancer targets

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Load balancer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerAPI.LoadbalancerIdPut(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.LoadbalancerIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Load balancer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadbalancerIdPutRequest struct via the builder pattern


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


## LoadbalancerPost

> LoadbalancerPost(ctx).LoadbalancerPostRequest(loadbalancerPostRequest).Execute()

Create a load balancer

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
	loadbalancerPostRequest := *openapiclient.NewLoadbalancerPostRequest() // LoadbalancerPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerAPI.LoadbalancerPost(context.Background()).LoadbalancerPostRequest(loadbalancerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.LoadbalancerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadbalancerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loadbalancerPostRequest** | [**LoadbalancerPostRequest**](LoadbalancerPostRequest.md) |  | 

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

