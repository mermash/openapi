/*
session.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SessionNothing struct for SessionNothing
type SessionNothing struct {
	Dummy *bool `json:"dummy,omitempty"`
}

// NewSessionNothing instantiates a new SessionNothing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionNothing() *SessionNothing {
	this := SessionNothing{}
	return &this
}

// NewSessionNothingWithDefaults instantiates a new SessionNothing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionNothingWithDefaults() *SessionNothing {
	this := SessionNothing{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise.
func (o *SessionNothing) GetDummy() bool {
	if o == nil || o.Dummy == nil {
		var ret bool
		return ret
	}
	return *o.Dummy
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionNothing) GetDummyOk() (*bool, bool) {
	if o == nil || o.Dummy == nil {
		return nil, false
	}
	return o.Dummy, true
}

// HasDummy returns a boolean if a field has been set.
func (o *SessionNothing) HasDummy() bool {
	if o != nil && o.Dummy != nil {
		return true
	}

	return false
}

// SetDummy gets a reference to the given bool and assigns it to the Dummy field.
func (o *SessionNothing) SetDummy(v bool) {
	o.Dummy = &v
}

func (o SessionNothing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dummy != nil {
		toSerialize["dummy"] = o.Dummy
	}
	return json.Marshal(toSerialize)
}

type NullableSessionNothing struct {
	value *SessionNothing
	isSet bool
}

func (v NullableSessionNothing) Get() *SessionNothing {
	return v.value
}

func (v *NullableSessionNothing) Set(val *SessionNothing) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionNothing) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionNothing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionNothing(val *SessionNothing) *NullableSessionNothing {
	return &NullableSessionNothing{value: val, isSet: true}
}

func (v NullableSessionNothing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionNothing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

