//lint:file-ignore * Autogenerated code
/*
BaSyx Asset Administration Shell HTTP REST-API

The full description of the generic BaSyx Asset Administration Shell HTTP REST-API

API version: v1
Contact: constantin.ziesche@bosch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SubmodelReference struct for SubmodelReference
type SubmodelReference struct {
	Keys []Key `json:"keys,omitempty"`
}

// NewSubmodelReference instantiates a new SubmodelReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmodelReference() *SubmodelReference {
	this := SubmodelReference{}
	return &this
}

// NewSubmodelReferenceWithDefaults instantiates a new SubmodelReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmodelReferenceWithDefaults() *SubmodelReference {
	this := SubmodelReference{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubmodelReference) GetKeys() []Key {
	if o == nil {
		var ret []Key
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmodelReference) GetKeysOk() ([]Key, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SubmodelReference) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []Key and assigns it to the Keys field.
func (o *SubmodelReference) SetKeys(v []Key) {
	o.Keys = v
}

func (o SubmodelReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableSubmodelReference struct {
	value *SubmodelReference
	isSet bool
}

func (v NullableSubmodelReference) Get() *SubmodelReference {
	return v.value
}

func (v *NullableSubmodelReference) Set(val *SubmodelReference) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmodelReference) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmodelReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmodelReference(val *SubmodelReference) *NullableSubmodelReference {
	return &NullableSubmodelReference{value: val, isSet: true}
}

func (v NullableSubmodelReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmodelReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
