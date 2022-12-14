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

// KeyTypes the model 'KeyTypes'
type KeyTypes string

// List of KeyTypes
const (
	KEYTYPES_ANNOTATED_RELATIONSHIP_ELEMENT KeyTypes = "AnnotatedRelationshipElement"
	KEYTYPES_ASSET_ADMINISTRATION_SHELL     KeyTypes = "AssetAdministrationShell"
	KEYTYPES_BASIC_EVENT_ELEMENT            KeyTypes = "BasicEventElement"
	KEYTYPES_BLOB                           KeyTypes = "Blob"
	KEYTYPES_CAPABILITY                     KeyTypes = "Capability"
	KEYTYPES_CONCEPT_DESCRIPTION            KeyTypes = "ConceptDescription"
	KEYTYPES_DATA_ELEMENT                   KeyTypes = "DataElement"
	KEYTYPES_ENTITY                         KeyTypes = "Entity"
	KEYTYPES_EVENT_ELEMENT                  KeyTypes = "EventElement"
	KEYTYPES_FILE                           KeyTypes = "File"
	KEYTYPES_FRAGMENT_REFERENCE             KeyTypes = "FragmentReference"
	KEYTYPES_GLOBAL_REFERENCE               KeyTypes = "GlobalReference"
	KEYTYPES_IDENTIFIABLE                   KeyTypes = "Identifiable"
	KEYTYPES_MULTI_LANGUAGE_PROPERTY        KeyTypes = "MultiLanguageProperty"
	KEYTYPES_OPERATION                      KeyTypes = "Operation"
	KEYTYPES_PROPERTY                       KeyTypes = "Property"
	KEYTYPES_RANGE                          KeyTypes = "Range"
	KEYTYPES_REFERABLE                      KeyTypes = "Referable"
	KEYTYPES_REFERENCE_ELEMENT              KeyTypes = "ReferenceElement"
	KEYTYPES_RELATIONSHIP_ELEMENT           KeyTypes = "RelationshipElement"
	KEYTYPES_SUBMODEL                       KeyTypes = "Submodel"
	KEYTYPES_SUBMODEL_ELEMENT               KeyTypes = "SubmodelElement"
	KEYTYPES_SUBMODEL_ELEMENT_COLLECTION    KeyTypes = "SubmodelElementCollection"
	KEYTYPES_SUBMODEL_ELEMENT_LIST          KeyTypes = "SubmodelElementList"
)

// All allowed values of KeyTypes enum
var AllowedKeyTypesEnumValues = []KeyTypes{
	"AnnotatedRelationshipElement",
	"AssetAdministrationShell",
	"BasicEventElement",
	"Blob",
	"Capability",
	"ConceptDescription",
	"DataElement",
	"Entity",
	"EventElement",
	"File",
	"FragmentReference",
	"GlobalReference",
	"Identifiable",
	"MultiLanguageProperty",
	"Operation",
	"Property",
	"Range",
	"Referable",
	"ReferenceElement",
	"RelationshipElement",
	"Submodel",
	"SubmodelElement",
	"SubmodelElementCollection",
	"SubmodelElementList",
}

func (v *KeyTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyTypes(value)
	for _, existing := range AllowedKeyTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyTypes", value)
}

// NewKeyTypesFromValue returns a pointer to a valid KeyTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyTypesFromValue(v string) (*KeyTypes, error) {
	ev := KeyTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyTypes: valid values are %v", v, AllowedKeyTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyTypes) IsValid() bool {
	for _, existing := range AllowedKeyTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyTypes value
func (v KeyTypes) Ptr() *KeyTypes {
	return &v
}

type NullableKeyTypes struct {
	value *KeyTypes
	isSet bool
}

func (v NullableKeyTypes) Get() *KeyTypes {
	return v.value
}

func (v *NullableKeyTypes) Set(val *KeyTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyTypes(val *KeyTypes) *NullableKeyTypes {
	return &NullableKeyTypes{value: val, isSet: true}
}

func (v NullableKeyTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
