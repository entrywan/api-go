# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | Load balancer account owner | [optional] 
**Created** | Pointer to **string** | Load balancer creation date | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Listeners** | Pointer to [**[]LoadbalancerPostRequestListenersInner**](LoadbalancerPostRequestListenersInner.md) |  | [optional] 

## Methods

### NewLoadBalancer

`func NewLoadBalancer() *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *LoadBalancer) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LoadBalancer) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LoadBalancer) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LoadBalancer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreated

`func (o *LoadBalancer) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoadBalancer) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoadBalancer) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LoadBalancer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *LoadBalancer) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoadBalancer) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoadBalancer) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LoadBalancer) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIp

`func (o *LoadBalancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LoadBalancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LoadBalancer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LoadBalancer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetListeners

`func (o *LoadBalancer) GetListeners() []LoadbalancerPostRequestListenersInner`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LoadBalancer) GetListenersOk() (*[]LoadbalancerPostRequestListenersInner, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *LoadBalancer) SetListeners(v []LoadbalancerPostRequestListenersInner)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *LoadBalancer) HasListeners() bool`

HasListeners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


