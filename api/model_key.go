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

// Key struct for Key
type Key struct {
	Type  KeyTypes `json:"type"`
	Value string   `json:"value"`
}

// NewKey instantiates a new Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKey(type_ KeyTypes, value string) *Key {
	this := Key{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewKeyWithDefaults instantiates a new Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyWithDefaults() *Key {
	this := Key{}
	return &this
}

// GetType returns the Type field value
func (o *Key) GetType() KeyTypes {
	if o == nil {
		var ret KeyTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Key) GetTypeOk() (*KeyTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Key) SetType(v KeyTypes) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *Key) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Key) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Key) SetValue(v string) {
	o.Value = v
}

func (o Key) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableKey struct {
	value *Key
	isSet bool
}

func (v NullableKey) Get() *Key {
	return v.value
}

func (v *NullableKey) Set(val *Key) {
	v.value = val
	v.isSet = true
}

func (v NullableKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKey(val *Key) *NullableKey {
	return &NullableKey{value: val, isSet: true}
}

func (v NullableKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
