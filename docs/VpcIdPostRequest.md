# VpcIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the VPC | [optional] 
**Prefix** | Pointer to **string** | CIDR prefix of the VPC | [optional] 

## Methods

### NewVpcIdPostRequest

`func NewVpcIdPostRequest() *VpcIdPostRequest`

NewVpcIdPostRequest instantiates a new VpcIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcIdPostRequestWithDefaults

`func NewVpcIdPostRequestWithDefaults() *VpcIdPostRequest`

NewVpcIdPostRequestWithDefaults instantiates a new VpcIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpcIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpcIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *VpcIdPostRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VpcIdPostRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VpcIdPostRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VpcIdPostRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


