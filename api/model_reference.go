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

// Reference struct for Reference
type Reference struct {
	Keys               []Key            `json:"keys"`
	Type               ReferenceTypes   `json:"type"`
	ReferredSemanticId *ReferenceParent `json:"referredSemanticId,omitempty"`
}

// NewReference instantiates a new Reference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReference(keys []Key, type_ ReferenceTypes) *Reference {
	this := Reference{}
	this.Keys = keys
	this.Type = type_
	return &this
}

// NewReferenceWithDefaults instantiates a new Reference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceWithDefaults() *Reference {
	this := Reference{}
	return &this
}

// GetKeys returns the Keys field value
func (o *Reference) GetKeys() []Key {
	if o == nil {
		var ret []Key
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *Reference) GetKeysOk() ([]Key, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *Reference) SetKeys(v []Key) {
	o.Keys = v
}

// GetType returns the Type field value
func (o *Reference) GetType() ReferenceTypes {
	if o == nil {
		var ret ReferenceTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Reference) GetTypeOk() (*ReferenceTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Reference) SetType(v ReferenceTypes) {
	o.Type = v
}

// GetReferredSemanticId returns the ReferredSemanticId field value if set, zero value otherwise.
func (o *Reference) GetReferredSemanticId() ReferenceParent {
	if o == nil || o.ReferredSemanticId == nil {
		var ret ReferenceParent
		return ret
	}
	return *o.ReferredSemanticId
}

// GetReferredSemanticIdOk returns a tuple with the ReferredSemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reference) GetReferredSemanticIdOk() (*ReferenceParent, bool) {
	if o == nil || o.ReferredSemanticId == nil {
		return nil, false
	}
	return o.ReferredSemanticId, true
}

// HasReferredSemanticId returns a boolean if a field has been set.
func (o *Reference) HasReferredSemanticId() bool {
	if o != nil && o.ReferredSemanticId != nil {
		return true
	}

	return false
}

// SetReferredSemanticId gets a reference to the given ReferenceParent and assigns it to the ReferredSemanticId field.
func (o *Reference) SetReferredSemanticId(v ReferenceParent) {
	o.ReferredSemanticId = &v
}

func (o Reference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["keys"] = o.Keys
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ReferredSemanticId != nil {
		toSerialize["referredSemanticId"] = o.ReferredSemanticId
	}
	return json.Marshal(toSerialize)
}

type NullableReference struct {
	value *Reference
	isSet bool
}

func (v NullableReference) Get() *Reference {
	return v.value
}

func (v *NullableReference) Set(val *Reference) {
	v.value = val
	v.isSet = true
}

func (v NullableReference) IsSet() bool {
	return v.isSet
}

func (v *NullableReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReference(val *Reference) *NullableReference {
	return &NullableReference{value: val, isSet: true}
}

func (v NullableReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
