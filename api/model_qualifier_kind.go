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
	"fmt"
)

// QualifierKind the model 'QualifierKind'
type QualifierKind string

// List of QualifierKind
const (
	QUALIFIERKIND_CONCEPT_QUALIFIER  QualifierKind = "ConceptQualifier"
	QUALIFIERKIND_TEMPLATE_QUALIFIER QualifierKind = "TemplateQualifier"
	QUALIFIERKIND_VALUE_QUALIFIER    QualifierKind = "ValueQualifier"
)

// All allowed values of QualifierKind enum
var AllowedQualifierKindEnumValues = []QualifierKind{
	"ConceptQualifier",
	"TemplateQualifier",
	"ValueQualifier",
}

func (v *QualifierKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QualifierKind(value)
	for _, existing := range AllowedQualifierKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QualifierKind", value)
}

// NewQualifierKindFromValue returns a pointer to a valid QualifierKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQualifierKindFromValue(v string) (*QualifierKind, error) {
	ev := QualifierKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QualifierKind: valid values are %v", v, AllowedQualifierKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QualifierKind) IsValid() bool {
	for _, existing := range AllowedQualifierKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QualifierKind value
func (v QualifierKind) Ptr() *QualifierKind {
	return &v
}

type NullableQualifierKind struct {
	value *QualifierKind
	isSet bool
}

func (v NullableQualifierKind) Get() *QualifierKind {
	return v.value
}

func (v *NullableQualifierKind) Set(val *QualifierKind) {
	v.value = val
	v.isSet = true
}

func (v NullableQualifierKind) IsSet() bool {
	return v.isSet
}

func (v *NullableQualifierKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualifierKind(val *QualifierKind) *NullableQualifierKind {
	return &NullableQualifierKind{value: val, isSet: true}
}

func (v NullableQualifierKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualifierKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}