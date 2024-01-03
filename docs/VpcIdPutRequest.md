# VpcIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip4public** | **string** | Public IPv4 address of the member | 
**Ip4private** | Pointer to **string** | Private IPv4 address of the member | [optional] 

## Methods

### NewVpcIdPutRequest

`func NewVpcIdPutRequest(ip4public string, ) *VpcIdPutRequest`

NewVpcIdPutRequest instantiates a new VpcIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcIdPutRequestWithDefaults

`func NewVpcIdPutRequestWithDefaults() *VpcIdPutRequest`

NewVpcIdPutRequestWithDefaults instantiates a new VpcIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp4public

`func (o *VpcIdPutRequest) GetIp4public() string`

GetIp4public returns the Ip4public field if non-nil, zero value otherwise.

### GetIp4publicOk

`func (o *VpcIdPutRequest) GetIp4publicOk() (*string, bool)`

GetIp4publicOk returns a tuple with the Ip4public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4public

`func (o *VpcIdPutRequest) SetIp4public(v string)`

SetIp4public sets Ip4public field to given value.


### GetIp4private

`func (o *VpcIdPutRequest) GetIp4private() string`

GetIp4private returns the Ip4private field if non-nil, zero value otherwise.

### GetIp4privateOk

`func (o *VpcIdPutRequest) GetIp4privateOk() (*string, bool)`

GetIp4privateOk returns a tuple with the Ip4private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4private

`func (o *VpcIdPutRequest) SetIp4private(v string)`

SetIp4private sets Ip4private field to given value.

### HasIp4private

`func (o *VpcIdPutRequest) HasIp4private() bool`

HasIp4private returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


