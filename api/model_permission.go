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

// Permission struct for Permission
type Permission struct {
	KindOfPermission string    `json:"kindOfPermission"`
	Permission       Reference `json:"permission"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission(kindOfPermission string, permission Reference) *Permission {
	this := Permission{}
	this.KindOfPermission = kindOfPermission
	this.Permission = permission
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetKindOfPermission returns the KindOfPermission field value
func (o *Permission) GetKindOfPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KindOfPermission
}

// GetKindOfPermissionOk returns a tuple with the KindOfPermission field value
// and a boolean to check if the value has been set.
func (o *Permission) GetKindOfPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KindOfPermission, true
}

// SetKindOfPermission sets field value
func (o *Permission) SetKindOfPermission(v string) {
	o.KindOfPermission = v
}

// GetPermission returns the Permission field value
func (o *Permission) GetPermission() Reference {
	if o == nil {
		var ret Reference
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *Permission) GetPermissionOk() (*Reference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *Permission) SetPermission(v Reference) {
	o.Permission = v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kindOfPermission"] = o.KindOfPermission
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}