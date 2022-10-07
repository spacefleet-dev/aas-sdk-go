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

// ExtensionAllOf struct for ExtensionAllOf
type ExtensionAllOf struct {
	Name      string          `json:"name"`
	RefersTo  *Reference      `json:"refersTo,omitempty"`
	Value     *string         `json:"value,omitempty"`
	ValueType *DataTypeDefXsd `json:"valueType,omitempty"`
}

// NewExtensionAllOf instantiates a new ExtensionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionAllOf(name string) *ExtensionAllOf {
	this := ExtensionAllOf{}
	this.Name = name
	return &this
}

// NewExtensionAllOfWithDefaults instantiates a new ExtensionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionAllOfWithDefaults() *ExtensionAllOf {
	this := ExtensionAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *ExtensionAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtensionAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtensionAllOf) SetName(v string) {
	o.Name = v
}

// GetRefersTo returns the RefersTo field value if set, zero value otherwise.
func (o *ExtensionAllOf) GetRefersTo() Reference {
	if o == nil || o.RefersTo == nil {
		var ret Reference
		return ret
	}
	return *o.RefersTo
}

// GetRefersToOk returns a tuple with the RefersTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAllOf) GetRefersToOk() (*Reference, bool) {
	if o == nil || o.RefersTo == nil {
		return nil, false
	}
	return o.RefersTo, true
}

// HasRefersTo returns a boolean if a field has been set.
func (o *ExtensionAllOf) HasRefersTo() bool {
	if o != nil && o.RefersTo != nil {
		return true
	}

	return false
}

// SetRefersTo gets a reference to the given Reference and assigns it to the RefersTo field.
func (o *ExtensionAllOf) SetRefersTo(v Reference) {
	o.RefersTo = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ExtensionAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ExtensionAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ExtensionAllOf) SetValue(v string) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ExtensionAllOf) GetValueType() DataTypeDefXsd {
	if o == nil || o.ValueType == nil {
		var ret DataTypeDefXsd
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAllOf) GetValueTypeOk() (*DataTypeDefXsd, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ExtensionAllOf) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given DataTypeDefXsd and assigns it to the ValueType field.
func (o *ExtensionAllOf) SetValueType(v DataTypeDefXsd) {
	o.ValueType = &v
}

func (o ExtensionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RefersTo != nil {
		toSerialize["refersTo"] = o.RefersTo
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionAllOf struct {
	value *ExtensionAllOf
	isSet bool
}

func (v NullableExtensionAllOf) Get() *ExtensionAllOf {
	return v.value
}

func (v *NullableExtensionAllOf) Set(val *ExtensionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionAllOf(val *ExtensionAllOf) *NullableExtensionAllOf {
	return &NullableExtensionAllOf{value: val, isSet: true}
}

func (v NullableExtensionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
