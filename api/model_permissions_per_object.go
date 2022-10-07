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

// PermissionsPerObject struct for PermissionsPerObject
type PermissionsPerObject struct {
	Object                 *Reference        `json:"object,omitempty"`
	Permission             []Permission      `json:"permission,omitempty"`
	TargetObjectAttributes *ObjectAttributes `json:"targetObjectAttributes,omitempty"`
}

// NewPermissionsPerObject instantiates a new PermissionsPerObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsPerObject() *PermissionsPerObject {
	this := PermissionsPerObject{}
	return &this
}

// NewPermissionsPerObjectWithDefaults instantiates a new PermissionsPerObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsPerObjectWithDefaults() *PermissionsPerObject {
	this := PermissionsPerObject{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *PermissionsPerObject) GetObject() Reference {
	if o == nil || o.Object == nil {
		var ret Reference
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsPerObject) GetObjectOk() (*Reference, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *PermissionsPerObject) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given Reference and assigns it to the Object field.
func (o *PermissionsPerObject) SetObject(v Reference) {
	o.Object = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *PermissionsPerObject) GetPermission() []Permission {
	if o == nil || o.Permission == nil {
		var ret []Permission
		return ret
	}
	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsPerObject) GetPermissionOk() ([]Permission, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *PermissionsPerObject) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given []Permission and assigns it to the Permission field.
func (o *PermissionsPerObject) SetPermission(v []Permission) {
	o.Permission = v
}

// GetTargetObjectAttributes returns the TargetObjectAttributes field value if set, zero value otherwise.
func (o *PermissionsPerObject) GetTargetObjectAttributes() ObjectAttributes {
	if o == nil || o.TargetObjectAttributes == nil {
		var ret ObjectAttributes
		return ret
	}
	return *o.TargetObjectAttributes
}

// GetTargetObjectAttributesOk returns a tuple with the TargetObjectAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsPerObject) GetTargetObjectAttributesOk() (*ObjectAttributes, bool) {
	if o == nil || o.TargetObjectAttributes == nil {
		return nil, false
	}
	return o.TargetObjectAttributes, true
}

// HasTargetObjectAttributes returns a boolean if a field has been set.
func (o *PermissionsPerObject) HasTargetObjectAttributes() bool {
	if o != nil && o.TargetObjectAttributes != nil {
		return true
	}

	return false
}

// SetTargetObjectAttributes gets a reference to the given ObjectAttributes and assigns it to the TargetObjectAttributes field.
func (o *PermissionsPerObject) SetTargetObjectAttributes(v ObjectAttributes) {
	o.TargetObjectAttributes = &v
}

func (o PermissionsPerObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.TargetObjectAttributes != nil {
		toSerialize["targetObjectAttributes"] = o.TargetObjectAttributes
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionsPerObject struct {
	value *PermissionsPerObject
	isSet bool
}

func (v NullablePermissionsPerObject) Get() *PermissionsPerObject {
	return v.value
}

func (v *NullablePermissionsPerObject) Set(val *PermissionsPerObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsPerObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsPerObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsPerObject(val *PermissionsPerObject) *NullablePermissionsPerObject {
	return &NullablePermissionsPerObject{value: val, isSet: true}
}

func (v NullablePermissionsPerObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsPerObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
