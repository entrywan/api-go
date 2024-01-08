# LoadbalancerIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listeners** | Pointer to [**[]LoadbalancerPostRequestListenersInner**](LoadbalancerPostRequestListenersInner.md) |  | [optional] 

## Methods

### NewLoadbalancerIdPutRequest

`func NewLoadbalancerIdPutRequest() *LoadbalancerIdPutRequest`

NewLoadbalancerIdPutRequest instantiates a new LoadbalancerIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerIdPutRequestWithDefaults

`func NewLoadbalancerIdPutRequestWithDefaults() *LoadbalancerIdPutRequest`

NewLoadbalancerIdPutRequestWithDefaults instantiates a new LoadbalancerIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListeners

`func (o *LoadbalancerIdPutRequest) GetListeners() []LoadbalancerPostRequestListenersInner`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LoadbalancerIdPutRequest) GetListenersOk() (*[]LoadbalancerPostRequestListenersInner, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *LoadbalancerIdPutRequest) SetListeners(v []LoadbalancerPostRequestListenersInner)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *LoadbalancerIdPutRequest) HasListeners() bool`

HasListeners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


