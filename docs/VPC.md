# VPC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | VPC account owner | [optional] 
**Created** | Pointer to **string** | VPC creation date | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Members** | Pointer to [**[]VPCMembersInner**](VPCMembersInner.md) |  | [optional] 

## Methods

### NewVPC

`func NewVPC() *VPC`

NewVPC instantiates a new VPC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPCWithDefaults

`func NewVPCWithDefaults() *VPC`

NewVPCWithDefaults instantiates a new VPC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPC) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPC) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPC) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPC) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *VPC) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VPC) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VPC) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VPC) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreated

`func (o *VPC) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPC) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPC) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VPC) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *VPC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPC) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VPC) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *VPC) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VPC) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VPC) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VPC) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPort

`func (o *VPC) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VPC) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VPC) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VPC) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMembers

`func (o *VPC) GetMembers() []VPCMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *VPC) GetMembersOk() (*[]VPCMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *VPC) SetMembers(v []VPCMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *VPC) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


