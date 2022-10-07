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

// HasExtensions struct for HasExtensions
type HasExtensions struct {
	Extensions []Extension `json:"extensions,omitempty"`
}

// NewHasExtensions instantiates a new HasExtensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasExtensions() *HasExtensions {
	this := HasExtensions{}
	return &this
}

// NewHasExtensionsWithDefaults instantiates a new HasExtensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasExtensionsWithDefaults() *HasExtensions {
	this := HasExtensions{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *HasExtensions) GetExtensions() []Extension {
	if o == nil || o.Extensions == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasExtensions) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *HasExtensions) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *HasExtensions) SetExtensions(v []Extension) {
	o.Extensions = v
}

func (o HasExtensions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	return json.Marshal(toSerialize)
}

type NullableHasExtensions struct {
	value *HasExtensions
	isSet bool
}

func (v NullableHasExtensions) Get() *HasExtensions {
	return v.value
}

func (v *NullableHasExtensions) Set(val *HasExtensions) {
	v.value = val
	v.isSet = true
}

func (v NullableHasExtensions) IsSet() bool {
	return v.isSet
}

func (v *NullableHasExtensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasExtensions(val *HasExtensions) *NullableHasExtensions {
	return &NullableHasExtensions{value: val, isSet: true}
}

func (v NullableHasExtensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasExtensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
