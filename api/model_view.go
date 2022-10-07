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

// View struct for View
type View struct {
	Extensions                 []Extension                 `json:"extensions,omitempty"`
	Category                   *string                     `json:"category,omitempty"`
	Checksum                   *string                     `json:"checksum,omitempty"`
	Description                []LangString                `json:"description,omitempty"`
	DisplayName                []LangString                `json:"displayName,omitempty"`
	IdShort                    *string                     `json:"idShort,omitempty"`
	ModelType                  ModelType                   `json:"modelType"`
	EmbeddedDataSpecifications []EmbeddedDataSpecification `json:"embeddedDataSpecifications,omitempty"`
	SemanticId                 *Reference                  `json:"semanticId,omitempty"`
	SupplementalSemanticIds    []Reference                 `json:"supplementalSemanticIds,omitempty"`
	ContainedElements          []Reference                 `json:"containedElements,omitempty"`
}

// NewView instantiates a new View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewView(modelType ModelType) *View {
	this := View{}
	this.ModelType = modelType
	return &this
}

// NewViewWithDefaults instantiates a new View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *View) GetExtensions() []Extension {
	if o == nil || o.Extensions == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *View) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *View) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *View) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *View) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *View) SetCategory(v string) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *View) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *View) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *View) SetChecksum(v string) {
	o.Checksum = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *View) GetDescription() []LangString {
	if o == nil || o.Description == nil {
		var ret []LangString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetDescriptionOk() ([]LangString, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *View) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LangString and assigns it to the Description field.
func (o *View) SetDescription(v []LangString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *View) GetDisplayName() []LangString {
	if o == nil || o.DisplayName == nil {
		var ret []LangString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetDisplayNameOk() ([]LangString, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *View) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LangString and assigns it to the DisplayName field.
func (o *View) SetDisplayName(v []LangString) {
	o.DisplayName = v
}

// GetIdShort returns the IdShort field value if set, zero value otherwise.
func (o *View) GetIdShort() string {
	if o == nil || o.IdShort == nil {
		var ret string
		return ret
	}
	return *o.IdShort
}

// GetIdShortOk returns a tuple with the IdShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetIdShortOk() (*string, bool) {
	if o == nil || o.IdShort == nil {
		return nil, false
	}
	return o.IdShort, true
}

// HasIdShort returns a boolean if a field has been set.
func (o *View) HasIdShort() bool {
	if o != nil && o.IdShort != nil {
		return true
	}

	return false
}

// SetIdShort gets a reference to the given string and assigns it to the IdShort field.
func (o *View) SetIdShort(v string) {
	o.IdShort = &v
}

// GetModelType returns the ModelType field value
func (o *View) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *View) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *View) SetModelType(v ModelType) {
	o.ModelType = v
}

// GetEmbeddedDataSpecifications returns the EmbeddedDataSpecifications field value if set, zero value otherwise.
func (o *View) GetEmbeddedDataSpecifications() []EmbeddedDataSpecification {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		var ret []EmbeddedDataSpecification
		return ret
	}
	return o.EmbeddedDataSpecifications
}

// GetEmbeddedDataSpecificationsOk returns a tuple with the EmbeddedDataSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetEmbeddedDataSpecificationsOk() ([]EmbeddedDataSpecification, bool) {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		return nil, false
	}
	return o.EmbeddedDataSpecifications, true
}

// HasEmbeddedDataSpecifications returns a boolean if a field has been set.
func (o *View) HasEmbeddedDataSpecifications() bool {
	if o != nil && o.EmbeddedDataSpecifications != nil {
		return true
	}

	return false
}

// SetEmbeddedDataSpecifications gets a reference to the given []EmbeddedDataSpecification and assigns it to the EmbeddedDataSpecifications field.
func (o *View) SetEmbeddedDataSpecifications(v []EmbeddedDataSpecification) {
	o.EmbeddedDataSpecifications = v
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *View) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *View) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *View) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *View) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *View) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *View) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetContainedElements returns the ContainedElements field value if set, zero value otherwise.
func (o *View) GetContainedElements() []Reference {
	if o == nil || o.ContainedElements == nil {
		var ret []Reference
		return ret
	}
	return o.ContainedElements
}

// GetContainedElementsOk returns a tuple with the ContainedElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetContainedElementsOk() ([]Reference, bool) {
	if o == nil || o.ContainedElements == nil {
		return nil, false
	}
	return o.ContainedElements, true
}

// HasContainedElements returns a boolean if a field has been set.
func (o *View) HasContainedElements() bool {
	if o != nil && o.ContainedElements != nil {
		return true
	}

	return false
}

// SetContainedElements gets a reference to the given []Reference and assigns it to the ContainedElements field.
func (o *View) SetContainedElements(v []Reference) {
	o.ContainedElements = v
}

func (o View) MarshalJSON() ([]byte, error) {
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
	if o.EmbeddedDataSpecifications != nil {
		toSerialize["embeddedDataSpecifications"] = o.EmbeddedDataSpecifications
	}
	if o.SemanticId != nil {
		toSerialize["semanticId"] = o.SemanticId
	}
	if o.SupplementalSemanticIds != nil {
		toSerialize["supplementalSemanticIds"] = o.SupplementalSemanticIds
	}
	if o.ContainedElements != nil {
		toSerialize["containedElements"] = o.ContainedElements
	}
	return json.Marshal(toSerialize)
}

type NullableView struct {
	value *View
	isSet bool
}

func (v NullableView) Get() *View {
	return v.value
}

func (v *NullableView) Set(val *View) {
	v.value = val
	v.isSet = true
}

func (v NullableView) IsSet() bool {
	return v.isSet
}

func (v *NullableView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableView(val *View) *NullableView {
	return &NullableView{value: val, isSet: true}
}

func (v NullableView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
