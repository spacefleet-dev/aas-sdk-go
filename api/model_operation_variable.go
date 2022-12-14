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

// OperationVariable struct for OperationVariable
type OperationVariable struct {
	Value ISubmodelElement `json:"value"`
}

// NewOperationVariable instantiates a new OperationVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationVariable(value ISubmodelElement) *OperationVariable {
	this := OperationVariable{}
	this.Value = value
	return &this
}

// NewOperationVariableWithDefaults instantiates a new OperationVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationVariableWithDefaults() *OperationVariable {
	this := OperationVariable{}
	return &this
}

// GetValue returns the Value field value
func (o *OperationVariable) GetValue() ISubmodelElement {
	if o == nil {
		var ret ISubmodelElement
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OperationVariable) GetValueOk() (*ISubmodelElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OperationVariable) SetValue(v ISubmodelElement) {
	o.Value = v
}

func (o OperationVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOperationVariable struct {
	value *OperationVariable
	isSet bool
}

func (v NullableOperationVariable) Get() *OperationVariable {
	return v.value
}

func (v *NullableOperationVariable) Set(val *OperationVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationVariable(val *OperationVariable) *NullableOperationVariable {
	return &NullableOperationVariable{value: val, isSet: true}
}

func (v NullableOperationVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
