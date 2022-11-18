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

// DataSpecificationIEC61360Content struct for DataSpecificationIEC61360Content
type DataSpecificationIEC61360Content struct {
	Value              *string      `json:"value,omitempty"`
	ValueId            *Reference   `json:"valueId,omitempty"`
	ValueType          *string      `json:"valueType,omitempty"`
	DataType           *string      `json:"dataType,omitempty"`
	Definition         []LangString `json:"definition,omitempty"`
	LevelType          []LevelType  `json:"levelType,omitempty"`
	PreferredName      []LangString `json:"preferredName"`
	ShortName          []LangString `json:"shortName,omitempty"`
	SourceOfDefinition *string      `json:"sourceOfDefinition,omitempty"`
	Symbol             *string      `json:"symbol,omitempty"`
	Unit               *string      `json:"unit,omitempty"`
	UnitId             *Reference   `json:"unitId,omitempty"`
	ValueFormat        *string      `json:"valueFormat,omitempty"`
	ValueList          *ValueList   `json:"valueList,omitempty"`
}

// NewDataSpecificationIEC61360Content instantiates a new DataSpecificationIEC61360Content object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSpecificationIEC61360Content(preferredName []LangString) *DataSpecificationIEC61360Content {
	this := DataSpecificationIEC61360Content{}
	this.PreferredName = preferredName
	return &this
}

// NewDataSpecificationIEC61360ContentWithDefaults instantiates a new DataSpecificationIEC61360Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSpecificationIEC61360ContentWithDefaults() *DataSpecificationIEC61360Content {
	this := DataSpecificationIEC61360Content{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataSpecificationIEC61360Content) SetValue(v string) {
	o.Value = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetValueId() Reference {
	if o == nil || o.ValueId == nil {
		var ret Reference
		return ret
	}
	return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetValueIdOk() (*Reference, bool) {
	if o == nil || o.ValueId == nil {
		return nil, false
	}
	return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasValueId() bool {
	if o != nil && o.ValueId != nil {
		return true
	}

	return false
}

// SetValueId gets a reference to the given Reference and assigns it to the ValueId field.
func (o *DataSpecificationIEC61360Content) SetValueId(v Reference) {
	o.ValueId = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *DataSpecificationIEC61360Content) SetValueType(v string) {
	o.ValueType = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *DataSpecificationIEC61360Content) SetDataType(v string) {
	o.DataType = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetDefinition() []LangString {
	if o == nil || o.Definition == nil {
		var ret []LangString
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetDefinitionOk() ([]LangString, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given []LangString and assigns it to the Definition field.
func (o *DataSpecificationIEC61360Content) SetDefinition(v []LangString) {
	o.Definition = v
}

// GetLevelType returns the LevelType field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetLevelType() []LevelType {
	if o == nil || o.LevelType == nil {
		var ret []LevelType
		return ret
	}
	return o.LevelType
}

// GetLevelTypeOk returns a tuple with the LevelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetLevelTypeOk() ([]LevelType, bool) {
	if o == nil || o.LevelType == nil {
		return nil, false
	}
	return o.LevelType, true
}

// HasLevelType returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasLevelType() bool {
	if o != nil && o.LevelType != nil {
		return true
	}

	return false
}

// SetLevelType gets a reference to the given []LevelType and assigns it to the LevelType field.
func (o *DataSpecificationIEC61360Content) SetLevelType(v []LevelType) {
	o.LevelType = v
}

// GetPreferredName returns the PreferredName field value
func (o *DataSpecificationIEC61360Content) GetPreferredName() []LangString {
	if o == nil {
		var ret []LangString
		return ret
	}

	return o.PreferredName
}

// GetPreferredNameOk returns a tuple with the PreferredName field value
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetPreferredNameOk() ([]LangString, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredName, true
}

// SetPreferredName sets field value
func (o *DataSpecificationIEC61360Content) SetPreferredName(v []LangString) {
	o.PreferredName = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetShortName() []LangString {
	if o == nil || o.ShortName == nil {
		var ret []LangString
		return ret
	}
	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetShortNameOk() ([]LangString, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given []LangString and assigns it to the ShortName field.
func (o *DataSpecificationIEC61360Content) SetShortName(v []LangString) {
	o.ShortName = v
}

// GetSourceOfDefinition returns the SourceOfDefinition field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetSourceOfDefinition() string {
	if o == nil || o.SourceOfDefinition == nil {
		var ret string
		return ret
	}
	return *o.SourceOfDefinition
}

// GetSourceOfDefinitionOk returns a tuple with the SourceOfDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetSourceOfDefinitionOk() (*string, bool) {
	if o == nil || o.SourceOfDefinition == nil {
		return nil, false
	}
	return o.SourceOfDefinition, true
}

// HasSourceOfDefinition returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasSourceOfDefinition() bool {
	if o != nil && o.SourceOfDefinition != nil {
		return true
	}

	return false
}

// SetSourceOfDefinition gets a reference to the given string and assigns it to the SourceOfDefinition field.
func (o *DataSpecificationIEC61360Content) SetSourceOfDefinition(v string) {
	o.SourceOfDefinition = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *DataSpecificationIEC61360Content) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DataSpecificationIEC61360Content) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitId returns the UnitId field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetUnitId() Reference {
	if o == nil || o.UnitId == nil {
		var ret Reference
		return ret
	}
	return *o.UnitId
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetUnitIdOk() (*Reference, bool) {
	if o == nil || o.UnitId == nil {
		return nil, false
	}
	return o.UnitId, true
}

// HasUnitId returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasUnitId() bool {
	if o != nil && o.UnitId != nil {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given Reference and assigns it to the UnitId field.
func (o *DataSpecificationIEC61360Content) SetUnitId(v Reference) {
	o.UnitId = &v
}

// GetValueFormat returns the ValueFormat field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetValueFormat() string {
	if o == nil || o.ValueFormat == nil {
		var ret string
		return ret
	}
	return *o.ValueFormat
}

// GetValueFormatOk returns a tuple with the ValueFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetValueFormatOk() (*string, bool) {
	if o == nil || o.ValueFormat == nil {
		return nil, false
	}
	return o.ValueFormat, true
}

// HasValueFormat returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasValueFormat() bool {
	if o != nil && o.ValueFormat != nil {
		return true
	}

	return false
}

// SetValueFormat gets a reference to the given string and assigns it to the ValueFormat field.
func (o *DataSpecificationIEC61360Content) SetValueFormat(v string) {
	o.ValueFormat = &v
}

// GetValueList returns the ValueList field value if set, zero value otherwise.
func (o *DataSpecificationIEC61360Content) GetValueList() ValueList {
	if o == nil || o.ValueList == nil {
		var ret ValueList
		return ret
	}
	return *o.ValueList
}

// GetValueListOk returns a tuple with the ValueList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSpecificationIEC61360Content) GetValueListOk() (*ValueList, bool) {
	if o == nil || o.ValueList == nil {
		return nil, false
	}
	return o.ValueList, true
}

// HasValueList returns a boolean if a field has been set.
func (o *DataSpecificationIEC61360Content) HasValueList() bool {
	if o != nil && o.ValueList != nil {
		return true
	}

	return false
}

// SetValueList gets a reference to the given ValueList and assigns it to the ValueList field.
func (o *DataSpecificationIEC61360Content) SetValueList(v ValueList) {
	o.ValueList = &v
}

func (o DataSpecificationIEC61360Content) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueId != nil {
		toSerialize["valueId"] = o.ValueId
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	if o.LevelType != nil {
		toSerialize["levelType"] = o.LevelType
	}
	if true {
		toSerialize["preferredName"] = o.PreferredName
	}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.SourceOfDefinition != nil {
		toSerialize["sourceOfDefinition"] = o.SourceOfDefinition
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitId != nil {
		toSerialize["unitId"] = o.UnitId
	}
	if o.ValueFormat != nil {
		toSerialize["valueFormat"] = o.ValueFormat
	}
	if o.ValueList != nil {
		toSerialize["valueList"] = o.ValueList
	}
	return json.Marshal(toSerialize)
}

type NullableDataSpecificationIEC61360Content struct {
	value *DataSpecificationIEC61360Content
	isSet bool
}

func (v NullableDataSpecificationIEC61360Content) Get() *DataSpecificationIEC61360Content {
	return v.value
}

func (v *NullableDataSpecificationIEC61360Content) Set(val *DataSpecificationIEC61360Content) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSpecificationIEC61360Content) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSpecificationIEC61360Content) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSpecificationIEC61360Content(val *DataSpecificationIEC61360Content) *NullableDataSpecificationIEC61360Content {
	return &NullableDataSpecificationIEC61360Content{value: val, isSet: true}
}

func (v NullableDataSpecificationIEC61360Content) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSpecificationIEC61360Content) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
