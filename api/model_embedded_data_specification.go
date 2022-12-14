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

// EmbeddedDataSpecification struct for EmbeddedDataSpecification
type EmbeddedDataSpecification struct {
	DataSpecification        Reference                `json:"dataSpecification"`
	DataSpecificationContent DataSpecificationContent `json:"dataSpecificationContent"`
}

// NewEmbeddedDataSpecification instantiates a new EmbeddedDataSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedDataSpecification(dataSpecification Reference, dataSpecificationContent DataSpecificationContent) *EmbeddedDataSpecification {
	this := EmbeddedDataSpecification{}
	this.DataSpecification = dataSpecification
	this.DataSpecificationContent = dataSpecificationContent
	return &this
}

// NewEmbeddedDataSpecificationWithDefaults instantiates a new EmbeddedDataSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedDataSpecificationWithDefaults() *EmbeddedDataSpecification {
	this := EmbeddedDataSpecification{}
	return &this
}

// GetDataSpecification returns the DataSpecification field value
func (o *EmbeddedDataSpecification) GetDataSpecification() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.DataSpecification
}

// GetDataSpecificationOk returns a tuple with the DataSpecification field value
// and a boolean to check if the value has been set.
func (o *EmbeddedDataSpecification) GetDataSpecificationOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSpecification, true
}

// SetDataSpecification sets field value
func (o *EmbeddedDataSpecification) SetDataSpecification(v Reference) {
	o.DataSpecification = v
}

// GetDataSpecificationContent returns the DataSpecificationContent field value
func (o *EmbeddedDataSpecification) GetDataSpecificationContent() DataSpecificationContent {
	if o == nil {
		var ret DataSpecificationContent
		return ret
	}

	return o.DataSpecificationContent
}

// GetDataSpecificationContentOk returns a tuple with the DataSpecificationContent field value
// and a boolean to check if the value has been set.
func (o *EmbeddedDataSpecification) GetDataSpecificationContentOk() (*DataSpecificationContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSpecificationContent, true
}

// SetDataSpecificationContent sets field value
func (o *EmbeddedDataSpecification) SetDataSpecificationContent(v DataSpecificationContent) {
	o.DataSpecificationContent = v
}

func (o EmbeddedDataSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSpecification"] = o.DataSpecification
	}
	if true {
		toSerialize["dataSpecificationContent"] = o.DataSpecificationContent
	}
	return json.Marshal(toSerialize)
}

type NullableEmbeddedDataSpecification struct {
	value *EmbeddedDataSpecification
	isSet bool
}

func (v NullableEmbeddedDataSpecification) Get() *EmbeddedDataSpecification {
	return v.value
}

func (v *NullableEmbeddedDataSpecification) Set(val *EmbeddedDataSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedDataSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedDataSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedDataSpecification(val *EmbeddedDataSpecification) *NullableEmbeddedDataSpecification {
	return &NullableEmbeddedDataSpecification{value: val, isSet: true}
}

func (v NullableEmbeddedDataSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedDataSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
