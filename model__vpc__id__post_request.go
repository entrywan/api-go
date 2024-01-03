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

// checks if the VpcIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpcIdPostRequest{}

// VpcIdPostRequest struct for VpcIdPostRequest
type VpcIdPostRequest struct {
	// Name of the VPC
	Name *string `json:"name,omitempty"`
	// CIDR prefix of the VPC
	Prefix *string `json:"prefix,omitempty"`
}

// NewVpcIdPostRequest instantiates a new VpcIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpcIdPostRequest() *VpcIdPostRequest {
	this := VpcIdPostRequest{}
	return &this
}

// NewVpcIdPostRequestWithDefaults instantiates a new VpcIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpcIdPostRequestWithDefaults() *VpcIdPostRequest {
	this := VpcIdPostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VpcIdPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpcIdPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VpcIdPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VpcIdPostRequest) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *VpcIdPostRequest) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpcIdPostRequest) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *VpcIdPostRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *VpcIdPostRequest) SetPrefix(v string) {
	o.Prefix = &v
}

func (o VpcIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpcIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	return toSerialize, nil
}

type NullableVpcIdPostRequest struct {
	value *VpcIdPostRequest
	isSet bool
}

func (v NullableVpcIdPostRequest) Get() *VpcIdPostRequest {
	return v.value
}

func (v *NullableVpcIdPostRequest) Set(val *VpcIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVpcIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVpcIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpcIdPostRequest(val *VpcIdPostRequest) *NullableVpcIdPostRequest {
	return &NullableVpcIdPostRequest{value: val, isSet: true}
}

func (v NullableVpcIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpcIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


