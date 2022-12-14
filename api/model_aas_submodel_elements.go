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
	"fmt"
)

// AasSubmodelElements the model 'AasSubmodelElements'
type AasSubmodelElements string

// List of AasSubmodelElements
const (
	AASSUBMODELELEMENTS_ANNOTATED_RELATIONSHIP_ELEMENT AasSubmodelElements = "AnnotatedRelationshipElement"
	AASSUBMODELELEMENTS_BASIC_EVENT_ELEMENT            AasSubmodelElements = "BasicEventElement"
	AASSUBMODELELEMENTS_BLOB                           AasSubmodelElements = "Blob"
	AASSUBMODELELEMENTS_CAPABILITY                     AasSubmodelElements = "Capability"
	AASSUBMODELELEMENTS_DATA_ELEMENT                   AasSubmodelElements = "DataElement"
	AASSUBMODELELEMENTS_ENTITY                         AasSubmodelElements = "Entity"
	AASSUBMODELELEMENTS_EVENT_ELEMENT                  AasSubmodelElements = "EventElement"
	AASSUBMODELELEMENTS_FILE                           AasSubmodelElements = "File"
	AASSUBMODELELEMENTS_MULTI_LANGUAGE_PROPERTY        AasSubmodelElements = "MultiLanguageProperty"
	AASSUBMODELELEMENTS_OPERATION                      AasSubmodelElements = "Operation"
	AASSUBMODELELEMENTS_PROPERTY                       AasSubmodelElements = "Property"
	AASSUBMODELELEMENTS_RANGE                          AasSubmodelElements = "Range"
	AASSUBMODELELEMENTS_REFERENCE_ELEMENT              AasSubmodelElements = "ReferenceElement"
	AASSUBMODELELEMENTS_RELATIONSHIP_ELEMENT           AasSubmodelElements = "RelationshipElement"
	AASSUBMODELELEMENTS_SUBMODEL_ELEMENT               AasSubmodelElements = "SubmodelElement"
	AASSUBMODELELEMENTS_SUBMODEL_ELEMENT_COLLECTION    AasSubmodelElements = "SubmodelElementCollection"
	AASSUBMODELELEMENTS_SUBMODEL_ELEMENT_LIST          AasSubmodelElements = "SubmodelElementList"
)

// All allowed values of AasSubmodelElements enum
var AllowedAasSubmodelElementsEnumValues = []AasSubmodelElements{
	"AnnotatedRelationshipElement",
	"BasicEventElement",
	"Blob",
	"Capability",
	"DataElement",
	"Entity",
	"EventElement",
	"File",
	"MultiLanguageProperty",
	"Operation",
	"Property",
	"Range",
	"ReferenceElement",
	"RelationshipElement",
	"SubmodelElement",
	"SubmodelElementCollection",
	"SubmodelElementList",
}

func (v *AasSubmodelElements) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AasSubmodelElements(value)
	for _, existing := range AllowedAasSubmodelElementsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AasSubmodelElements", value)
}

// NewAasSubmodelElementsFromValue returns a pointer to a valid AasSubmodelElements
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAasSubmodelElementsFromValue(v string) (*AasSubmodelElements, error) {
	ev := AasSubmodelElements(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AasSubmodelElements: valid values are %v", v, AllowedAasSubmodelElementsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AasSubmodelElements) IsValid() bool {
	for _, existing := range AllowedAasSubmodelElementsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AasSubmodelElements value
func (v AasSubmodelElements) Ptr() *AasSubmodelElements {
	return &v
}

type NullableAasSubmodelElements struct {
	value *AasSubmodelElements
	isSet bool
}

func (v NullableAasSubmodelElements) Get() *AasSubmodelElements {
	return v.value
}

func (v *NullableAasSubmodelElements) Set(val *AasSubmodelElements) {
	v.value = val
	v.isSet = true
}

func (v NullableAasSubmodelElements) IsSet() bool {
	return v.isSet
}

func (v *NullableAasSubmodelElements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAasSubmodelElements(val *AasSubmodelElements) *NullableAasSubmodelElements {
	return &NullableAasSubmodelElements{value: val, isSet: true}
}

func (v NullableAasSubmodelElements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAasSubmodelElements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
