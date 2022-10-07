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

// Range struct for Range
type Range struct {
	Extensions                 []Extension                 `json:"extensions,omitempty"`
	Category                   *string                     `json:"category,omitempty"`
	Checksum                   *string                     `json:"checksum,omitempty"`
	Description                []LangString                `json:"description,omitempty"`
	DisplayName                []LangString                `json:"displayName,omitempty"`
	IdShort                    *string                     `json:"idShort,omitempty"`
	ModelType                  ModelType                   `json:"modelType"`
	Kind                       *ModelingKind               `json:"kind,omitempty"`
	SemanticId                 *Reference                  `json:"semanticId,omitempty"`
	SupplementalSemanticIds    []Reference                 `json:"supplementalSemanticIds,omitempty"`
	Qualifiers                 []Qualifier                 `json:"qualifiers,omitempty"`
	EmbeddedDataSpecifications []EmbeddedDataSpecification `json:"embeddedDataSpecifications,omitempty"`
	Max                        *string                     `json:"max,omitempty"`
	Min                        *string                     `json:"min,omitempty"`
	ValueType                  DataTypeDefXsd              `json:"valueType"`
}

// NewRange instantiates a new Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRange(modelType ModelType, valueType DataTypeDefXsd) *Range {
	this := Range{}
	this.ModelType = modelType
	this.ValueType = valueType
	return &this
}

// NewRangeWithDefaults instantiates a new Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeWithDefaults() *Range {
	this := Range{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Range) GetExtensions() []Extension {
	if o == nil || o.Extensions == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Range) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *Range) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Range) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Range) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Range) SetCategory(v string) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *Range) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *Range) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *Range) SetChecksum(v string) {
	o.Checksum = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Range) GetDescription() []LangString {
	if o == nil || o.Description == nil {
		var ret []LangString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetDescriptionOk() ([]LangString, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Range) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LangString and assigns it to the Description field.
func (o *Range) SetDescription(v []LangString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Range) GetDisplayName() []LangString {
	if o == nil || o.DisplayName == nil {
		var ret []LangString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetDisplayNameOk() ([]LangString, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Range) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LangString and assigns it to the DisplayName field.
func (o *Range) SetDisplayName(v []LangString) {
	o.DisplayName = v
}

// GetIdShort returns the IdShort field value if set, zero value otherwise.
func (o *Range) GetIdShort() string {
	if o == nil || o.IdShort == nil {
		var ret string
		return ret
	}
	return *o.IdShort
}

// GetIdShortOk returns a tuple with the IdShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetIdShortOk() (*string, bool) {
	if o == nil || o.IdShort == nil {
		return nil, false
	}
	return o.IdShort, true
}

// HasIdShort returns a boolean if a field has been set.
func (o *Range) HasIdShort() bool {
	if o != nil && o.IdShort != nil {
		return true
	}

	return false
}

// SetIdShort gets a reference to the given string and assigns it to the IdShort field.
func (o *Range) SetIdShort(v string) {
	o.IdShort = &v
}

// GetModelType returns the ModelType field value
func (o *Range) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *Range) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *Range) SetModelType(v ModelType) {
	o.ModelType = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Range) GetKind() ModelingKind {
	if o == nil || o.Kind == nil {
		var ret ModelingKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetKindOk() (*ModelingKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Range) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ModelingKind and assigns it to the Kind field.
func (o *Range) SetKind(v ModelingKind) {
	o.Kind = &v
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *Range) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *Range) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *Range) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *Range) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *Range) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *Range) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetQualifiers returns the Qualifiers field value if set, zero value otherwise.
func (o *Range) GetQualifiers() []Qualifier {
	if o == nil || o.Qualifiers == nil {
		var ret []Qualifier
		return ret
	}
	return o.Qualifiers
}

// GetQualifiersOk returns a tuple with the Qualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetQualifiersOk() ([]Qualifier, bool) {
	if o == nil || o.Qualifiers == nil {
		return nil, false
	}
	return o.Qualifiers, true
}

// HasQualifiers returns a boolean if a field has been set.
func (o *Range) HasQualifiers() bool {
	if o != nil && o.Qualifiers != nil {
		return true
	}

	return false
}

// SetQualifiers gets a reference to the given []Qualifier and assigns it to the Qualifiers field.
func (o *Range) SetQualifiers(v []Qualifier) {
	o.Qualifiers = v
}

// GetEmbeddedDataSpecifications returns the EmbeddedDataSpecifications field value if set, zero value otherwise.
func (o *Range) GetEmbeddedDataSpecifications() []EmbeddedDataSpecification {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		var ret []EmbeddedDataSpecification
		return ret
	}
	return o.EmbeddedDataSpecifications
}

// GetEmbeddedDataSpecificationsOk returns a tuple with the EmbeddedDataSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetEmbeddedDataSpecificationsOk() ([]EmbeddedDataSpecification, bool) {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		return nil, false
	}
	return o.EmbeddedDataSpecifications, true
}

// HasEmbeddedDataSpecifications returns a boolean if a field has been set.
func (o *Range) HasEmbeddedDataSpecifications() bool {
	if o != nil && o.EmbeddedDataSpecifications != nil {
		return true
	}

	return false
}

// SetEmbeddedDataSpecifications gets a reference to the given []EmbeddedDataSpecification and assigns it to the EmbeddedDataSpecifications field.
func (o *Range) SetEmbeddedDataSpecifications(v []EmbeddedDataSpecification) {
	o.EmbeddedDataSpecifications = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Range) GetMax() string {
	if o == nil || o.Max == nil {
		var ret string
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetMaxOk() (*string, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Range) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given string and assigns it to the Max field.
func (o *Range) SetMax(v string) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Range) GetMin() string {
	if o == nil || o.Min == nil {
		var ret string
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetMinOk() (*string, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Range) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given string and assigns it to the Min field.
func (o *Range) SetMin(v string) {
	o.Min = &v
}

// GetValueType returns the ValueType field value
func (o *Range) GetValueType() DataTypeDefXsd {
	if o == nil {
		var ret DataTypeDefXsd
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *Range) GetValueTypeOk() (*DataTypeDefXsd, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *Range) SetValueType(v DataTypeDefXsd) {
	o.ValueType = v
}

func (o Range) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.IdShort != nil {
		toSerialize["idShort"] = o.IdShort
	}
	if true {
		toSerialize["modelType"] = o.ModelType
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.SemanticId != nil {
		toSerialize["semanticId"] = o.SemanticId
	}
	if o.SupplementalSemanticIds != nil {
		toSerialize["supplementalSemanticIds"] = o.SupplementalSemanticIds
	}
	if o.Qualifiers != nil {
		toSerialize["qualifiers"] = o.Qualifiers
	}
	if o.EmbeddedDataSpecifications != nil {
		toSerialize["embeddedDataSpecifications"] = o.EmbeddedDataSpecifications
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if true {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableRange struct {
	value *Range
	isSet bool
}

func (v NullableRange) Get() *Range {
	return v.value
}

func (v *NullableRange) Set(val *Range) {
	v.value = val
	v.isSet = true
}

func (v NullableRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRange(val *Range) *NullableRange {
	return &NullableRange{value: val, isSet: true}
}

func (v NullableRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}