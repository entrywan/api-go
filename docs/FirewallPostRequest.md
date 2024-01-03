# FirewallPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the firewall | [optional] 
**Rules** | Pointer to [**[]FirewallPostRequestRulesInner**](FirewallPostRequestRulesInner.md) |  | [optional] 

## Methods

### NewFirewallPostRequest

`func NewFirewallPostRequest() *FirewallPostRequest`

NewFirewallPostRequest instantiates a new FirewallPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallPostRequestWithDefaults

`func NewFirewallPostRequestWithDefaults() *FirewallPostRequest`

NewFirewallPostRequestWithDefaults instantiates a new FirewallPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *FirewallPostRequest) GetRules() []FirewallPostRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallPostRequest) GetRulesOk() (*[]FirewallPostRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallPostRequest) SetRules(v []FirewallPostRequestRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FirewallPostRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


