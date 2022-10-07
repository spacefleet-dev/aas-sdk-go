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

// BlobCertificateAllOf struct for BlobCertificateAllOf
type BlobCertificateAllOf struct {
	BlobCertificate    *Blob       `json:"blobCertificate,omitempty"`
	ContainedExtension []Reference `json:"containedExtension,omitempty"`
	LastCertificate    *bool       `json:"lastCertificate,omitempty"`
}

// NewBlobCertificateAllOf instantiates a new BlobCertificateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobCertificateAllOf() *BlobCertificateAllOf {
	this := BlobCertificateAllOf{}
	return &this
}

// NewBlobCertificateAllOfWithDefaults instantiates a new BlobCertificateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobCertificateAllOfWithDefaults() *BlobCertificateAllOf {
	this := BlobCertificateAllOf{}
	return &this
}

// GetBlobCertificate returns the BlobCertificate field value if set, zero value otherwise.
func (o *BlobCertificateAllOf) GetBlobCertificate() Blob {
	if o == nil || o.BlobCertificate == nil {
		var ret Blob
		return ret
	}
	return *o.BlobCertificate
}

// GetBlobCertificateOk returns a tuple with the BlobCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobCertificateAllOf) GetBlobCertificateOk() (*Blob, bool) {
	if o == nil || o.BlobCertificate == nil {
		return nil, false
	}
	return o.BlobCertificate, true
}

// HasBlobCertificate returns a boolean if a field has been set.
func (o *BlobCertificateAllOf) HasBlobCertificate() bool {
	if o != nil && o.BlobCertificate != nil {
		return true
	}

	return false
}

// SetBlobCertificate gets a reference to the given Blob and assigns it to the BlobCertificate field.
func (o *BlobCertificateAllOf) SetBlobCertificate(v Blob) {
	o.BlobCertificate = &v
}

// GetContainedExtension returns the ContainedExtension field value if set, zero value otherwise.
func (o *BlobCertificateAllOf) GetContainedExtension() []Reference {
	if o == nil || o.ContainedExtension == nil {
		var ret []Reference
		return ret
	}
	return o.ContainedExtension
}

// GetContainedExtensionOk returns a tuple with the ContainedExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobCertificateAllOf) GetContainedExtensionOk() ([]Reference, bool) {
	if o == nil || o.ContainedExtension == nil {
		return nil, false
	}
	return o.ContainedExtension, true
}

// HasContainedExtension returns a boolean if a field has been set.
func (o *BlobCertificateAllOf) HasContainedExtension() bool {
	if o != nil && o.ContainedExtension != nil {
		return true
	}

	return false
}

// SetContainedExtension gets a reference to the given []Reference and assigns it to the ContainedExtension field.
func (o *BlobCertificateAllOf) SetContainedExtension(v []Reference) {
	o.ContainedExtension = v
}

// GetLastCertificate returns the LastCertificate field value if set, zero value otherwise.
func (o *BlobCertificateAllOf) GetLastCertificate() bool {
	if o == nil || o.LastCertificate == nil {
		var ret bool
		return ret
	}
	return *o.LastCertificate
}

// GetLastCertificateOk returns a tuple with the LastCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobCertificateAllOf) GetLastCertificateOk() (*bool, bool) {
	if o == nil || o.LastCertificate == nil {
		return nil, false
	}
	return o.LastCertificate, true
}

// HasLastCertificate returns a boolean if a field has been set.
func (o *BlobCertificateAllOf) HasLastCertificate() bool {
	if o != nil && o.LastCertificate != nil {
		return true
	}

	return false
}

// SetLastCertificate gets a reference to the given bool and assigns it to the LastCertificate field.
func (o *BlobCertificateAllOf) SetLastCertificate(v bool) {
	o.LastCertificate = &v
}

func (o BlobCertificateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlobCertificate != nil {
		toSerialize["blobCertificate"] = o.BlobCertificate
	}
	if o.ContainedExtension != nil {
		toSerialize["containedExtension"] = o.ContainedExtension
	}
	if o.LastCertificate != nil {
		toSerialize["lastCertificate"] = o.LastCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableBlobCertificateAllOf struct {
	value *BlobCertificateAllOf
	isSet bool
}

func (v NullableBlobCertificateAllOf) Get() *BlobCertificateAllOf {
	return v.value
}

func (v *NullableBlobCertificateAllOf) Set(val *BlobCertificateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobCertificateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobCertificateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobCertificateAllOf(val *BlobCertificateAllOf) *NullableBlobCertificateAllOf {
	return &NullableBlobCertificateAllOf{value: val, isSet: true}
}

func (v NullableBlobCertificateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobCertificateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
