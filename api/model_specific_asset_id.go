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

// SpecificAssetId struct for SpecificAssetId
type SpecificAssetId struct {
	SemanticId              *Reference  `json:"semanticId,omitempty"`
	SupplementalSemanticIds []Reference `json:"supplementalSemanticIds,omitempty"`
	ExternalSubjectId       Reference   `json:"externalSubjectId"`
	Name                    string      `json:"name"`
	Value                   string      `json:"value"`
}

// NewSpecificAssetId instantiates a new SpecificAssetId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecificAssetId(externalSubjectId Reference, name string, value string) *SpecificAssetId {
	this := SpecificAssetId{}
	this.ExternalSubjectId = externalSubjectId
	this.Name = name
	this.Value = value
	return &this
}

// NewSpecificAssetIdWithDefaults instantiates a new SpecificAssetId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecificAssetIdWithDefaults() *SpecificAssetId {
	this := SpecificAssetId{}
	return &this
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *SpecificAssetId) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificAssetId) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *SpecificAssetId) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *SpecificAssetId) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *SpecificAssetId) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificAssetId) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *SpecificAssetId) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *SpecificAssetId) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetExternalSubjectId returns the ExternalSubjectId field value
func (o *SpecificAssetId) GetExternalSubjectId() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.ExternalSubjectId
}

// GetExternalSubjectIdOk returns a tuple with the ExternalSubjectId field value
// and a boolean to check if the value has been set.
func (o *SpecificAssetId) GetExternalSubjectIdOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalSubjectId, true
}

// SetExternalSubjectId sets field value
func (o *SpecificAssetId) SetExternalSubjectId(v Reference) {
	o.ExternalSubjectId = v
}

// GetName returns the Name field value
func (o *SpecificAssetId) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpecificAssetId) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpecificAssetId) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *SpecificAssetId) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SpecificAssetId) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SpecificAssetId) SetValue(v string) {
	o.Value = v
}

func (o SpecificAssetId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SemanticId != nil {
		toSerialize["semanticId"] = o.SemanticId
	}
	if o.SupplementalSemanticIds != nil {
		toSerialize["supplementalSemanticIds"] = o.SupplementalSemanticIds
	}
	if true {
		toSerialize["externalSubjectId"] = o.ExternalSubjectId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSpecificAssetId struct {
	value *SpecificAssetId
	isSet bool
}

func (v NullableSpecificAssetId) Get() *SpecificAssetId {
	return v.value
}

func (v *NullableSpecificAssetId) Set(val *SpecificAssetId) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecificAssetId) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecificAssetId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecificAssetId(val *SpecificAssetId) *NullableSpecificAssetId {
	return &NullableSpecificAssetId{value: val, isSet: true}
}

func (v NullableSpecificAssetId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecificAssetId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}