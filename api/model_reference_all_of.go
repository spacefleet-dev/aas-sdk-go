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

// ReferenceAllOf struct for ReferenceAllOf
type ReferenceAllOf struct {
	ReferredSemanticId *ReferenceParent `json:"referredSemanticId,omitempty"`
}

// NewReferenceAllOf instantiates a new ReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceAllOf() *ReferenceAllOf {
	this := ReferenceAllOf{}
	return &this
}

// NewReferenceAllOfWithDefaults instantiates a new ReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceAllOfWithDefaults() *ReferenceAllOf {
	this := ReferenceAllOf{}
	return &this
}

// GetReferredSemanticId returns the ReferredSemanticId field value if set, zero value otherwise.
func (o *ReferenceAllOf) GetReferredSemanticId() ReferenceParent {
	if o == nil || o.ReferredSemanticId == nil {
		var ret ReferenceParent
		return ret
	}
	return *o.ReferredSemanticId
}

// GetReferredSemanticIdOk returns a tuple with the ReferredSemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceAllOf) GetReferredSemanticIdOk() (*ReferenceParent, bool) {
	if o == nil || o.ReferredSemanticId == nil {
		return nil, false
	}
	return o.ReferredSemanticId, true
}

// HasReferredSemanticId returns a boolean if a field has been set.
func (o *ReferenceAllOf) HasReferredSemanticId() bool {
	if o != nil && o.ReferredSemanticId != nil {
		return true
	}

	return false
}

// SetReferredSemanticId gets a reference to the given ReferenceParent and assigns it to the ReferredSemanticId field.
func (o *ReferenceAllOf) SetReferredSemanticId(v ReferenceParent) {
	o.ReferredSemanticId = &v
}

func (o ReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferredSemanticId != nil {
		toSerialize["referredSemanticId"] = o.ReferredSemanticId
	}
	return json.Marshal(toSerialize)
}

type NullableReferenceAllOf struct {
	value *ReferenceAllOf
	isSet bool
}

func (v NullableReferenceAllOf) Get() *ReferenceAllOf {
	return v.value
}

func (v *NullableReferenceAllOf) Set(val *ReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceAllOf(val *ReferenceAllOf) *NullableReferenceAllOf {
	return &NullableReferenceAllOf{value: val, isSet: true}
}

func (v NullableReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
