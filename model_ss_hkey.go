/*
Entrywan API

Manage Entrywan resources programmatically using the API.  All API requests are authenticated using [IAM tokens](https://entrywan.com/docs#iam).  Tokens can be generated and retrieved from the [portal](https://portal.entrywan.com).  The portal itself is an API client that uses an unrestricted token to access resources for an account.  This documentation is generated using an OpenAPI 3.1.0 [specification](https://spec.openapis.org/oas/latest.html).  More information about OpenAPI can be found on its [site](https://openapis.org).  The current version of [Entrywan's API spec](https://entrywan.com/openapi.yaml) is also available for inspection.  On the left of this page are links to the <i>Endpoints</i> grouped by tag and <i>Schemas</i> the API exposes.  <i>Endpoints</i> are URLs that can be accessed with any HTTP client or device. <i>Schemas</i> are machine-readable data models that represent resources.  To learn more, have a look at the [documentation](https://entrywan.com/docs).  If you have any questions, contact [support](mailto:support@entrywan.com) or your account representative.

API version: v1beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SSHkey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSHkey{}

// SSHkey struct for SSHkey
type SSHkey struct {
	Id *string `json:"id,omitempty"`
	// SSH key account owner
	Account *string `json:"account,omitempty"`
	// SSH key creation date
	Created *string `json:"created,omitempty"`
	// SSH key name
	Name *string `json:"name,omitempty"`
	// SSH public key
	Pub *string `json:"pub,omitempty"`
	// SSH key algorithm
	Type *string `json:"type,omitempty"`
	// SSH key description
	Comment *string `json:"comment,omitempty"`
}

// NewSSHkey instantiates a new SSHkey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHkey() *SSHkey {
	this := SSHkey{}
	return &this
}

// NewSSHkeyWithDefaults instantiates a new SSHkey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHkeyWithDefaults() *SSHkey {
	this := SSHkey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SSHkey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SSHkey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SSHkey) SetId(v string) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SSHkey) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SSHkey) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *SSHkey) SetAccount(v string) {
	o.Account = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SSHkey) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SSHkey) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *SSHkey) SetCreated(v string) {
	o.Created = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SSHkey) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SSHkey) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SSHkey) SetName(v string) {
	o.Name = &v
}

// GetPub returns the Pub field value if set, zero value otherwise.
func (o *SSHkey) GetPub() string {
	if o == nil || IsNil(o.Pub) {
		var ret string
		return ret
	}
	return *o.Pub
}

// GetPubOk returns a tuple with the Pub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetPubOk() (*string, bool) {
	if o == nil || IsNil(o.Pub) {
		return nil, false
	}
	return o.Pub, true
}

// HasPub returns a boolean if a field has been set.
func (o *SSHkey) HasPub() bool {
	if o != nil && !IsNil(o.Pub) {
		return true
	}

	return false
}

// SetPub gets a reference to the given string and assigns it to the Pub field.
func (o *SSHkey) SetPub(v string) {
	o.Pub = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SSHkey) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SSHkey) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SSHkey) SetType(v string) {
	o.Type = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SSHkey) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHkey) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SSHkey) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SSHkey) SetComment(v string) {
	o.Comment = &v
}

func (o SSHkey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSHkey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pub) {
		toSerialize["pub"] = o.Pub
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableSSHkey struct {
	value *SSHkey
	isSet bool
}

func (v NullableSSHkey) Get() *SSHkey {
	return v.value
}

func (v *NullableSSHkey) Set(val *SSHkey) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHkey) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHkey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHkey(val *SSHkey) *NullableSSHkey {
	return &NullableSSHkey{value: val, isSet: true}
}

func (v NullableSSHkey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHkey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

