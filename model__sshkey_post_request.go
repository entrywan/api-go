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

// checks if the SshkeyPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshkeyPostRequest{}

// SshkeyPostRequest struct for SshkeyPostRequest
type SshkeyPostRequest struct {
	// Name of the ssh key
	Name *string `json:"name,omitempty"`
	// SSH public key
	Pub *string `json:"pub,omitempty"`
}

// NewSshkeyPostRequest instantiates a new SshkeyPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshkeyPostRequest() *SshkeyPostRequest {
	this := SshkeyPostRequest{}
	return &this
}

// NewSshkeyPostRequestWithDefaults instantiates a new SshkeyPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshkeyPostRequestWithDefaults() *SshkeyPostRequest {
	this := SshkeyPostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SshkeyPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshkeyPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SshkeyPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SshkeyPostRequest) SetName(v string) {
	o.Name = &v
}

// GetPub returns the Pub field value if set, zero value otherwise.
func (o *SshkeyPostRequest) GetPub() string {
	if o == nil || IsNil(o.Pub) {
		var ret string
		return ret
	}
	return *o.Pub
}

// GetPubOk returns a tuple with the Pub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshkeyPostRequest) GetPubOk() (*string, bool) {
	if o == nil || IsNil(o.Pub) {
		return nil, false
	}
	return o.Pub, true
}

// HasPub returns a boolean if a field has been set.
func (o *SshkeyPostRequest) HasPub() bool {
	if o != nil && !IsNil(o.Pub) {
		return true
	}

	return false
}

// SetPub gets a reference to the given string and assigns it to the Pub field.
func (o *SshkeyPostRequest) SetPub(v string) {
	o.Pub = &v
}

func (o SshkeyPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshkeyPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pub) {
		toSerialize["pub"] = o.Pub
	}
	return toSerialize, nil
}

type NullableSshkeyPostRequest struct {
	value *SshkeyPostRequest
	isSet bool
}

func (v NullableSshkeyPostRequest) Get() *SshkeyPostRequest {
	return v.value
}

func (v *NullableSshkeyPostRequest) Set(val *SshkeyPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSshkeyPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSshkeyPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshkeyPostRequest(val *SshkeyPostRequest) *NullableSshkeyPostRequest {
	return &NullableSshkeyPostRequest{value: val, isSet: true}
}

func (v NullableSshkeyPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshkeyPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


