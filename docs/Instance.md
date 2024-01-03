# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | Instance state | [optional] 
**Hostname** | Pointer to **string** | Instance hostname | [optional] 
**Account** | Pointer to **string** | Instance account owner | [optional] 
**Ip4** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** | Instance datacenter location | [optional] 
**Cpus** | Pointer to **int** | Instance CPU cores | [optional] 
**Ram** | Pointer to **int** | Instance memory in gigabytes (GB) | [optional] 
**Disk** | Pointer to **int** | Instance hard drive size in gigabytes (GB) | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Instance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Instance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetHostname

`func (o *Instance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Instance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Instance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Instance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAccount

`func (o *Instance) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Instance) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Instance) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Instance) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIp4

`func (o *Instance) GetIp4() string`

GetIp4 returns the Ip4 field if non-nil, zero value otherwise.

### GetIp4Ok

`func (o *Instance) GetIp4Ok() (*string, bool)`

GetIp4Ok returns a tuple with the Ip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4

`func (o *Instance) SetIp4(v string)`

SetIp4 sets Ip4 field to given value.

### HasIp4

`func (o *Instance) HasIp4() bool`

HasIp4 returns a boolean if a field has been set.

### GetLocation

`func (o *Instance) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Instance) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Instance) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Instance) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCpus

`func (o *Instance) GetCpus() int`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *Instance) GetCpusOk() (*int, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *Instance) SetCpus(v int)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *Instance) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetRam

`func (o *Instance) GetRam() int`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *Instance) GetRamOk() (*int, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *Instance) SetRam(v int)`

SetRam sets Ram field to given value.

### HasRam

`func (o *Instance) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetDisk

`func (o *Instance) GetDisk() int`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Instance) GetDiskOk() (*int, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Instance) SetDisk(v int)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Instance) HasDisk() bool`

HasDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


