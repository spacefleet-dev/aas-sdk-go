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

// SubmodelElementList struct for SubmodelElementList
type SubmodelElementList struct {
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
	OrderRelevant              *bool                       `json:"orderRelevant,omitempty"`
	SemanticIdListElement      *Reference                  `json:"semanticIdListElement,omitempty"`
	TypeValueListElement       AasSubmodelElements         `json:"typeValueListElement"`
	Value                      []ISubmodelElement          `json:"value,omitempty"`
	ValueTypeListElement       *DataTypeDefXsd             `json:"valueTypeListElement,omitempty"`
}

// NewSubmodelElementList instantiates a new SubmodelElementList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmodelElementList(modelType ModelType, typeValueListElement AasSubmodelElements) *SubmodelElementList {
	this := SubmodelElementList{}
	this.ModelType = modelType
	this.TypeValueListElement = typeValueListElement
	return &this
}

// NewSubmodelElementListWithDefaults instantiates a new SubmodelElementList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmodelElementListWithDefaults() *SubmodelElementList {
	this := SubmodelElementList{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *SubmodelElementList) GetExtensions() []Extension {
	if o == nil || o.Extensions == nil {
		var ret []Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetExtensionsOk() ([]Extension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *SubmodelElementList) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []Extension and assigns it to the Extensions field.
func (o *SubmodelElementList) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SubmodelElementList) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SubmodelElementList) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SubmodelElementList) SetCategory(v string) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *SubmodelElementList) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *SubmodelElementList) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *SubmodelElementList) SetChecksum(v string) {
	o.Checksum = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubmodelElementList) GetDescription() []LangString {
	if o == nil || o.Description == nil {
		var ret []LangString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetDescriptionOk() ([]LangString, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubmodelElementList) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LangString and assigns it to the Description field.
func (o *SubmodelElementList) SetDescription(v []LangString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubmodelElementList) GetDisplayName() []LangString {
	if o == nil || o.DisplayName == nil {
		var ret []LangString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetDisplayNameOk() ([]LangString, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubmodelElementList) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LangString and assigns it to the DisplayName field.
func (o *SubmodelElementList) SetDisplayName(v []LangString) {
	o.DisplayName = v
}

// GetIdShort returns the IdShort field value if set, zero value otherwise.
func (o *SubmodelElementList) GetIdShort() string {
	if o == nil || o.IdShort == nil {
		var ret string
		return ret
	}
	return *o.IdShort
}

// GetIdShortOk returns a tuple with the IdShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetIdShortOk() (*string, bool) {
	if o == nil || o.IdShort == nil {
		return nil, false
	}
	return o.IdShort, true
}

// HasIdShort returns a boolean if a field has been set.
func (o *SubmodelElementList) HasIdShort() bool {
	if o != nil && o.IdShort != nil {
		return true
	}

	return false
}

// SetIdShort gets a reference to the given string and assigns it to the IdShort field.
func (o *SubmodelElementList) SetIdShort(v string) {
	o.IdShort = &v
}

// GetModelType returns the ModelType field value
func (o *SubmodelElementList) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *SubmodelElementList) SetModelType(v ModelType) {
	o.ModelType = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SubmodelElementList) GetKind() ModelingKind {
	if o == nil || o.Kind == nil {
		var ret ModelingKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetKindOk() (*ModelingKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SubmodelElementList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given ModelingKind and assigns it to the Kind field.
func (o *SubmodelElementList) SetKind(v ModelingKind) {
	o.Kind = &v
}

// GetSemanticId returns the SemanticId field value if set, zero value otherwise.
func (o *SubmodelElementList) GetSemanticId() Reference {
	if o == nil || o.SemanticId == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticId
}

// GetSemanticIdOk returns a tuple with the SemanticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetSemanticIdOk() (*Reference, bool) {
	if o == nil || o.SemanticId == nil {
		return nil, false
	}
	return o.SemanticId, true
}

// HasSemanticId returns a boolean if a field has been set.
func (o *SubmodelElementList) HasSemanticId() bool {
	if o != nil && o.SemanticId != nil {
		return true
	}

	return false
}

// SetSemanticId gets a reference to the given Reference and assigns it to the SemanticId field.
func (o *SubmodelElementList) SetSemanticId(v Reference) {
	o.SemanticId = &v
}

// GetSupplementalSemanticIds returns the SupplementalSemanticIds field value if set, zero value otherwise.
func (o *SubmodelElementList) GetSupplementalSemanticIds() []Reference {
	if o == nil || o.SupplementalSemanticIds == nil {
		var ret []Reference
		return ret
	}
	return o.SupplementalSemanticIds
}

// GetSupplementalSemanticIdsOk returns a tuple with the SupplementalSemanticIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetSupplementalSemanticIdsOk() ([]Reference, bool) {
	if o == nil || o.SupplementalSemanticIds == nil {
		return nil, false
	}
	return o.SupplementalSemanticIds, true
}

// HasSupplementalSemanticIds returns a boolean if a field has been set.
func (o *SubmodelElementList) HasSupplementalSemanticIds() bool {
	if o != nil && o.SupplementalSemanticIds != nil {
		return true
	}

	return false
}

// SetSupplementalSemanticIds gets a reference to the given []Reference and assigns it to the SupplementalSemanticIds field.
func (o *SubmodelElementList) SetSupplementalSemanticIds(v []Reference) {
	o.SupplementalSemanticIds = v
}

// GetQualifiers returns the Qualifiers field value if set, zero value otherwise.
func (o *SubmodelElementList) GetQualifiers() []Qualifier {
	if o == nil || o.Qualifiers == nil {
		var ret []Qualifier
		return ret
	}
	return o.Qualifiers
}

// GetQualifiersOk returns a tuple with the Qualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetQualifiersOk() ([]Qualifier, bool) {
	if o == nil || o.Qualifiers == nil {
		return nil, false
	}
	return o.Qualifiers, true
}

// HasQualifiers returns a boolean if a field has been set.
func (o *SubmodelElementList) HasQualifiers() bool {
	if o != nil && o.Qualifiers != nil {
		return true
	}

	return false
}

// SetQualifiers gets a reference to the given []Qualifier and assigns it to the Qualifiers field.
func (o *SubmodelElementList) SetQualifiers(v []Qualifier) {
	o.Qualifiers = v
}

// GetEmbeddedDataSpecifications returns the EmbeddedDataSpecifications field value if set, zero value otherwise.
func (o *SubmodelElementList) GetEmbeddedDataSpecifications() []EmbeddedDataSpecification {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		var ret []EmbeddedDataSpecification
		return ret
	}
	return o.EmbeddedDataSpecifications
}

// GetEmbeddedDataSpecificationsOk returns a tuple with the EmbeddedDataSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetEmbeddedDataSpecificationsOk() ([]EmbeddedDataSpecification, bool) {
	if o == nil || o.EmbeddedDataSpecifications == nil {
		return nil, false
	}
	return o.EmbeddedDataSpecifications, true
}

// HasEmbeddedDataSpecifications returns a boolean if a field has been set.
func (o *SubmodelElementList) HasEmbeddedDataSpecifications() bool {
	if o != nil && o.EmbeddedDataSpecifications != nil {
		return true
	}

	return false
}

// SetEmbeddedDataSpecifications gets a reference to the given []EmbeddedDataSpecification and assigns it to the EmbeddedDataSpecifications field.
func (o *SubmodelElementList) SetEmbeddedDataSpecifications(v []EmbeddedDataSpecification) {
	o.EmbeddedDataSpecifications = v
}

// GetOrderRelevant returns the OrderRelevant field value if set, zero value otherwise.
func (o *SubmodelElementList) GetOrderRelevant() bool {
	if o == nil || o.OrderRelevant == nil {
		var ret bool
		return ret
	}
	return *o.OrderRelevant
}

// GetOrderRelevantOk returns a tuple with the OrderRelevant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetOrderRelevantOk() (*bool, bool) {
	if o == nil || o.OrderRelevant == nil {
		return nil, false
	}
	return o.OrderRelevant, true
}

// HasOrderRelevant returns a boolean if a field has been set.
func (o *SubmodelElementList) HasOrderRelevant() bool {
	if o != nil && o.OrderRelevant != nil {
		return true
	}

	return false
}

// SetOrderRelevant gets a reference to the given bool and assigns it to the OrderRelevant field.
func (o *SubmodelElementList) SetOrderRelevant(v bool) {
	o.OrderRelevant = &v
}

// GetSemanticIdListElement returns the SemanticIdListElement field value if set, zero value otherwise.
func (o *SubmodelElementList) GetSemanticIdListElement() Reference {
	if o == nil || o.SemanticIdListElement == nil {
		var ret Reference
		return ret
	}
	return *o.SemanticIdListElement
}

// GetSemanticIdListElementOk returns a tuple with the SemanticIdListElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetSemanticIdListElementOk() (*Reference, bool) {
	if o == nil || o.SemanticIdListElement == nil {
		return nil, false
	}
	return o.SemanticIdListElement, true
}

// HasSemanticIdListElement returns a boolean if a field has been set.
func (o *SubmodelElementList) HasSemanticIdListElement() bool {
	if o != nil && o.SemanticIdListElement != nil {
		return true
	}

	return false
}

// SetSemanticIdListElement gets a reference to the given Reference and assigns it to the SemanticIdListElement field.
func (o *SubmodelElementList) SetSemanticIdListElement(v Reference) {
	o.SemanticIdListElement = &v
}

// GetTypeValueListElement returns the TypeValueListElement field value
func (o *SubmodelElementList) GetTypeValueListElement() AasSubmodelElements {
	if o == nil {
		var ret AasSubmodelElements
		return ret
	}

	return o.TypeValueListElement
}

// GetTypeValueListElementOk returns a tuple with the TypeValueListElement field value
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetTypeValueListElementOk() (*AasSubmodelElements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeValueListElement, true
}

// SetTypeValueListElement sets field value
func (o *SubmodelElementList) SetTypeValueListElement(v AasSubmodelElements) {
	o.TypeValueListElement = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SubmodelElementList) GetValue() []ISubmodelElement {
	if o == nil || o.Value == nil {
		var ret []ISubmodelElement
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetValueOk() ([]ISubmodelElement, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SubmodelElementList) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []ISubmodelElement and assigns it to the Value field.
func (o *SubmodelElementList) SetValue(v []ISubmodelElement) {
	o.Value = v
}

// GetValueTypeListElement returns the ValueTypeListElement field value if set, zero value otherwise.
func (o *SubmodelElementList) GetValueTypeListElement() DataTypeDefXsd {
	if o == nil || o.ValueTypeListElement == nil {
		var ret DataTypeDefXsd
		return ret
	}
	return *o.ValueTypeListElement
}

// GetValueTypeListElementOk returns a tuple with the ValueTypeListElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmodelElementList) GetValueTypeListElementOk() (*DataTypeDefXsd, bool) {
	if o == nil || o.ValueTypeListElement == nil {
		return nil, false
	}
	return o.ValueTypeListElement, true
}

// HasValueTypeListElement returns a boolean if a field has been set.
func (o *SubmodelElementList) HasValueTypeListElement() bool {
	if o != nil && o.ValueTypeListElement != nil {
		return true
	}

	return false
}

// SetValueTypeListElement gets a reference to the given DataTypeDefXsd and assigns it to the ValueTypeListElement field.
func (o *SubmodelElementList) SetValueTypeListElement(v DataTypeDefXsd) {
	o.ValueTypeListElement = &v
}

func (o SubmodelElementList) MarshalJSON() ([]byte, error) {
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
	if o.OrderRelevant != nil {
		toSerialize["orderRelevant"] = o.OrderRelevant
	}
	if o.SemanticIdListElement != nil {
		toSerialize["semanticIdListElement"] = o.SemanticIdListElement
	}
	if true {
		toSerialize["typeValueListElement"] = o.TypeValueListElement
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueTypeListElement != nil {
		toSerialize["valueTypeListElement"] = o.ValueTypeListElement
	}
	return json.Marshal(toSerialize)
}

type NullableSubmodelElementList struct {
	value *SubmodelElementList
	isSet bool
}

func (v NullableSubmodelElementList) Get() *SubmodelElementList {
	return v.value
}

func (v *NullableSubmodelElementList) Set(val *SubmodelElementList) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmodelElementList) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmodelElementList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmodelElementList(val *SubmodelElementList) *NullableSubmodelElementList {
	return &NullableSubmodelElementList{value: val, isSet: true}
}

func (v NullableSubmodelElementList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmodelElementList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
