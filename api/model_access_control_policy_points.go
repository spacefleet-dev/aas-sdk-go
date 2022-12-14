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

// AccessControlPolicyPoints struct for AccessControlPolicyPoints
type AccessControlPolicyPoints struct {
	PolicyAdministrationPoint PolicyAdministrationPoint `json:"policyAdministrationPoint"`
	PolicyDecisionPoint       PolicyDecisionPoint       `json:"policyDecisionPoint"`
	PolicyEnforcementPoint    PolicyEnforcementPoint    `json:"policyEnforcementPoint"`
	PolicyInformationPoints   *PolicyInformationPoints  `json:"policyInformationPoints,omitempty"`
}

// NewAccessControlPolicyPoints instantiates a new AccessControlPolicyPoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessControlPolicyPoints(policyAdministrationPoint PolicyAdministrationPoint, policyDecisionPoint PolicyDecisionPoint, policyEnforcementPoint PolicyEnforcementPoint) *AccessControlPolicyPoints {
	this := AccessControlPolicyPoints{}
	this.PolicyAdministrationPoint = policyAdministrationPoint
	this.PolicyDecisionPoint = policyDecisionPoint
	this.PolicyEnforcementPoint = policyEnforcementPoint
	return &this
}

// NewAccessControlPolicyPointsWithDefaults instantiates a new AccessControlPolicyPoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessControlPolicyPointsWithDefaults() *AccessControlPolicyPoints {
	this := AccessControlPolicyPoints{}
	return &this
}

// GetPolicyAdministrationPoint returns the PolicyAdministrationPoint field value
func (o *AccessControlPolicyPoints) GetPolicyAdministrationPoint() PolicyAdministrationPoint {
	if o == nil {
		var ret PolicyAdministrationPoint
		return ret
	}

	return o.PolicyAdministrationPoint
}

// GetPolicyAdministrationPointOk returns a tuple with the PolicyAdministrationPoint field value
// and a boolean to check if the value has been set.
func (o *AccessControlPolicyPoints) GetPolicyAdministrationPointOk() (*PolicyAdministrationPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyAdministrationPoint, true
}

// SetPolicyAdministrationPoint sets field value
func (o *AccessControlPolicyPoints) SetPolicyAdministrationPoint(v PolicyAdministrationPoint) {
	o.PolicyAdministrationPoint = v
}

// GetPolicyDecisionPoint returns the PolicyDecisionPoint field value
func (o *AccessControlPolicyPoints) GetPolicyDecisionPoint() PolicyDecisionPoint {
	if o == nil {
		var ret PolicyDecisionPoint
		return ret
	}

	return o.PolicyDecisionPoint
}

// GetPolicyDecisionPointOk returns a tuple with the PolicyDecisionPoint field value
// and a boolean to check if the value has been set.
func (o *AccessControlPolicyPoints) GetPolicyDecisionPointOk() (*PolicyDecisionPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyDecisionPoint, true
}

// SetPolicyDecisionPoint sets field value
func (o *AccessControlPolicyPoints) SetPolicyDecisionPoint(v PolicyDecisionPoint) {
	o.PolicyDecisionPoint = v
}

// GetPolicyEnforcementPoint returns the PolicyEnforcementPoint field value
func (o *AccessControlPolicyPoints) GetPolicyEnforcementPoint() PolicyEnforcementPoint {
	if o == nil {
		var ret PolicyEnforcementPoint
		return ret
	}

	return o.PolicyEnforcementPoint
}

// GetPolicyEnforcementPointOk returns a tuple with the PolicyEnforcementPoint field value
// and a boolean to check if the value has been set.
func (o *AccessControlPolicyPoints) GetPolicyEnforcementPointOk() (*PolicyEnforcementPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyEnforcementPoint, true
}

// SetPolicyEnforcementPoint sets field value
func (o *AccessControlPolicyPoints) SetPolicyEnforcementPoint(v PolicyEnforcementPoint) {
	o.PolicyEnforcementPoint = v
}

// GetPolicyInformationPoints returns the PolicyInformationPoints field value if set, zero value otherwise.
func (o *AccessControlPolicyPoints) GetPolicyInformationPoints() PolicyInformationPoints {
	if o == nil || o.PolicyInformationPoints == nil {
		var ret PolicyInformationPoints
		return ret
	}
	return *o.PolicyInformationPoints
}

// GetPolicyInformationPointsOk returns a tuple with the PolicyInformationPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessControlPolicyPoints) GetPolicyInformationPointsOk() (*PolicyInformationPoints, bool) {
	if o == nil || o.PolicyInformationPoints == nil {
		return nil, false
	}
	return o.PolicyInformationPoints, true
}

// HasPolicyInformationPoints returns a boolean if a field has been set.
func (o *AccessControlPolicyPoints) HasPolicyInformationPoints() bool {
	if o != nil && o.PolicyInformationPoints != nil {
		return true
	}

	return false
}

// SetPolicyInformationPoints gets a reference to the given PolicyInformationPoints and assigns it to the PolicyInformationPoints field.
func (o *AccessControlPolicyPoints) SetPolicyInformationPoints(v PolicyInformationPoints) {
	o.PolicyInformationPoints = &v
}

func (o AccessControlPolicyPoints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policyAdministrationPoint"] = o.PolicyAdministrationPoint
	}
	if true {
		toSerialize["policyDecisionPoint"] = o.PolicyDecisionPoint
	}
	if true {
		toSerialize["policyEnforcementPoint"] = o.PolicyEnforcementPoint
	}
	if o.PolicyInformationPoints != nil {
		toSerialize["policyInformationPoints"] = o.PolicyInformationPoints
	}
	return json.Marshal(toSerialize)
}

type NullableAccessControlPolicyPoints struct {
	value *AccessControlPolicyPoints
	isSet bool
}

func (v NullableAccessControlPolicyPoints) Get() *AccessControlPolicyPoints {
	return v.value
}

func (v *NullableAccessControlPolicyPoints) Set(val *AccessControlPolicyPoints) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessControlPolicyPoints) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessControlPolicyPoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessControlPolicyPoints(val *AccessControlPolicyPoints) *NullableAccessControlPolicyPoints {
	return &NullableAccessControlPolicyPoints{value: val, isSet: true}
}

func (v NullableAccessControlPolicyPoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessControlPolicyPoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
