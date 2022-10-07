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

// DataSpecificationContent struct for DataSpecificationContent
type DataSpecificationContent struct {
	ModelType ModelType `json:"modelType"`
}

// NewDataSpecificationContent instantiates a new DataSpecificationContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSpecificationContent(modelType ModelType) *DataSpecificationContent {
	this := DataSpecificationContent{}
	this.ModelType = modelType
	return &this
}

// NewDataSpecificationContentWithDefaults instantiates a new DataSpecificationContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSpecificationContentWithDefaults() *DataSpecificationContent {
	this := DataSpecificationContent{}
	return &this
}

// GetModelType returns the ModelType field value
func (o *DataSpecificationContent) GetModelType() ModelType {
	if o == nil {
		var ret ModelType
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *DataSpecificationContent) GetModelTypeOk() (*ModelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *DataSpecificationContent) SetModelType(v ModelType) {
	o.ModelType = v
}

func (o DataSpecificationContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["modelType"] = o.ModelType
	}
	return json.Marshal(toSerialize)
}

type NullableDataSpecificationContent struct {
	value *DataSpecificationContent
	isSet bool
}

func (v NullableDataSpecificationContent) Get() *DataSpecificationContent {
	return v.value
}

func (v *NullableDataSpecificationContent) Set(val *DataSpecificationContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSpecificationContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSpecificationContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSpecificationContent(val *DataSpecificationContent) *NullableDataSpecificationContent {
	return &NullableDataSpecificationContent{value: val, isSet: true}
}

func (v NullableDataSpecificationContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSpecificationContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
