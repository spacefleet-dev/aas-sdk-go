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

// SubmodelAllOf struct for SubmodelAllOf
type SubmodelAllOf struct {
	SubmodelElements []ISubmodelElement `json:"submodelElements,omitempty"`
}

// NewSubmodelAllOf instantiates a new SubmodelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmodelAllOf() *SubmodelAllOf {
	this := SubmodelAllOf{}
	return &this
}

// NewSubmodelAllOfWithDefaults instantiates a new SubmodelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmodelAllOfWithDefaults() *SubmodelAllOf {
	this := SubmodelAllOf{}
	return &this
}

// GetSubmodelElements returns the SubmodelElements field value if set, zero value otherwise.
func (o *SubmodelAllOf) GetSubmodelElements() []ISubmodelElement {
	if o == nil || o.SubmodelElements == nil {
		var ret []ISubmodelElement
		return ret
	}
	return o.SubmodelElements
}

// GetSubmodelElementsOk returns a tuple with the SubmodelElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelAllOf) GetSubmodelElementsOk() ([]ISubmodelElement, bool) {
	if o == nil || o.SubmodelElements == nil {
		return nil, false
	}
	return o.SubmodelElements, true
}

// HasSubmodelElements returns a boolean if a field has been set.
func (o *SubmodelAllOf) HasSubmodelElements() bool {
	if o != nil && o.SubmodelElements != nil {
		return true
	}

	return false
}

// SetSubmodelElements gets a reference to the given []ISubmodelElement and assigns it to the SubmodelElements field.
func (o *SubmodelAllOf) SetSubmodelElements(v []ISubmodelElement) {
	o.SubmodelElements = v
}

func (o SubmodelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubmodelElements != nil {
		toSerialize["submodelElements"] = o.SubmodelElements
	}
	return json.Marshal(toSerialize)
}

type NullableSubmodelAllOf struct {
	value *SubmodelAllOf
	isSet bool
}

func (v NullableSubmodelAllOf) Get() *SubmodelAllOf {
	return v.value
}

func (v *NullableSubmodelAllOf) Set(val *SubmodelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmodelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmodelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmodelAllOf(val *SubmodelAllOf) *NullableSubmodelAllOf {
	return &NullableSubmodelAllOf{value: val, isSet: true}
}

func (v NullableSubmodelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmodelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
