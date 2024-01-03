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

// checks if the VPCMembersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VPCMembersInner{}

// VPCMembersInner struct for VPCMembersInner
type VPCMembersInner struct {
	Id *string `json:"id,omitempty"`
	Ippublic *string `json:"ippublic,omitempty"`
	Ipprivate *string `json:"ipprivate,omitempty"`
	Pubkey *string `json:"pubkey,omitempty"`
	Privkey *string `json:"privkey,omitempty"`
	Config *string `json:"config,omitempty"`
}

// NewVPCMembersInner instantiates a new VPCMembersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVPCMembersInner() *VPCMembersInner {
	this := VPCMembersInner{}
	return &this
}

// NewVPCMembersInnerWithDefaults instantiates a new VPCMembersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVPCMembersInnerWithDefaults() *VPCMembersInner {
	this := VPCMembersInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VPCMembersInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCMembersInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VPCMembersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VPCMembersInner) SetId(v string) {
	o.Id = &v
}

// GetIppublic returns the Ippublic field value if set, zero value otherwise.
func (o *VPCMembersInner) GetIppublic() string {
	if o == nil || IsNil(o.Ippublic) {
		var ret string
		return ret
	}
	return *o.Ippublic
}

// GetIppublicOk returns a tuple with the Ippublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCMembersInner) GetIppublicOk() (*string, bool) {
	if o == nil || IsNil(o.Ippublic) {
		return nil, false
	}
	return o.Ippublic, true
}

// HasIppublic returns a boolean if a field has been set.
func (o *VPCMembersInner) HasIppublic() bool {
	if o != nil && !IsNil(o.Ippublic) {
		return true
	}

	return false
}

// SetIppublic gets a reference to the given string and assigns it to the Ippublic field.
func (o *VPCMembersInner) SetIppublic(v string) {
	o.Ippublic = &v
}

// GetIpprivate returns the Ipprivate field value if set, zero value otherwise.
func (o *VPCMembersInner) GetIpprivate() string {
	if o == nil || IsNil(o.Ipprivate) {
		var ret string
		return ret
	}
	return *o.Ipprivate
}

// GetIpprivateOk returns a tuple with the Ipprivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCMembersInner) GetIpprivateOk() (*string, bool) {
	if o == nil || IsNil(o.Ipprivate) {
		return nil, false
	}
	return o.Ipprivate, true
}

// HasIpprivate returns a boolean if a field has been set.
func (o *VPCMembersInner) HasIpprivate() bool {
	if o != nil && !IsNil(o.Ipprivate) {
		return true
	}

	return false
}

// SetIpprivate gets a reference to the given string and assigns it to the Ipprivate field.
func (o *VPCMembersInner) SetIpprivate(v string) {
	o.Ipprivate = &v
}

// GetPubkey returns the Pubkey field value if set, zero value otherwise.
func (o *VPCMembersInner) GetPubkey() string {
	if o == nil || IsNil(o.Pubkey) {
		var ret string
		return ret
	}
	return *o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCMembersInner) GetPubkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Pubkey) {
		return nil, false
	}
	return o.Pubkey, true
}

// HasPubkey returns a boolean if a field has been set.
func (o *VPCMembersInner) HasPubkey() bool {
	if o != nil && !IsNil(o.Pubkey) {
		return true
	}

	return false
}

// SetPubkey gets a reference to the given string and assigns it to the Pubkey field.
func (o *VPCMembersInner) SetPubkey(v string) {
	o.Pubkey = &v
}

// GetPrivkey returns the Privkey field value if set, zero value otherwise.
func (o *VPCMembersInner) GetPrivkey() string {
	if o == nil || IsNil(o.Privkey) {
		var ret string
		return ret
	}
	return *o.Privkey
}

// GetPrivkeyOk returns a tuple with the Privkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCMembersInner) GetPrivkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Privkey) {
		return nil, false
	}
	return o.Privkey, true
}

// HasPrivkey returns a boolean if a field has been set.
func (o *VPCMembersInner) HasPrivkey() bool {
	if o != nil && !IsNil(o.Privkey) {
		return true
	}

	return false
}

// SetPrivkey gets a reference to the given string and assigns it to the Privkey field.
func (o *VPCMembersInner) SetPrivkey(v string) {
	o.Privkey = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *VPCMembersInner) GetConfig() string {
	if o == nil || IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCMembersInner) GetConfigOk() (*string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *VPCMembersInner) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *VPCMembersInner) SetConfig(v string) {
	o.Config = &v
}

func (o VPCMembersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VPCMembersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ippublic) {
		toSerialize["ippublic"] = o.Ippublic
	}
	if !IsNil(o.Ipprivate) {
		toSerialize["ipprivate"] = o.Ipprivate
	}
	if !IsNil(o.Pubkey) {
		toSerialize["pubkey"] = o.Pubkey
	}
	if !IsNil(o.Privkey) {
		toSerialize["privkey"] = o.Privkey
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableVPCMembersInner struct {
	value *VPCMembersInner
	isSet bool
}

func (v NullableVPCMembersInner) Get() *VPCMembersInner {
	return v.value
}

func (v *NullableVPCMembersInner) Set(val *VPCMembersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVPCMembersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVPCMembersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVPCMembersInner(val *VPCMembersInner) *NullableVPCMembersInner {
	return &NullableVPCMembersInner{value: val, isSet: true}
}

func (v NullableVPCMembersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVPCMembersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

