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

// FileAllOf struct for FileAllOf
type FileAllOf struct {
	ContentType string  `json:"contentType"`
	Value       *string `json:"value,omitempty"`
}

// NewFileAllOf instantiates a new FileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAllOf(contentType string) *FileAllOf {
	this := FileAllOf{}
	this.ContentType = contentType
	return &this
}

// NewFileAllOfWithDefaults instantiates a new FileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAllOfWithDefaults() *FileAllOf {
	this := FileAllOf{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *FileAllOf) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *FileAllOf) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *FileAllOf) SetContentType(v string) {
	o.ContentType = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FileAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FileAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FileAllOf) SetValue(v string) {
	o.Value = &v
}

func (o FileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFileAllOf struct {
	value *FileAllOf
	isSet bool
}

func (v NullableFileAllOf) Get() *FileAllOf {
	return v.value
}

func (v *NullableFileAllOf) Set(val *FileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAllOf(val *FileAllOf) *NullableFileAllOf {
	return &NullableFileAllOf{value: val, isSet: true}
}

func (v NullableFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
