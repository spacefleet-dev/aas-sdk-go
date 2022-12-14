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

// Result struct for Result
type Result struct {
	Entity      map[string]interface{} `json:"entity,omitempty"`
	EntityType  *string                `json:"entityType,omitempty"`
	IsException *bool                  `json:"isException,omitempty"`
	Messages    []Message              `json:"messages,omitempty"`
	Success     *bool                  `json:"success,omitempty"`
}

// NewResult instantiates a new Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResult() *Result {
	this := Result{}
	return &this
}

// NewResultWithDefaults instantiates a new Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultWithDefaults() *Result {
	this := Result{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *Result) GetEntity() map[string]interface{} {
	if o == nil || o.Entity == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetEntityOk() (map[string]interface{}, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *Result) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *Result) SetEntity(v map[string]interface{}) {
	o.Entity = v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *Result) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *Result) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *Result) SetEntityType(v string) {
	o.EntityType = &v
}

// GetIsException returns the IsException field value if set, zero value otherwise.
func (o *Result) GetIsException() bool {
	if o == nil || o.IsException == nil {
		var ret bool
		return ret
	}
	return *o.IsException
}

// GetIsExceptionOk returns a tuple with the IsException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetIsExceptionOk() (*bool, bool) {
	if o == nil || o.IsException == nil {
		return nil, false
	}
	return o.IsException, true
}

// HasIsException returns a boolean if a field has been set.
func (o *Result) HasIsException() bool {
	if o != nil && o.IsException != nil {
		return true
	}

	return false
}

// SetIsException gets a reference to the given bool and assigns it to the IsException field.
func (o *Result) SetIsException(v bool) {
	o.IsException = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Result) GetMessages() []Message {
	if o == nil || o.Messages == nil {
		var ret []Message
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetMessagesOk() ([]Message, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Result) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []Message and assigns it to the Messages field.
func (o *Result) SetMessages(v []Message) {
	o.Messages = v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *Result) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *Result) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *Result) SetSuccess(v bool) {
	o.Success = &v
}

func (o Result) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.IsException != nil {
		toSerialize["isException"] = o.IsException
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableResult struct {
	value *Result
	isSet bool
}

func (v NullableResult) Get() *Result {
	return v.value
}

func (v *NullableResult) Set(val *Result) {
	v.value = val
	v.isSet = true
}

func (v NullableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResult(val *Result) *NullableResult {
	return &NullableResult{value: val, isSet: true}
}

func (v NullableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
