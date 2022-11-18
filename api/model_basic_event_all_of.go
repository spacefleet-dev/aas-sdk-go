//lint:file-ignore * Autogenerated code
/*
BaSyx Asset Administration Shell Repository HTTP REST-API

The full description of the generic BaSyx Asset Administration Shell Repository HTTP REST-API

API version: v1
Contact: constantin.ziesche@bosch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BasicEventAllOf struct for BasicEventAllOf
type BasicEventAllOf struct {
	Observed Reference `json:"observed"`
}

// NewBasicEventAllOf instantiates a new BasicEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicEventAllOf(observed Reference) *BasicEventAllOf {
	this := BasicEventAllOf{}
	this.Observed = observed
	return &this
}

// NewBasicEventAllOfWithDefaults instantiates a new BasicEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicEventAllOfWithDefaults() *BasicEventAllOf {
	this := BasicEventAllOf{}
	return &this
}

// GetObserved returns the Observed field value
func (o *BasicEventAllOf) GetObserved() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Observed
}

// GetObservedOk returns a tuple with the Observed field value
// and a boolean to check if the value has been set.
func (o *BasicEventAllOf) GetObservedOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observed, true
}

// SetObserved sets field value
func (o *BasicEventAllOf) SetObserved(v Reference) {
	o.Observed = v
}

func (o BasicEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["observed"] = o.Observed
	}
	return json.Marshal(toSerialize)
}

type NullableBasicEventAllOf struct {
	value *BasicEventAllOf
	isSet bool
}

func (v NullableBasicEventAllOf) Get() *BasicEventAllOf {
	return v.value
}

func (v *NullableBasicEventAllOf) Set(val *BasicEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicEventAllOf(val *BasicEventAllOf) *NullableBasicEventAllOf {
	return &NullableBasicEventAllOf{value: val, isSet: true}
}

func (v NullableBasicEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
