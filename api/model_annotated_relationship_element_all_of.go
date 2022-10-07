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

// AnnotatedRelationshipElementAllOf struct for AnnotatedRelationshipElementAllOf
type AnnotatedRelationshipElementAllOf struct {
	Annotations []ISubmodelElement `json:"annotations,omitempty"`
}

// NewAnnotatedRelationshipElementAllOf instantiates a new AnnotatedRelationshipElementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotatedRelationshipElementAllOf() *AnnotatedRelationshipElementAllOf {
	this := AnnotatedRelationshipElementAllOf{}
	return &this
}

// NewAnnotatedRelationshipElementAllOfWithDefaults instantiates a new AnnotatedRelationshipElementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotatedRelationshipElementAllOfWithDefaults() *AnnotatedRelationshipElementAllOf {
	this := AnnotatedRelationshipElementAllOf{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AnnotatedRelationshipElementAllOf) GetAnnotations() []ISubmodelElement {
	if o == nil || o.Annotations == nil {
		var ret []ISubmodelElement
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotatedRelationshipElementAllOf) GetAnnotationsOk() ([]ISubmodelElement, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AnnotatedRelationshipElementAllOf) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []ISubmodelElement and assigns it to the Annotations field.
func (o *AnnotatedRelationshipElementAllOf) SetAnnotations(v []ISubmodelElement) {
	o.Annotations = v
}

func (o AnnotatedRelationshipElementAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableAnnotatedRelationshipElementAllOf struct {
	value *AnnotatedRelationshipElementAllOf
	isSet bool
}

func (v NullableAnnotatedRelationshipElementAllOf) Get() *AnnotatedRelationshipElementAllOf {
	return v.value
}

func (v *NullableAnnotatedRelationshipElementAllOf) Set(val *AnnotatedRelationshipElementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotatedRelationshipElementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotatedRelationshipElementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotatedRelationshipElementAllOf(val *AnnotatedRelationshipElementAllOf) *NullableAnnotatedRelationshipElementAllOf {
	return &NullableAnnotatedRelationshipElementAllOf{value: val, isSet: true}
}

func (v NullableAnnotatedRelationshipElementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotatedRelationshipElementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}