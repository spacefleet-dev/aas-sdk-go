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

// SecurityCertificateInner - struct for SecurityCertificateInner
type SecurityCertificateInner struct {
	BlobCertificate *BlobCertificate
}

// BlobCertificateAsSecurityCertificateInner is a convenience function that returns BlobCertificate wrapped in SecurityCertificateInner
func BlobCertificateAsSecurityCertificateInner(v *BlobCertificate) SecurityCertificateInner {
	return SecurityCertificateInner{
		BlobCertificate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityCertificateInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlobCertificate
	err = json.Unmarshal(data, &dst.BlobCertificate)
	if err == nil {
		jsonBlobCertificate, _ := json.Marshal(dst.BlobCertificate)
		if string(jsonBlobCertificate) == "{}" { // empty struct
			dst.BlobCertificate = nil
		} else {
			match++
		}
	} else {
		dst.BlobCertificate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlobCertificate = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SecurityCertificateInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SecurityCertificateInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityCertificateInner) MarshalJSON() ([]byte, error) {
	if src.BlobCertificate != nil {
		return json.Marshal(&src.BlobCertificate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityCertificateInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BlobCertificate != nil {
		return obj.BlobCertificate
	}

	// all schemas are nil
	return nil
}

type NullableSecurityCertificateInner struct {
	value *SecurityCertificateInner
	isSet bool
}

func (v NullableSecurityCertificateInner) Get() *SecurityCertificateInner {
	return v.value
}

func (v *NullableSecurityCertificateInner) Set(val *SecurityCertificateInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityCertificateInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityCertificateInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityCertificateInner(val *SecurityCertificateInner) *NullableSecurityCertificateInner {
	return &NullableSecurityCertificateInner{value: val, isSet: true}
}

func (v NullableSecurityCertificateInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityCertificateInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
