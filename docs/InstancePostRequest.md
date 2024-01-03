# InstancePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** | Number of CPU cores | 
**Ram** | **int32** | Gigabytes (GB) of memory | 
**Disk** | **int32** | Hard drive size in gigabytes (GB) | 
**Location** | **string** | Data center location | 
**Userdata** | Pointer to **string** | Initial script to be performed on first boot | [optional] 
**Sshkey** | **string** | Name of ssh key to be planted on instance for root user | 
**Os** | Pointer to **string** | Name of operating system image | [optional] 
**Hostname** | Pointer to **string** | Hostname | [optional] 

## Methods

### NewInstancePostRequest

`func NewInstancePostRequest(cpu int32, ram int32, disk int32, location string, sshkey string, ) *InstancePostRequest`

NewInstancePostRequest instantiates a new InstancePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePostRequestWithDefaults

`func NewInstancePostRequestWithDefaults() *InstancePostRequest`

NewInstancePostRequestWithDefaults instantiates a new InstancePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *InstancePostRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstancePostRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstancePostRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *InstancePostRequest) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstancePostRequest) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstancePostRequest) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *InstancePostRequest) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *InstancePostRequest) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *InstancePostRequest) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetLocation

`func (o *InstancePostRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InstancePostRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InstancePostRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetUserdata

`func (o *InstancePostRequest) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *InstancePostRequest) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *InstancePostRequest) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *InstancePostRequest) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetSshkey

`func (o *InstancePostRequest) GetSshkey() string`

GetSshkey returns the Sshkey field if non-nil, zero value otherwise.

### GetSshkeyOk

`func (o *InstancePostRequest) GetSshkeyOk() (*string, bool)`

GetSshkeyOk returns a tuple with the Sshkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkey

`func (o *InstancePostRequest) SetSshkey(v string)`

SetSshkey sets Sshkey field to given value.


### GetOs

`func (o *InstancePostRequest) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstancePostRequest) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstancePostRequest) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *InstancePostRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetHostname

`func (o *InstancePostRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InstancePostRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InstancePostRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InstancePostRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


