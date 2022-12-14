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

// Qualifier struct for Qualifier
type Qualifier struct {
	SemanticId              *Reference     `json:"semanticId,omitempty"`
	SupplementalSemanticIds []Reference    `json:"supplementalSemanticIds,omitempty"`
	Kind                    *QualifierKind `json:"kind,omitempty"`
	Type                    string         `json:"type"`
	Value                   *string        `json:"value,omitempty"`
	ValueId                 *Reference     `json:"valueId,omitempty"`
	ValueType               DataTypeDefXsd `json:"valueType"`
}

// NewQualifier instantiates a new Qualifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualifier(type_ string, valueType DataTypeDefXsd) *Qualifier {
	this := Qualifier{}
	this.Type = type_
	this.ValueType = valueType
	return &this
}

// NewQualifierWithDefaults instantiates a new Qualifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualifierWithDefaults() *Qualifier {
	this := Qualifier{}
	return &this
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *Qualifier) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Qualifier) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *Qualifier) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *Qualifier) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *Qualifier) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Qualifier) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *Qualifier) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *Qualifier) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Qualifier) GetKind() QualifierKind {
	if o == nil || o.Kind == nil {
		var ret QualifierKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Qualifier) GetKindOk() (*QualifierKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Qualifier) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given QualifierKind and assigns it to the Kind field.
func (o *Qualifier) SetKind(v QualifierKind) {
	o.Kind = &v
}

// GetType returns the Type field value
func (o *Qualifier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Qualifier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Qualifier) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Qualifier) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Qualifier) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Qualifier) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Qualifier) SetValue(v string) {
	o.Value = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *Qualifier) GetValueId() Reference {
	if o == nil || o.ValueId == nil {
		var ret Reference
		return ret
	}
	return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Qualifier) GetValueIdOk() (*Reference, bool) {
	if o == nil || o.ValueId == nil {
		return nil, false
	}
	return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *Qualifier) HasValueId() bool {
	if o != nil && o.ValueId != nil {
		return true
	}

	return false
}

// SetValueId gets a reference to the given Reference and assigns it to the ValueId field.
func (o *Qualifier) SetValueId(v Reference) {
	o.ValueId = &v
}

// GetValueType returns the ValueType field value
func (o *Qualifier) GetValueType() DataTypeDefXsd {
	if o == nil {
		var ret DataTypeDefXsd
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *Qualifier) GetValueTypeOk() (*DataTypeDefXsd, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *Qualifier) SetValueType(v DataTypeDefXsd) {
	o.ValueType = v
}

func (o Qualifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SemanticId != nil {
		toSerialize["semanticId"] = o.SemanticId
	}
	if o.SupplementalSemanticIds != nil {
		toSerialize["supplementalSemanticIds"] = o.SupplementalSemanticIds
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueId != nil {
		toSerialize["valueId"] = o.ValueId
	}
	if true {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableQualifier struct {
	value *Qualifier
	isSet bool
}

func (v NullableQualifier) Get() *Qualifier {
	return v.value
}

func (v *NullableQualifier) Set(val *Qualifier) {
	v.value = val
	v.isSet = true
}

func (v NullableQualifier) IsSet() bool {
	return v.isSet
}

func (v *NullableQualifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualifier(val *Qualifier) *NullableQualifier {
	return &NullableQualifier{value: val, isSet: true}
}

func (v NullableQualifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
