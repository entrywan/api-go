# LoadbalancerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the load balancer | [optional] 
**Location** | Pointer to **string** | Location of the load balancer | [optional] 
**Listeners** | Pointer to [**[]LoadbalancerPostRequestListenersInner**](LoadbalancerPostRequestListenersInner.md) |  | [optional] 

## Methods

### NewLoadbalancerPostRequest

`func NewLoadbalancerPostRequest() *LoadbalancerPostRequest`

NewLoadbalancerPostRequest instantiates a new LoadbalancerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerPostRequestWithDefaults

`func NewLoadbalancerPostRequestWithDefaults() *LoadbalancerPostRequest`

NewLoadbalancerPostRequestWithDefaults instantiates a new LoadbalancerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadbalancerPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadbalancerPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadbalancerPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadbalancerPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *LoadbalancerPostRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoadbalancerPostRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoadbalancerPostRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LoadbalancerPostRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetListeners

`func (o *LoadbalancerPostRequest) GetListeners() []LoadbalancerPostRequestListenersInner`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LoadbalancerPostRequest) GetListenersOk() (*[]LoadbalancerPostRequestListenersInner, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *LoadbalancerPostRequest) SetListeners(v []LoadbalancerPostRequestListenersInner)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *LoadbalancerPostRequest) HasListeners() bool`

HasListeners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


