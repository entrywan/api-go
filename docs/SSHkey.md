# SSHkey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | SSH key account owner | [optional] 
**Created** | Pointer to **string** | SSH key creation date | [optional] 
**Name** | Pointer to **string** | SSH key name | [optional] 
**Pub** | Pointer to **string** | SSH public key | [optional] 
**Type** | Pointer to **string** | SSH key algorithm | [optional] 
**Comment** | Pointer to **string** | SSH key description | [optional] 

## Methods

### NewSSHkey

`func NewSSHkey() *SSHkey`

NewSSHkey instantiates a new SSHkey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHkeyWithDefaults

`func NewSSHkeyWithDefaults() *SSHkey`

NewSSHkeyWithDefaults instantiates a new SSHkey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSHkey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSHkey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSHkey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SSHkey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *SSHkey) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SSHkey) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SSHkey) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SSHkey) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreated

`func (o *SSHkey) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SSHkey) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SSHkey) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SSHkey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *SSHkey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSHkey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSHkey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SSHkey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPub

`func (o *SSHkey) GetPub() string`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *SSHkey) GetPubOk() (*string, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *SSHkey) SetPub(v string)`

SetPub sets Pub field to given value.

### HasPub

`func (o *SSHkey) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetType

`func (o *SSHkey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SSHkey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SSHkey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SSHkey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComment

`func (o *SSHkey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SSHkey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SSHkey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SSHkey) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


