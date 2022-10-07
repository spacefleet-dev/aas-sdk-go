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

// DataSpecificationPhysicalUnit struct for DataSpecificationPhysicalUnit
type DataSpecificationPhysicalUnit struct {
	ModelType               ModelType    `json:"modelType"`
	ConversionFactor        *string      `json:"conversionFactor,omitempty"`
	Definition              []LangString `json:"definition"`
	DinNotation             *string      `json:"dinNotation,omitempty"`
	EceCode                 *string      `json:"eceCode,omitempty"`
	EceName                 *string      `json:"eceName,omitempty"`
	NistName                *string      `json:"nistName,omitempty"`
	RegistrationAuthorityId *string      `json:"registrationAuthorityId,omitempty"`
	SiName                  *string      `json:"siName,omitempty"`
	SiNotation              *string      `json:"siNotation,omitempty"`
	SourceOfDefinition      *string      `json:"sourceOfDefinition,omitempty"`
	Supplier                *string      `json:"supplier,omitempty"`
	UnitName                string       `json:"unitName"`
	UnitSymbol              string       `json:"unitSymbol"`
}

// NewDataSpecificationPhysicalUnit instantiates a new DataSpecificationPhysicalUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSpecificationPhysicalUnit(modelType ModelType, definition []LangString, unitName string, unitSymbol string) *DataSpecificationPhysicalUnit {
	this := DataSpecificationPhysicalUnit{}
	this.ModelType = modelType
	this.Definition = definition
	this.UnitName = unitName
	this.UnitSymbol = unitSymbol
	return &this
}

// NewDataSpecificationPhysicalUnitWithDefaults instantiates a new DataSpecificationPhysicalUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSpecificationPhysicalUnitWithDefaults() *DataSpecificationPhysicalUnit {
	this := DataSpecificationPhysicalUnit{}
	return &this
}

// GetModelType returns the ModelType field value
func (o *DataSpecificationPhysicalUnit) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *DataSpecificationPhysicalUnit) SetModelType(v ModelType) {
	o.ModelType = v
}

// GetConversionFactor returns the ConversionFactor field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetConversionFactor() string {
	if o == nil || o.ConversionFactor == nil {
		var ret string
		return ret
	}
	return *o.ConversionFactor
}

// GetConversionFactorOk returns a tuple with the ConversionFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetConversionFactorOk() (*string, bool) {
	if o == nil || o.ConversionFactor == nil {
		return nil, false
	}
	return o.ConversionFactor, true
}

// HasConversionFactor returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasConversionFactor() bool {
	if o != nil && o.ConversionFactor != nil {
		return true
	}

	return false
}

// SetConversionFactor gets a reference to the given string and assigns it to the ConversionFactor field.
func (o *DataSpecificationPhysicalUnit) SetConversionFactor(v string) {
	o.ConversionFactor = &v
}

// GetDefinition returns the Definition field value
func (o *DataSpecificationPhysicalUnit) GetDefinition() []LangString {
	if o == nil {
		var ret []LangString
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetDefinitionOk() ([]LangString, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definition, true
}

// SetDefinition sets field value
func (o *DataSpecificationPhysicalUnit) SetDefinition(v []LangString) {
	o.Definition = v
}

// GetDinNotation returns the DinNotation field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetDinNotation() string {
	if o == nil || o.DinNotation == nil {
		var ret string
		return ret
	}
	return *o.DinNotation
}

// GetDinNotationOk returns a tuple with the DinNotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetDinNotationOk() (*string, bool) {
	if o == nil || o.DinNotation == nil {
		return nil, false
	}
	return o.DinNotation, true
}

// HasDinNotation returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasDinNotation() bool {
	if o != nil && o.DinNotation != nil {
		return true
	}

	return false
}

// SetDinNotation gets a reference to the given string and assigns it to the DinNotation field.
func (o *DataSpecificationPhysicalUnit) SetDinNotation(v string) {
	o.DinNotation = &v
}

// GetEceCode returns the EceCode field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetEceCode() string {
	if o == nil || o.EceCode == nil {
		var ret string
		return ret
	}
	return *o.EceCode
}

// GetEceCodeOk returns a tuple with the EceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetEceCodeOk() (*string, bool) {
	if o == nil || o.EceCode == nil {
		return nil, false
	}
	return o.EceCode, true
}

// HasEceCode returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasEceCode() bool {
	if o != nil && o.EceCode != nil {
		return true
	}

	return false
}

// SetEceCode gets a reference to the given string and assigns it to the EceCode field.
func (o *DataSpecificationPhysicalUnit) SetEceCode(v string) {
	o.EceCode = &v
}

// GetEceName returns the EceName field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetEceName() string {
	if o == nil || o.EceName == nil {
		var ret string
		return ret
	}
	return *o.EceName
}

// GetEceNameOk returns a tuple with the EceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetEceNameOk() (*string, bool) {
	if o == nil || o.EceName == nil {
		return nil, false
	}
	return o.EceName, true
}

// HasEceName returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasEceName() bool {
	if o != nil && o.EceName != nil {
		return true
	}

	return false
}

// SetEceName gets a reference to the given string and assigns it to the EceName field.
func (o *DataSpecificationPhysicalUnit) SetEceName(v string) {
	o.EceName = &v
}

// GetNistName returns the NistName field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetNistName() string {
	if o == nil || o.NistName == nil {
		var ret string
		return ret
	}
	return *o.NistName
}

// GetNistNameOk returns a tuple with the NistName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetNistNameOk() (*string, bool) {
	if o == nil || o.NistName == nil {
		return nil, false
	}
	return o.NistName, true
}

// HasNistName returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasNistName() bool {
	if o != nil && o.NistName != nil {
		return true
	}

	return false
}

// SetNistName gets a reference to the given string and assigns it to the NistName field.
func (o *DataSpecificationPhysicalUnit) SetNistName(v string) {
	o.NistName = &v
}

// GetRegistrationAuthorityId returns the RegistrationAuthorityId field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetRegistrationAuthorityId() string {
	if o == nil || o.RegistrationAuthorityId == nil {
		var ret string
		return ret
	}
	return *o.RegistrationAuthorityId
}

// GetRegistrationAuthorityIdOk returns a tuple with the RegistrationAuthorityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetRegistrationAuthorityIdOk() (*string, bool) {
	if o == nil || o.RegistrationAuthorityId == nil {
		return nil, false
	}
	return o.RegistrationAuthorityId, true
}

// HasRegistrationAuthorityId returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasRegistrationAuthorityId() bool {
	if o != nil && o.RegistrationAuthorityId != nil {
		return true
	}

	return false
}

// SetRegistrationAuthorityId gets a reference to the given string and assigns it to the RegistrationAuthorityId field.
func (o *DataSpecificationPhysicalUnit) SetRegistrationAuthorityId(v string) {
	o.RegistrationAuthorityId = &v
}

// GetSiName returns the SiName field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetSiName() string {
	if o == nil || o.SiName == nil {
		var ret string
		return ret
	}
	return *o.SiName
}

// GetSiNameOk returns a tuple with the SiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetSiNameOk() (*string, bool) {
	if o == nil || o.SiName == nil {
		return nil, false
	}
	return o.SiName, true
}

// HasSiName returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasSiName() bool {
	if o != nil && o.SiName != nil {
		return true
	}

	return false
}

// SetSiName gets a reference to the given string and assigns it to the SiName field.
func (o *DataSpecificationPhysicalUnit) SetSiName(v string) {
	o.SiName = &v
}

// GetSiNotation returns the SiNotation field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetSiNotation() string {
	if o == nil || o.SiNotation == nil {
		var ret string
		return ret
	}
	return *o.SiNotation
}

// GetSiNotationOk returns a tuple with the SiNotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetSiNotationOk() (*string, bool) {
	if o == nil || o.SiNotation == nil {
		return nil, false
	}
	return o.SiNotation, true
}

// HasSiNotation returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasSiNotation() bool {
	if o != nil && o.SiNotation != nil {
		return true
	}

	return false
}

// SetSiNotation gets a reference to the given string and assigns it to the SiNotation field.
func (o *DataSpecificationPhysicalUnit) SetSiNotation(v string) {
	o.SiNotation = &v
}

// GetSourceOfDefinition returns the SourceOfDefinition field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetSourceOfDefinition() string {
	if o == nil || o.SourceOfDefinition == nil {
		var ret string
		return ret
	}
	return *o.SourceOfDefinition
}

// GetSourceOfDefinitionOk returns a tuple with the SourceOfDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetSourceOfDefinitionOk() (*string, bool) {
	if o == nil || o.SourceOfDefinition == nil {
		return nil, false
	}
	return o.SourceOfDefinition, true
}

// HasSourceOfDefinition returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasSourceOfDefinition() bool {
	if o != nil && o.SourceOfDefinition != nil {
		return true
	}

	return false
}

// SetSourceOfDefinition gets a reference to the given string and assigns it to the SourceOfDefinition field.
func (o *DataSpecificationPhysicalUnit) SetSourceOfDefinition(v string) {
	o.SourceOfDefinition = &v
}

// GetSupplier returns the Supplier field value if set, zero value otherwise.
func (o *DataSpecificationPhysicalUnit) GetSupplier() string {
	if o == nil || o.Supplier == nil {
		var ret string
		return ret
	}
	return *o.Supplier
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetSupplierOk() (*string, bool) {
	if o == nil || o.Supplier == nil {
		return nil, false
	}
	return o.Supplier, true
}

// HasSupplier returns a boolean if a field has been set.
func (o *DataSpecificationPhysicalUnit) HasSupplier() bool {
	if o != nil && o.Supplier != nil {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given string and assigns it to the Supplier field.
func (o *DataSpecificationPhysicalUnit) SetSupplier(v string) {
	o.Supplier = &v
}

// GetUnitName returns the UnitName field value
func (o *DataSpecificationPhysicalUnit) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value
func (o *DataSpecificationPhysicalUnit) SetUnitName(v string) {
	o.UnitName = v
}

// GetUnitSymbol returns the UnitSymbol field value
func (o *DataSpecificationPhysicalUnit) GetUnitSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitSymbol
}

// GetUnitSymbolOk returns a tuple with the UnitSymbol field value
// and a boolean to check if the value has been set.
func (o *DataSpecificationPhysicalUnit) GetUnitSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitSymbol, true
}

// SetUnitSymbol sets field value
func (o *DataSpecificationPhysicalUnit) SetUnitSymbol(v string) {
	o.UnitSymbol = v
}

func (o DataSpecificationPhysicalUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["modelType"] = o.ModelType
	}
	if o.ConversionFactor != nil {
		toSerialize["conversionFactor"] = o.ConversionFactor
	}
	if true {
		toSerialize["definition"] = o.Definition
	}
	if o.DinNotation != nil {
		toSerialize["dinNotation"] = o.DinNotation
	}
	if o.EceCode != nil {
		toSerialize["eceCode"] = o.EceCode
	}
	if o.EceName != nil {
		toSerialize["eceName"] = o.EceName
	}
	if o.NistName != nil {
		toSerialize["nistName"] = o.NistName
	}
	if o.RegistrationAuthorityId != nil {
		toSerialize["registrationAuthorityId"] = o.RegistrationAuthorityId
	}
	if o.SiName != nil {
		toSerialize["siName"] = o.SiName
	}
	if o.SiNotation != nil {
		toSerialize["siNotation"] = o.SiNotation
	}
	if o.SourceOfDefinition != nil {
		toSerialize["sourceOfDefinition"] = o.SourceOfDefinition
	}
	if o.Supplier != nil {
		toSerialize["supplier"] = o.Supplier
	}
	if true {
		toSerialize["unitName"] = o.UnitName
	}
	if true {
		toSerialize["unitSymbol"] = o.UnitSymbol
	}
	return json.Marshal(toSerialize)
}

type NullableDataSpecificationPhysicalUnit struct {
	value *DataSpecificationPhysicalUnit
	isSet bool
}

func (v NullableDataSpecificationPhysicalUnit) Get() *DataSpecificationPhysicalUnit {
	return v.value
}

func (v *NullableDataSpecificationPhysicalUnit) Set(val *DataSpecificationPhysicalUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSpecificationPhysicalUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSpecificationPhysicalUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSpecificationPhysicalUnit(val *DataSpecificationPhysicalUnit) *NullableDataSpecificationPhysicalUnit {
	return &NullableDataSpecificationPhysicalUnit{value: val, isSet: true}
}

func (v NullableDataSpecificationPhysicalUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSpecificationPhysicalUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
