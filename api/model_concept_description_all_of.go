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

// ConceptDescriptionAllOf struct for ConceptDescriptionAllOf
type ConceptDescriptionAllOf struct {
	IsCaseOf []Reference `json:"isCaseOf,omitempty"`
}

// NewConceptDescriptionAllOf instantiates a new ConceptDescriptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConceptDescriptionAllOf() *ConceptDescriptionAllOf {
	this := ConceptDescriptionAllOf{}
	return &this
}

// NewConceptDescriptionAllOfWithDefaults instantiates a new ConceptDescriptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConceptDescriptionAllOfWithDefaults() *ConceptDescriptionAllOf {
	this := ConceptDescriptionAllOf{}
	return &this
}

// GetIsCaseOf returns the IsCaseOf field value if set, zero value otherwise.
func (o *ConceptDescriptionAllOf) GetIsCaseOf() []Reference {
	if o == nil || o.IsCaseOf == nil {
		var ret []Reference
		return ret
	}
	return o.IsCaseOf
}

// GetIsCaseOfOk returns a tuple with the IsCaseOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConceptDescriptionAllOf) GetIsCaseOfOk() ([]Reference, bool) {
	if o == nil || o.IsCaseOf == nil {
		return nil, false
	}
	return o.IsCaseOf, true
}

// HasIsCaseOf returns a boolean if a field has been set.
func (o *ConceptDescriptionAllOf) HasIsCaseOf() bool {
	if o != nil && o.IsCaseOf != nil {
		return true
	}

	return false
}

// SetIsCaseOf gets a reference to the given []Reference and assigns it to the IsCaseOf field.
func (o *ConceptDescriptionAllOf) SetIsCaseOf(v []Reference) {
	o.IsCaseOf = v
}

func (o ConceptDescriptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsCaseOf != nil {
		toSerialize["isCaseOf"] = o.IsCaseOf
	}
	return json.Marshal(toSerialize)
}

type NullableConceptDescriptionAllOf struct {
	value *ConceptDescriptionAllOf
	isSet bool
}

func (v NullableConceptDescriptionAllOf) Get() *ConceptDescriptionAllOf {
	return v.value
}

func (v *NullableConceptDescriptionAllOf) Set(val *ConceptDescriptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConceptDescriptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConceptDescriptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConceptDescriptionAllOf(val *ConceptDescriptionAllOf) *NullableConceptDescriptionAllOf {
	return &NullableConceptDescriptionAllOf{value: val, isSet: true}
}

func (v NullableConceptDescriptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConceptDescriptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
