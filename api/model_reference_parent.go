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

// ReferenceParent struct for ReferenceParent
type ReferenceParent struct {
	Keys []Key          `json:"keys"`
	Type ReferenceTypes `json:"type"`
}

// NewReferenceParent instantiates a new ReferenceParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceParent(keys []Key, type_ ReferenceTypes) *ReferenceParent {
	this := ReferenceParent{}
	this.Keys = keys
	this.Type = type_
	return &this
}

// NewReferenceParentWithDefaults instantiates a new ReferenceParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceParentWithDefaults() *ReferenceParent {
	this := ReferenceParent{}
	return &this
}

// GetKeys returns the Keys field value
func (o *ReferenceParent) GetKeys() []Key {
	if o == nil {
		var ret []Key
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *ReferenceParent) GetKeysOk() ([]Key, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *ReferenceParent) SetKeys(v []Key) {
	o.Keys = v
}

// GetType returns the Type field value
func (o *ReferenceParent) GetType() ReferenceTypes {
	if o == nil {
		var ret ReferenceTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReferenceParent) GetTypeOk() (*ReferenceTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReferenceParent) SetType(v ReferenceTypes) {
	o.Type = v
}

func (o ReferenceParent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["keys"] = o.Keys
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableReferenceParent struct {
	value *ReferenceParent
	isSet bool
}

func (v NullableReferenceParent) Get() *ReferenceParent {
	return v.value
}

func (v *NullableReferenceParent) Set(val *ReferenceParent) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceParent) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceParent(val *ReferenceParent) *NullableReferenceParent {
	return &NullableReferenceParent{value: val, isSet: true}
}

func (v NullableReferenceParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
