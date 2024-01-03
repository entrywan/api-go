# SshkeyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the ssh key | [optional] 
**Pub** | Pointer to **string** | SSH public key | [optional] 

## Methods

### NewSshkeyPostRequest

`func NewSshkeyPostRequest() *SshkeyPostRequest`

NewSshkeyPostRequest instantiates a new SshkeyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshkeyPostRequestWithDefaults

`func NewSshkeyPostRequestWithDefaults() *SshkeyPostRequest`

NewSshkeyPostRequestWithDefaults instantiates a new SshkeyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SshkeyPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshkeyPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshkeyPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SshkeyPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPub

`func (o *SshkeyPostRequest) GetPub() string`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *SshkeyPostRequest) GetPubOk() (*string, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *SshkeyPostRequest) SetPub(v string)`

SetPub sets Pub field to given value.

### HasPub

`func (o *SshkeyPostRequest) HasPub() bool`

HasPub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


