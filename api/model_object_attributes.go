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

// ObjectAttributes struct for ObjectAttributes
type ObjectAttributes struct {
	ObjectAttribute []Property `json:"objectAttribute,omitempty"`
}

// NewObjectAttributes instantiates a new ObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectAttributes() *ObjectAttributes {
	this := ObjectAttributes{}
	return &this
}

// NewObjectAttributesWithDefaults instantiates a new ObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectAttributesWithDefaults() *ObjectAttributes {
	this := ObjectAttributes{}
	return &this
}

// GetObjectAttribute returns the ObjectAttribute field value if set, zero value otherwise.
func (o *ObjectAttributes) GetObjectAttribute() []Property {
	if o == nil || o.ObjectAttribute == nil {
		var ret []Property
		return ret
	}
	return o.ObjectAttribute
}

// GetObjectAttributeOk returns a tuple with the ObjectAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectAttributes) GetObjectAttributeOk() ([]Property, bool) {
	if o == nil || o.ObjectAttribute == nil {
		return nil, false
	}
	return o.ObjectAttribute, true
}

// HasObjectAttribute returns a boolean if a field has been set.
func (o *ObjectAttributes) HasObjectAttribute() bool {
	if o != nil && o.ObjectAttribute != nil {
		return true
	}

	return false
}

// SetObjectAttribute gets a reference to the given []Property and assigns it to the ObjectAttribute field.
func (o *ObjectAttributes) SetObjectAttribute(v []Property) {
	o.ObjectAttribute = v
}

func (o ObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectAttribute != nil {
		toSerialize["objectAttribute"] = o.ObjectAttribute
	}
	return json.Marshal(toSerialize)
}

type NullableObjectAttributes struct {
	value *ObjectAttributes
	isSet bool
}

func (v NullableObjectAttributes) Get() *ObjectAttributes {
	return v.value
}

func (v *NullableObjectAttributes) Set(val *ObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectAttributes(val *ObjectAttributes) *NullableObjectAttributes {
	return &NullableObjectAttributes{value: val, isSet: true}
}

func (v NullableObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
