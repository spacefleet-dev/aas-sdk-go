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

// Submodel struct for Submodel
type Submodel struct {
	Extensions                 []Extension                 `json:"extensions,omitempty"`
	Category                   *string                     `json:"category,omitempty"`
	Checksum                   *string                     `json:"checksum,omitempty"`
	Description                []LangString                `json:"description,omitempty"`
	DisplayName                []LangString                `json:"displayName,omitempty"`
	IdShort                    *string                     `json:"idShort,omitempty"`
	ModelType                  ModelType                   `json:"modelType"`
	Administration             *AdministrativeInformation  `json:"administration,omitempty"`
	Identification             Identifier                  `json:"identification"`
	Kind                       *ModelingKind               `json:"kind,omitempty"`
	SemanticId                 *Reference                  `json:"semanticId,omitempty"`
	SupplementalSemanticIds    []Reference                 `json:"supplementalSemanticIds,omitempty"`
	Qualifiers                 []Qualifier                 `json:"qualifiers,omitempty"`
	EmbeddedDataSpecifications []EmbeddedDataSpecification `json:"embeddedDataSpecifications,omitempty"`
	SubmodelElements           []ISubmodelElement          `json:"submodelElements,omitempty"`
}

// NewSubmodel instantiates a new Submodel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmodel(modelType ModelType, identification Identifier) *Submodel {
	this := Submodel{}
	this.ModelType = modelType
	this.Identification = identification
	return &this
}

// NewSubmodelWithDefaults instantiates a new Submodel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmodelWithDefaults() *Submodel {
	this := Submodel{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Submodel) GetExtensions() []Extension {
	if o == nil || o.Extensions == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Submodel) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *Submodel) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Submodel) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Submodel) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Submodel) SetCategory(v string) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *Submodel) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *Submodel) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *Submodel) SetChecksum(v string) {
	o.Checksum = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Submodel) GetDescription() []LangString {
	if o == nil || o.Description == nil {
		var ret []LangString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetDescriptionOk() ([]LangString, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Submodel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LangString and assigns it to the Description field.
func (o *Submodel) SetDescription(v []LangString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Submodel) GetDisplayName() []LangString {
	if o == nil || o.DisplayName == nil {
		var ret []LangString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetDisplayNameOk() ([]LangString, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Submodel) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LangString and assigns it to the DisplayName field.
func (o *Submodel) SetDisplayName(v []LangString) {
	o.DisplayName = v
}

// GetIdShort returns the IdShort field value if set, zero value otherwise.
func (o *Submodel) GetIdShort() string {
	if o == nil || o.IdShort == nil {
		var ret string
		return ret
	}
	return *o.IdShort
}

// GetIdShortOk returns a tuple with the IdShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetIdShortOk() (*string, bool) {
	if o == nil || o.IdShort == nil {
		return nil, false
	}
	return o.IdShort, true
}

// HasIdShort returns a boolean if a field has been set.
func (o *Submodel) HasIdShort() bool {
	if o != nil && o.IdShort != nil {
		return true
	}

	return false
}

// SetIdShort gets a reference to the given string and assigns it to the IdShort field.
func (o *Submodel) SetIdShort(v string) {
	o.IdShort = &v
}

// GetModelType returns the ModelType field value
func (o *Submodel) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *Submodel) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *Submodel) SetModelType(v ModelType) {
	o.ModelType = v
}

// GetAdministration returns the Administration field value if set, zero value otherwise.
func (o *Submodel) GetAdministration() AdministrativeInformation {
	if o == nil || o.Administration == nil {
		var ret AdministrativeInformation
		return ret
	}
	return *o.Administration
}

// GetAdministrationOk returns a tuple with the Administration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetAdministrationOk() (*AdministrativeInformation, bool) {
	if o == nil || o.Administration == nil {
		return nil, false
	}
	return o.Administration, true
}

// HasAdministration returns a boolean if a field has been set.
func (o *Submodel) HasAdministration() bool {
	if o != nil && o.Administration != nil {
		return true
	}

	return false
}

// SetAdministration gets a reference to the given AdministrativeInformation and assigns it to the Administration field.
func (o *Submodel) SetAdministration(v AdministrativeInformation) {
	o.Administration = &v
}

// GetIdentification returns the Identification field value
func (o *Submodel) GetIdentification() Identifier {
	if o == nil {
		var ret Identifier
		return ret
	}

	return o.Identification
}

// GetIdentificationOk returns a tuple with the Identification field value
// and a boolean to check if the value has been set.
func (o *Submodel) GetIdentificationOk() (*Identifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identification, true
}

// SetIdentification sets field value
func (o *Submodel) SetIdentification(v Identifier) {
	o.Identification = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Submodel) GetKind() ModelingKind {
	if o == nil || o.Kind == nil {
		var ret ModelingKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetKindOk() (*ModelingKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Submodel) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ModelingKind and assigns it to the Kind field.
func (o *Submodel) SetKind(v ModelingKind) {
	o.Kind = &v
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *Submodel) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *Submodel) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *Submodel) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *Submodel) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *Submodel) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *Submodel) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetQualifiers returns the Qualifiers field value if set, zero value otherwise.
func (o *Submodel) GetQualifiers() []Qualifier {
	if o == nil || o.Qualifiers == nil {
		var ret []Qualifier
		return ret
	}
	return o.Qualifiers
}

// GetQualifiersOk returns a tuple with the Qualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetQualifiersOk() ([]Qualifier, bool) {
	if o == nil || o.Qualifiers == nil {
		return nil, false
	}
	return o.Qualifiers, true
}

// HasQualifiers returns a boolean if a field has been set.
func (o *Submodel) HasQualifiers() bool {
	if o != nil && o.Qualifiers != nil {
		return true
	}

	return false
}

// SetQualifiers gets a reference to the given []Qualifier and assigns it to the Qualifiers field.
func (o *Submodel) SetQualifiers(v []Qualifier) {
	o.Qualifiers = v
}

// GetEmbeddedDataSpecifications returns the EmbeddedDataSpecifications field value if set, zero value otherwise.
func (o *Submodel) GetEmbeddedDataSpecifications() []EmbeddedDataSpecification {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		var ret []EmbeddedDataSpecification
		return ret
	}
	return o.EmbeddedDataSpecifications
}

// GetEmbeddedDataSpecificationsOk returns a tuple with the EmbeddedDataSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetEmbeddedDataSpecificationsOk() ([]EmbeddedDataSpecification, bool) {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		return nil, false
	}
	return o.EmbeddedDataSpecifications, true
}

// HasEmbeddedDataSpecifications returns a boolean if a field has been set.
func (o *Submodel) HasEmbeddedDataSpecifications() bool {
	if o != nil && o.EmbeddedDataSpecifications != nil {
		return true
	}

	return false
}

// SetEmbeddedDataSpecifications gets a reference to the given []EmbeddedDataSpecification and assigns it to the EmbeddedDataSpecifications field.
func (o *Submodel) SetEmbeddedDataSpecifications(v []EmbeddedDataSpecification) {
	o.EmbeddedDataSpecifications = v
}

// GetSubmodelElements returns the SubmodelElements field value if set, zero value otherwise.
func (o *Submodel) GetSubmodelElements() []ISubmodelElement {
	if o == nil || o.SubmodelElements == nil {
		var ret []ISubmodelElement
		return ret
	}
	return o.SubmodelElements
}

// GetSubmodelElementsOk returns a tuple with the SubmodelElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submodel) GetSubmodelElementsOk() ([]ISubmodelElement, bool) {
	if o == nil || o.SubmodelElements == nil {
		return nil, false
	}
	return o.SubmodelElements, true
}

// HasSubmodelElements returns a boolean if a field has been set.
func (o *Submodel) HasSubmodelElements() bool {
	if o != nil && o.SubmodelElements != nil {
		return true
	}

	return false
}

// SetSubmodelElements gets a reference to the given []ISubmodelElement and assigns it to the SubmodelElements field.
func (o *Submodel) SetSubmodelElements(v []ISubmodelElement) {
	o.SubmodelElements = v
}

func (o Submodel) MarshalJSON() ([]byte, error) {
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
	if o.Administration != nil {
		toSerialize["administration"] = o.Administration
	}
	if true {
		toSerialize["identification"] = o.Identification
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
	if o.SubmodelElements != nil {
		toSerialize["submodelElements"] = o.SubmodelElements
	}
	return json.Marshal(toSerialize)
}

type NullableSubmodel struct {
	value *Submodel
	isSet bool
}

func (v NullableSubmodel) Get() *Submodel {
	return v.value
}

func (v *NullableSubmodel) Set(val *Submodel) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmodel) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmodel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmodel(val *Submodel) *NullableSubmodel {
	return &NullableSubmodel{value: val, isSet: true}
}

func (v NullableSubmodel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmodel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
