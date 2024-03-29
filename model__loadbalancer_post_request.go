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

// checks if the LoadbalancerPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerPostRequest{}

// LoadbalancerPostRequest struct for LoadbalancerPostRequest
type LoadbalancerPostRequest struct {
	// Name of the load balancer
	Name *string `json:"name,omitempty"`
	// Location of the load balancer
	Location *string `json:"location,omitempty"`
	Listeners []LoadbalancerPostRequestListenersInner `json:"listeners,omitempty"`
}

// NewLoadbalancerPostRequest instantiates a new LoadbalancerPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerPostRequest() *LoadbalancerPostRequest {
	this := LoadbalancerPostRequest{}
	return &this
}

// NewLoadbalancerPostRequestWithDefaults instantiates a new LoadbalancerPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerPostRequestWithDefaults() *LoadbalancerPostRequest {
	this := LoadbalancerPostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadbalancerPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadbalancerPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadbalancerPostRequest) SetName(v string) {
	o.Name = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LoadbalancerPostRequest) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequest) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LoadbalancerPostRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *LoadbalancerPostRequest) SetLocation(v string) {
	o.Location = &v
}

// GetListeners returns the Listeners field value if set, zero value otherwise.
func (o *LoadbalancerPostRequest) GetListeners() []LoadbalancerPostRequestListenersInner {
	if o == nil || IsNil(o.Listeners) {
		var ret []LoadbalancerPostRequestListenersInner
		return ret
	}
	return o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerPostRequest) GetListenersOk() ([]LoadbalancerPostRequestListenersInner, bool) {
	if o == nil || IsNil(o.Listeners) {
		return nil, false
	}
	return o.Listeners, true
}

// HasListeners returns a boolean if a field has been set.
func (o *LoadbalancerPostRequest) HasListeners() bool {
	if o != nil && !IsNil(o.Listeners) {
		return true
	}

	return false
}

// SetListeners gets a reference to the given []LoadbalancerPostRequestListenersInner and assigns it to the Listeners field.
func (o *LoadbalancerPostRequest) SetListeners(v []LoadbalancerPostRequestListenersInner) {
	o.Listeners = v
}

func (o LoadbalancerPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadbalancerPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Listeners) {
		toSerialize["listeners"] = o.Listeners
	}
	return toSerialize, nil
}

type NullableLoadbalancerPostRequest struct {
	value *LoadbalancerPostRequest
	isSet bool
}

func (v NullableLoadbalancerPostRequest) Get() *LoadbalancerPostRequest {
	return v.value
}

func (v *NullableLoadbalancerPostRequest) Set(val *LoadbalancerPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerPostRequest(val *LoadbalancerPostRequest) *NullableLoadbalancerPostRequest {
	return &NullableLoadbalancerPostRequest{value: val, isSet: true}
}

func (v NullableLoadbalancerPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


