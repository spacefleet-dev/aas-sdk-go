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

// AccessPermissionRuleAllOf struct for AccessPermissionRuleAllOf
type AccessPermissionRuleAllOf struct {
	PermissionsPerObject    []PermissionsPerObject `json:"permissionsPerObject,omitempty"`
	TargetSubjectAttributes []SubjectAttributes    `json:"targetSubjectAttributes"`
}

// NewAccessPermissionRuleAllOf instantiates a new AccessPermissionRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPermissionRuleAllOf(targetSubjectAttributes []SubjectAttributes) *AccessPermissionRuleAllOf {
	this := AccessPermissionRuleAllOf{}
	this.TargetSubjectAttributes = targetSubjectAttributes
	return &this
}

// NewAccessPermissionRuleAllOfWithDefaults instantiates a new AccessPermissionRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPermissionRuleAllOfWithDefaults() *AccessPermissionRuleAllOf {
	this := AccessPermissionRuleAllOf{}
	return &this
}

// GetPermissionsPerObject returns the PermissionsPerObject field value if set, zero value otherwise.
func (o *AccessPermissionRuleAllOf) GetPermissionsPerObject() []PermissionsPerObject {
	if o == nil || o.PermissionsPerObject == nil {
		var ret []PermissionsPerObject
		return ret
	}
	return o.PermissionsPerObject
}

// GetPermissionsPerObjectOk returns a tuple with the PermissionsPerObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPermissionRuleAllOf) GetPermissionsPerObjectOk() ([]PermissionsPerObject, bool) {
	if o == nil || o.PermissionsPerObject == nil {
		return nil, false
	}
	return o.PermissionsPerObject, true
}

// HasPermissionsPerObject returns a boolean if a field has been set.
func (o *AccessPermissionRuleAllOf) HasPermissionsPerObject() bool {
	if o != nil && o.PermissionsPerObject != nil {
		return true
	}

	return false
}

// SetPermissionsPerObject gets a reference to the given []PermissionsPerObject and assigns it to the PermissionsPerObject field.
func (o *AccessPermissionRuleAllOf) SetPermissionsPerObject(v []PermissionsPerObject) {
	o.PermissionsPerObject = v
}

// GetTargetSubjectAttributes returns the TargetSubjectAttributes field value
func (o *AccessPermissionRuleAllOf) GetTargetSubjectAttributes() []SubjectAttributes {
	if o == nil {
		var ret []SubjectAttributes
		return ret
	}

	return o.TargetSubjectAttributes
}

// GetTargetSubjectAttributesOk returns a tuple with the TargetSubjectAttributes field value
// and a boolean to check if the value has been set.
func (o *AccessPermissionRuleAllOf) GetTargetSubjectAttributesOk() ([]SubjectAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSubjectAttributes, true
}

// SetTargetSubjectAttributes sets field value
func (o *AccessPermissionRuleAllOf) SetTargetSubjectAttributes(v []SubjectAttributes) {
	o.TargetSubjectAttributes = v
}

func (o AccessPermissionRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionsPerObject != nil {
		toSerialize["permissionsPerObject"] = o.PermissionsPerObject
	}
	if true {
		toSerialize["targetSubjectAttributes"] = o.TargetSubjectAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableAccessPermissionRuleAllOf struct {
	value *AccessPermissionRuleAllOf
	isSet bool
}

func (v NullableAccessPermissionRuleAllOf) Get() *AccessPermissionRuleAllOf {
	return v.value
}

func (v *NullableAccessPermissionRuleAllOf) Set(val *AccessPermissionRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPermissionRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPermissionRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPermissionRuleAllOf(val *AccessPermissionRuleAllOf) *NullableAccessPermissionRuleAllOf {
	return &NullableAccessPermissionRuleAllOf{value: val, isSet: true}
}

func (v NullableAccessPermissionRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPermissionRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
