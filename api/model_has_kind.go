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

// HasKind struct for HasKind
type HasKind struct {
	Kind *ModelingKind `json:"kind,omitempty"`
}

// NewHasKind instantiates a new HasKind object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasKind() *HasKind {
	this := HasKind{}
	return &this
}

// NewHasKindWithDefaults instantiates a new HasKind object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasKindWithDefaults() *HasKind {
	this := HasKind{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *HasKind) GetKind() ModelingKind {
	if o == nil || o.Kind == nil {
		var ret ModelingKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasKind) GetKindOk() (*ModelingKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *HasKind) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ModelingKind and assigns it to the Kind field.
func (o *HasKind) SetKind(v ModelingKind) {
	o.Kind = &v
}

func (o HasKind) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableHasKind struct {
	value *HasKind
	isSet bool
}

func (v NullableHasKind) Get() *HasKind {
	return v.value
}

func (v *NullableHasKind) Set(val *HasKind) {
	v.value = val
	v.isSet = true
}

func (v NullableHasKind) IsSet() bool {
	return v.isSet
}

func (v *NullableHasKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasKind(val *HasKind) *NullableHasKind {
	return &NullableHasKind{value: val, isSet: true}
}

func (v NullableHasKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
