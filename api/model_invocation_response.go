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

// InvocationResponse struct for InvocationResponse
type InvocationResponse struct {
	ExecutionState    *string             `json:"executionState,omitempty"`
	InoutputArguments []OperationVariable `json:"inoutputArguments,omitempty"`
	OperationResult   *Result             `json:"operationResult,omitempty"`
	OutputArguments   []OperationVariable `json:"outputArguments,omitempty"`
	RequestId         *string             `json:"requestId,omitempty"`
}

// NewInvocationResponse instantiates a new InvocationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvocationResponse() *InvocationResponse {
	this := InvocationResponse{}
	return &this
}

// NewInvocationResponseWithDefaults instantiates a new InvocationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvocationResponseWithDefaults() *InvocationResponse {
	this := InvocationResponse{}
	return &this
}

// GetExecutionState returns the ExecutionState field value if set, zero value otherwise.
func (o *InvocationResponse) GetExecutionState() string {
	if o == nil || o.ExecutionState == nil {
		var ret string
		return ret
	}
	return *o.ExecutionState
}

// GetExecutionStateOk returns a tuple with the ExecutionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResponse) GetExecutionStateOk() (*string, bool) {
	if o == nil || o.ExecutionState == nil {
		return nil, false
	}
	return o.ExecutionState, true
}

// HasExecutionState returns a boolean if a field has been set.
func (o *InvocationResponse) HasExecutionState() bool {
	if o != nil && o.ExecutionState != nil {
		return true
	}

	return false
}

// SetExecutionState gets a reference to the given string and assigns it to the ExecutionState field.
func (o *InvocationResponse) SetExecutionState(v string) {
	o.ExecutionState = &v
}

// GetInoutputArguments returns the InoutputArguments field value if set, zero value otherwise.
func (o *InvocationResponse) GetInoutputArguments() []OperationVariable {
	if o == nil || o.InoutputArguments == nil {
		var ret []OperationVariable
		return ret
	}
	return o.InoutputArguments
}

// GetInoutputArgumentsOk returns a tuple with the InoutputArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResponse) GetInoutputArgumentsOk() ([]OperationVariable, bool) {
	if o == nil || o.InoutputArguments == nil {
		return nil, false
	}
	return o.InoutputArguments, true
}

// HasInoutputArguments returns a boolean if a field has been set.
func (o *InvocationResponse) HasInoutputArguments() bool {
	if o != nil && o.InoutputArguments != nil {
		return true
	}

	return false
}

// SetInoutputArguments gets a reference to the given []OperationVariable and assigns it to the InoutputArguments field.
func (o *InvocationResponse) SetInoutputArguments(v []OperationVariable) {
	o.InoutputArguments = v
}

// GetOperationResult returns the OperationResult field value if set, zero value otherwise.
func (o *InvocationResponse) GetOperationResult() Result {
	if o == nil || o.OperationResult == nil {
		var ret Result
		return ret
	}
	return *o.OperationResult
}

// GetOperationResultOk returns a tuple with the OperationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResponse) GetOperationResultOk() (*Result, bool) {
	if o == nil || o.OperationResult == nil {
		return nil, false
	}
	return o.OperationResult, true
}

// HasOperationResult returns a boolean if a field has been set.
func (o *InvocationResponse) HasOperationResult() bool {
	if o != nil && o.OperationResult != nil {
		return true
	}

	return false
}

// SetOperationResult gets a reference to the given Result and assigns it to the OperationResult field.
func (o *InvocationResponse) SetOperationResult(v Result) {
	o.OperationResult = &v
}

// GetOutputArguments returns the OutputArguments field value if set, zero value otherwise.
func (o *InvocationResponse) GetOutputArguments() []OperationVariable {
	if o == nil || o.OutputArguments == nil {
		var ret []OperationVariable
		return ret
	}
	return o.OutputArguments
}

// GetOutputArgumentsOk returns a tuple with the OutputArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResponse) GetOutputArgumentsOk() ([]OperationVariable, bool) {
	if o == nil || o.OutputArguments == nil {
		return nil, false
	}
	return o.OutputArguments, true
}

// HasOutputArguments returns a boolean if a field has been set.
func (o *InvocationResponse) HasOutputArguments() bool {
	if o != nil && o.OutputArguments != nil {
		return true
	}

	return false
}

// SetOutputArguments gets a reference to the given []OperationVariable and assigns it to the OutputArguments field.
func (o *InvocationResponse) SetOutputArguments(v []OperationVariable) {
	o.OutputArguments = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *InvocationResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *InvocationResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *InvocationResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o InvocationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExecutionState != nil {
		toSerialize["executionState"] = o.ExecutionState
	}
	if o.InoutputArguments != nil {
		toSerialize["inoutputArguments"] = o.InoutputArguments
	}
	if o.OperationResult != nil {
		toSerialize["operationResult"] = o.OperationResult
	}
	if o.OutputArguments != nil {
		toSerialize["outputArguments"] = o.OutputArguments
	}
	if o.RequestId != nil {
		toSerialize["requestId"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableInvocationResponse struct {
	value *InvocationResponse
	isSet bool
}

func (v NullableInvocationResponse) Get() *InvocationResponse {
	return v.value
}

func (v *NullableInvocationResponse) Set(val *InvocationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvocationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvocationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvocationResponse(val *InvocationResponse) *NullableInvocationResponse {
	return &NullableInvocationResponse{value: val, isSet: true}
}

func (v NullableInvocationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvocationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
