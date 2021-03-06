/*
 * Power API v1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: lowell@lanl.gov
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package powerapi

import (
	"encoding/json"
)

// ErrorError Properties that describe the error
type ErrorError struct {
	// Response code
	Code string `json:"code"`
	// A human-readable error message string
	Message string `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _ErrorError ErrorError

// NewErrorError instantiates a new ErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorError(code string, message string, ) *ErrorError {
	this := ErrorError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorErrorWithDefaults instantiates a new ErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorErrorWithDefaults() *ErrorError {
	this := ErrorError{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorError) GetCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorError) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorError) SetMessage(v string) {
	o.Message = v
}

func (o ErrorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorError) UnmarshalJSON(bytes []byte) (err error) {
	varErrorError := _ErrorError{}

	if err = json.Unmarshal(bytes, &varErrorError); err == nil {
		*o = ErrorError(varErrorError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorError struct {
	value *ErrorError
	isSet bool
}

func (v NullableErrorError) Get() *ErrorError {
	return v.value
}

func (v *NullableErrorError) Set(val *ErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorError(val *ErrorError) *NullableErrorError {
	return &NullableErrorError{value: val, isSet: true}
}

func (v NullableErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


