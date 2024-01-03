# LoadbalancerPostRequestListenersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Targets** | Pointer to [**[]LoadbalancerPostRequestListenersInnerTargetsInner**](LoadbalancerPostRequestListenersInnerTargetsInner.md) |  | [optional] 

## Methods

### NewLoadbalancerPostRequestListenersInner

`func NewLoadbalancerPostRequestListenersInner() *LoadbalancerPostRequestListenersInner`

NewLoadbalancerPostRequestListenersInner instantiates a new LoadbalancerPostRequestListenersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerPostRequestListenersInnerWithDefaults

`func NewLoadbalancerPostRequestListenersInnerWithDefaults() *LoadbalancerPostRequestListenersInner`

NewLoadbalancerPostRequestListenersInnerWithDefaults instantiates a new LoadbalancerPostRequestListenersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadbalancerPostRequestListenersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadbalancerPostRequestListenersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadbalancerPostRequestListenersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadbalancerPostRequestListenersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *LoadbalancerPostRequestListenersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadbalancerPostRequestListenersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadbalancerPostRequestListenersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoadbalancerPostRequestListenersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTargets

`func (o *LoadbalancerPostRequestListenersInner) GetTargets() []LoadbalancerPostRequestListenersInnerTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *LoadbalancerPostRequestListenersInner) GetTargetsOk() (*[]LoadbalancerPostRequestListenersInnerTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *LoadbalancerPostRequestListenersInner) SetTargets(v []LoadbalancerPostRequestListenersInnerTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *LoadbalancerPostRequestListenersInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


