# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | Load balancer account owner | [optional] 
**Created** | Pointer to **string** | Load balancer creation date | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** | Number of worker nodes assigned to the cluster | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Apiserver** | Pointer to **string** |  | [optional] 
**Cni** | Pointer to **string** |  | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *Cluster) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Cluster) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Cluster) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Cluster) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreated

`func (o *Cluster) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Cluster) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Cluster) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Cluster) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *Cluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Cluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Cluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Cluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSize

`func (o *Cluster) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Cluster) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Cluster) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Cluster) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *Cluster) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Cluster) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Cluster) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Cluster) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiserver

`func (o *Cluster) GetApiserver() string`

GetApiserver returns the Apiserver field if non-nil, zero value otherwise.

### GetApiserverOk

`func (o *Cluster) GetApiserverOk() (*string, bool)`

GetApiserverOk returns a tuple with the Apiserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiserver

`func (o *Cluster) SetApiserver(v string)`

SetApiserver sets Apiserver field to given value.

### HasApiserver

`func (o *Cluster) HasApiserver() bool`

HasApiserver returns a boolean if a field has been set.

### GetCni

`func (o *Cluster) GetCni() string`

GetCni returns the Cni field if non-nil, zero value otherwise.

### GetCniOk

`func (o *Cluster) GetCniOk() (*string, bool)`

GetCniOk returns a tuple with the Cni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCni

`func (o *Cluster) SetCni(v string)`

SetCni sets Cni field to given value.

### HasCni

`func (o *Cluster) HasCni() bool`

HasCni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


