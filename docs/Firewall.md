# Firewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | Firewall account owner | [optional] 
**Created** | Pointer to **string** | Firewall creation date | [optional] 
**Rules** | Pointer to [**FirewallPostRequestRulesInner**](FirewallPostRequestRulesInner.md) |  | [optional] 

## Methods

### NewFirewall

`func NewFirewall() *Firewall`

NewFirewall instantiates a new Firewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallWithDefaults

`func NewFirewallWithDefaults() *Firewall`

NewFirewallWithDefaults instantiates a new Firewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Firewall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Firewall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Firewall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Firewall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *Firewall) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Firewall) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Firewall) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Firewall) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreated

`func (o *Firewall) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Firewall) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Firewall) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Firewall) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRules

`func (o *Firewall) GetRules() FirewallPostRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Firewall) GetRulesOk() (*FirewallPostRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Firewall) SetRules(v FirewallPostRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Firewall) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


