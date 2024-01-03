# ClusterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster name | [optional] 
**Location** | Pointer to **string** | Cluster location | [optional] 
**Size** | Pointer to **int32** | Number of initial worker nodes | [optional] 
**Cni** | Pointer to **string** | Networking plugin | [optional] 

## Methods

### NewClusterPostRequest

`func NewClusterPostRequest() *ClusterPostRequest`

NewClusterPostRequest instantiates a new ClusterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPostRequestWithDefaults

`func NewClusterPostRequestWithDefaults() *ClusterPostRequest`

NewClusterPostRequestWithDefaults instantiates a new ClusterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *ClusterPostRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterPostRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterPostRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClusterPostRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSize

`func (o *ClusterPostRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ClusterPostRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ClusterPostRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ClusterPostRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCni

`func (o *ClusterPostRequest) GetCni() string`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *ClusterPostRequest) GetCniOk() (*string, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *ClusterPostRequest) SetCni(v string)`

SetCni sets Cni field to given value.

### HasCni

`func (o *ClusterPostRequest) HasCni() bool`

HasCni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


