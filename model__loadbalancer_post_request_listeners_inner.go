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

// checks if the LoadbalancerPostRequestListenersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerPostRequestListenersInner{}

// LoadbalancerPostRequestListenersInner struct for LoadbalancerPostRequestListenersInner
type LoadbalancerPostRequestListenersInner struct {
	Name *string `json:"name,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Targets []LoadbalancerPostRequestListenersInnerTargetsInner `json:"targets,omitempty"`
}

// NewLoadbalancerPostRequestListenersInner instantiates a new LoadbalancerPostRequestListenersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerPostRequestListenersInner() *LoadbalancerPostRequestListenersInner {
	this := LoadbalancerPostRequestListenersInner{}
	return &this
}

// NewLoadbalancerPostRequestListenersInnerWithDefaults instantiates a new LoadbalancerPostRequestListenersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerPostRequestListenersInnerWithDefaults() *LoadbalancerPostRequestListenersInner {
	this := LoadbalancerPostRequestListenersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadbalancerPostRequestListenersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequestListenersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadbalancerPostRequestListenersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadbalancerPostRequestListenersInner) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoadbalancerPostRequestListenersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequestListenersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoadbalancerPostRequestListenersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoadbalancerPostRequestListenersInner) SetPort(v int32) {
	o.Port = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *LoadbalancerPostRequestListenersInner) GetTargets() []LoadbalancerPostRequestListenersInnerTargetsInner {
	if o == nil || IsNil(o.Targets) {
		var ret []LoadbalancerPostRequestListenersInnerTargetsInner
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequestListenersInner) GetTargetsOk() ([]LoadbalancerPostRequestListenersInnerTargetsInner, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *LoadbalancerPostRequestListenersInner) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []LoadbalancerPostRequestListenersInnerTargetsInner and assigns it to the Targets field.
func (o *LoadbalancerPostRequestListenersInner) SetTargets(v []LoadbalancerPostRequestListenersInnerTargetsInner) {
	o.Targets = v
}

func (o LoadbalancerPostRequestListenersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadbalancerPostRequestListenersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableLoadbalancerPostRequestListenersInner struct {
	value *LoadbalancerPostRequestListenersInner
	isSet bool
}

func (v NullableLoadbalancerPostRequestListenersInner) Get() *LoadbalancerPostRequestListenersInner {
	return v.value
}

func (v *NullableLoadbalancerPostRequestListenersInner) Set(val *LoadbalancerPostRequestListenersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerPostRequestListenersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerPostRequestListenersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerPostRequestListenersInner(val *LoadbalancerPostRequestListenersInner) *NullableLoadbalancerPostRequestListenersInner {
	return &NullableLoadbalancerPostRequestListenersInner{value: val, isSet: true}
}

func (v NullableLoadbalancerPostRequestListenersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerPostRequestListenersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


