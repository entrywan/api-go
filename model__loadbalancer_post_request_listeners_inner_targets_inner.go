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

// checks if the LoadbalancerPostRequestListenersInnerTargetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerPostRequestListenersInnerTargetsInner{}

// LoadbalancerPostRequestListenersInnerTargetsInner struct for LoadbalancerPostRequestListenersInnerTargetsInner
type LoadbalancerPostRequestListenersInnerTargetsInner struct {
	Ip *string `json:"ip,omitempty"`
	Port *int32 `json:"port,omitempty"`
}

// NewLoadbalancerPostRequestListenersInnerTargetsInner instantiates a new LoadbalancerPostRequestListenersInnerTargetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerPostRequestListenersInnerTargetsInner() *LoadbalancerPostRequestListenersInnerTargetsInner {
	this := LoadbalancerPostRequestListenersInnerTargetsInner{}
	return &this
}

// NewLoadbalancerPostRequestListenersInnerTargetsInnerWithDefaults instantiates a new LoadbalancerPostRequestListenersInnerTargetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerPostRequestListenersInnerTargetsInnerWithDefaults() *LoadbalancerPostRequestListenersInnerTargetsInner {
	this := LoadbalancerPostRequestListenersInnerTargetsInner{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) SetIp(v string) {
	o.Ip = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoadbalancerPostRequestListenersInnerTargetsInner) SetPort(v int32) {
	o.Port = &v
}

func (o LoadbalancerPostRequestListenersInnerTargetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadbalancerPostRequestListenersInnerTargetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableLoadbalancerPostRequestListenersInnerTargetsInner struct {
	value *LoadbalancerPostRequestListenersInnerTargetsInner
	isSet bool
}

func (v NullableLoadbalancerPostRequestListenersInnerTargetsInner) Get() *LoadbalancerPostRequestListenersInnerTargetsInner {
	return v.value
}

func (v *NullableLoadbalancerPostRequestListenersInnerTargetsInner) Set(val *LoadbalancerPostRequestListenersInnerTargetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerPostRequestListenersInnerTargetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerPostRequestListenersInnerTargetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerPostRequestListenersInnerTargetsInner(val *LoadbalancerPostRequestListenersInnerTargetsInner) *NullableLoadbalancerPostRequestListenersInnerTargetsInner {
	return &NullableLoadbalancerPostRequestListenersInnerTargetsInner{value: val, isSet: true}
}

func (v NullableLoadbalancerPostRequestListenersInnerTargetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerPostRequestListenersInnerTargetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


