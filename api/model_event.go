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

// Event struct for Event
type Event struct {
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
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(modelType ModelType) *Event {
	this := Event{}
	this.ModelType = modelType
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Event) GetExtensions() []Extension {
	if o == nil || o.Extensions == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Event) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *Event) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Event) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Event) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Event) SetCategory(v string) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *Event) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *Event) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *Event) SetChecksum(v string) {
	o.Checksum = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Event) GetDescription() []LangString {
	if o == nil || o.Description == nil {
		var ret []LangString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDescriptionOk() ([]LangString, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Event) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LangString and assigns it to the Description field.
func (o *Event) SetDescription(v []LangString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Event) GetDisplayName() []LangString {
	if o == nil || o.DisplayName == nil {
		var ret []LangString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDisplayNameOk() ([]LangString, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Event) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LangString and assigns it to the DisplayName field.
func (o *Event) SetDisplayName(v []LangString) {
	o.DisplayName = v
}

// GetIdShort returns the IdShort field value if set, zero value otherwise.
func (o *Event) GetIdShort() string {
	if o == nil || o.IdShort == nil {
		var ret string
		return ret
	}
	return *o.IdShort
}

// GetIdShortOk returns a tuple with the IdShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIdShortOk() (*string, bool) {
	if o == nil || o.IdShort == nil {
		return nil, false
	}
	return o.IdShort, true
}

// HasIdShort returns a boolean if a field has been set.
func (o *Event) HasIdShort() bool {
	if o != nil && o.IdShort != nil {
		return true
	}

	return false
}

// SetIdShort gets a reference to the given string and assigns it to the IdShort field.
func (o *Event) SetIdShort(v string) {
	o.IdShort = &v
}

// GetModelType returns the ModelType field value
func (o *Event) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *Event) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *Event) SetModelType(v ModelType) {
	o.ModelType = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Event) GetKind() ModelingKind {
	if o == nil || o.Kind == nil {
		var ret ModelingKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetKindOk() (*ModelingKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Event) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ModelingKind and assigns it to the Kind field.
func (o *Event) SetKind(v ModelingKind) {
	o.Kind = &v
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *Event) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *Event) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *Event) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *Event) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *Event) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *Event) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetQualifiers returns the Qualifiers field value if set, zero value otherwise.
func (o *Event) GetQualifiers() []Qualifier {
	if o == nil || o.Qualifiers == nil {
		var ret []Qualifier
		return ret
	}
	return o.Qualifiers
}

// GetQualifiersOk returns a tuple with the Qualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetQualifiersOk() ([]Qualifier, bool) {
	if o == nil || o.Qualifiers == nil {
		return nil, false
	}
	return o.Qualifiers, true
}

// HasQualifiers returns a boolean if a field has been set.
func (o *Event) HasQualifiers() bool {
	if o != nil && o.Qualifiers != nil {
		return true
	}

	return false
}

// SetQualifiers gets a reference to the given []Qualifier and assigns it to the Qualifiers field.
func (o *Event) SetQualifiers(v []Qualifier) {
	o.Qualifiers = v
}

// GetEmbeddedDataSpecifications returns the EmbeddedDataSpecifications field value if set, zero value otherwise.
func (o *Event) GetEmbeddedDataSpecifications() []EmbeddedDataSpecification {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		var ret []EmbeddedDataSpecification
		return ret
	}
	return o.EmbeddedDataSpecifications
}

// GetEmbeddedDataSpecificationsOk returns a tuple with the EmbeddedDataSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEmbeddedDataSpecificationsOk() ([]EmbeddedDataSpecification, bool) {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		return nil, false
	}
	return o.EmbeddedDataSpecifications, true
}

// HasEmbeddedDataSpecifications returns a boolean if a field has been set.
func (o *Event) HasEmbeddedDataSpecifications() bool {
	if o != nil && o.EmbeddedDataSpecifications != nil {
		return true
	}

	return false
}

// SetEmbeddedDataSpecifications gets a reference to the given []EmbeddedDataSpecification and assigns it to the EmbeddedDataSpecifications field.
func (o *Event) SetEmbeddedDataSpecifications(v []EmbeddedDataSpecification) {
	o.EmbeddedDataSpecifications = v
}

func (o Event) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}