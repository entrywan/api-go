# Sshkey

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

### NewSshkey

`func NewSshkey() *Sshkey`

NewSshkey instantiates a new Sshkey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshkeyWithDefaults

`func NewSshkeyWithDefaults() *Sshkey`

NewSshkeyWithDefaults instantiates a new Sshkey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sshkey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sshkey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sshkey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sshkey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *Sshkey) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Sshkey) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Sshkey) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Sshkey) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreated

`func (o *Sshkey) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Sshkey) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Sshkey) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Sshkey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *Sshkey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sshkey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sshkey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sshkey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPub

`func (o *Sshkey) GetPub() string`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *Sshkey) GetPubOk() (*string, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *Sshkey) SetPub(v string)`

SetPub sets Pub field to given value.

### HasPub

`func (o *Sshkey) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetType

`func (o *Sshkey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Sshkey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Sshkey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Sshkey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComment

`func (o *Sshkey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Sshkey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Sshkey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Sshkey) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


